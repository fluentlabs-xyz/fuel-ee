use alloc::{format, vec::Vec};
use alloy_sol_types::SolType;
use core::str::FromStr;
use fluentbase_sdk::Bytes;
use fluentbase_sdk::{basic_entrypoint, derive::Contract, ExitCode, SharedAPI, U256};
use fluentbase_sdk::derive::{router, signature};
use fuel_core_storage::{
    structured_storage::StructuredStorage, tables::Coins, StorageInspect,
    StorageMutate,
};
use fuel_core_types::{
    entities::coins::coin::{CompressedCoin, CompressedCoinV1},
    fuel_types::AssetId,
};
use fuel_ee_core::fvm::exec::_exec_fuel_tx;
use fuel_ee_core::fvm::types::{FVM_DEPOSIT_SIG_BYTES, FVM_DRY_RUN_SIG_BYTES, FVM_EXEC_SIG_BYTES, FVM_WITHDRAW_SIG_BYTES};
use fuel_ee_core::fvm::{
    helpers::FUEL_TESTNET_BASE_ASSET_ID,
    types::{
        FvmDepositInput, FvmWithdrawInput, WasmStorage,
    },
};
use fuel_ee_core::helpers_fvm::{log_deposit, log_withdraw};
use fuel_tx::{TxId, UtxoId};

#[derive(Contract)]
pub struct FvmLoader<SDK> {
    sdk: SDK,
}

impl<SDK: SharedAPI> FvmLoader<SDK> {
    pub fn deploy(&mut self) {
        self.sdk.exit(ExitCode::Ok.into_i32());
    }

    pub fn main(&mut self) {
        let exit_code = self.main_inner();
        self.sdk.exit(exit_code.into_i32());
    }
    pub fn deposit(sdk: &mut SDK, msg: &[u8], asset_id: &AssetId) -> ExitCode {
        let deposit_input: FvmDepositInput =
            <FvmDepositInput as SolType>::abi_decode(msg, true)
                .expect("valid fvm deposit input");
        let recipient_address = fuel_core_types::fuel_types::Address::new(deposit_input.address.0);

        let contract_ctx = sdk.contract_context();
        let caller = contract_ctx.caller;
        let value = contract_ctx.value;

        let evm_balance = sdk.balance(&caller);
        if evm_balance < value {
            return ExitCode::InsufficientBalance;
        }
        if value == U256::default() {
            panic!("value must be greater 0 and will be used as a deposit amount");
        }
        let value_gwei = value / U256::from(1_000_000_000);
        if value != value_gwei * U256::from(1_000_000_000) {
            panic!("can not convert deposit value into gwei without cutting least significant part");
        };

        let mut wasm_storage = WasmStorage { sdk };
        let deposit_withdraw_tx_index =
            wasm_storage.deposit_withdraw_tx_next_index().to_be_bytes();
        let mut storage = StructuredStorage::new(wasm_storage);
        let coin_amount = value_gwei.as_limbs()[0];

        let tx_output_index: u16 = 0;
        let tx_id: TxId = TxId::new(deposit_withdraw_tx_index);
        let utxo_id = UtxoId::new(tx_id, tx_output_index);

        let mut coin = CompressedCoin::V1(CompressedCoinV1::default());
        coin.set_owner(recipient_address);
        coin.set_amount(coin_amount);
        coin.set_asset_id(*asset_id);

        <StructuredStorage<WasmStorage<'_, SDK>> as StorageMutate<Coins>>::insert(
            &mut storage,
            &utxo_id,
            &coin,
        )
            .expect("failed to save deposit utxo");

        log_deposit(sdk, &recipient_address, coin_amount, &tx_id, tx_output_index,  &asset_id);

        ExitCode::Ok
    }
    pub fn withdraw(sdk: &mut SDK, msg: &[u8], asset_id: &AssetId) -> ExitCode {
        let contract_ctx = sdk.contract_context();
        let caller = contract_ctx.caller;
        let fvm_withdraw_input: FvmWithdrawInput =
            <FvmWithdrawInput as SolType>::abi_decode(msg, true)
                .expect("failed to decode FvmWithdrawInput");
        let FvmWithdrawInput {
            utxos,
            withdraw_amount,
        } = fvm_withdraw_input;
        let mut utxos_total_balance = 0;
        let utxos_to_spend: Vec<UtxoId> = utxos
            .iter()
            .map(|v| {
                UtxoId::new(
                    TxId::new(v.tx_id.0),
                    v.output_index,
                )
            })
            .collect();
        if utxos_to_spend.len() <= 0 {
            panic!("utxos must be provided when withdrawing funds")
        }
        let mut last_owner: Option<fuel_core_types::fuel_types::Address> = None;
        for utxo_id in &utxos_to_spend {
            let wasm_storage = WasmStorage { sdk };
            let mut storage = StructuredStorage::new(wasm_storage);
            let coin = <StructuredStorage<WasmStorage<'_, SDK>> as StorageInspect<Coins>>::get(
                &mut storage,
                &utxo_id,
            )
                .expect(&format!("got error when fetching utxo: {}", &utxo_id))
                .expect(&format!("utxo {} doesnt exist", &utxo_id));
            utxos_total_balance += coin.amount();
            if coin.asset_id() != asset_id {
                panic!(
                    "utxo {} asset id doesn't match base asset id {}",
                    &utxo_id, &asset_id
                )
            }
            if let Some(last_owner) = last_owner {
                if &last_owner != coin.owner() {
                    panic!("all utxo owners must be the same across all utxos")
                }
            }
            last_owner = Some(coin.owner().clone());
        }
        // sum all the utxos balances and check if it is more than provided in input
        if utxos_total_balance < withdraw_amount {
            panic!(
                "utxo balance ({}) must be greater withdraw amount ({})",
                &utxos_total_balance, &withdraw_amount
            )
        }

        let mut wasm_storage = WasmStorage { sdk };
        let deposit_withdraw_tx_index =
            wasm_storage.deposit_withdraw_tx_next_index().to_be_bytes();
        let mut storage = StructuredStorage::new(wasm_storage);

        let last_owner = last_owner.expect("utxo owner not found");
        // spend utxos
        for utxo_id in &utxos_to_spend {
            <StructuredStorage<WasmStorage<'_, SDK>> as StorageMutate<Coins>>::remove(
                &mut storage,
                &utxo_id,
            )
                .expect(&format!("failed to remove spent utxo: {}", utxo_id));
        }
        let balance_left = utxos_total_balance - withdraw_amount;
        let mut utxo_id_opt: Option<UtxoId> = None;
        if balance_left > 0 {
            // if there is fvm balance left - create utxo based on balance
            let mut coin = CompressedCoin::V1(CompressedCoinV1::default());
            coin.set_owner(last_owner);
            coin.set_amount(balance_left);
            coin.set_asset_id(*asset_id);
            let tx_id = TxId::new(deposit_withdraw_tx_index);
            let output_index: u16 = 0;
            let utxo_id = UtxoId::new(tx_id, output_index);
            <StructuredStorage<WasmStorage<'_, SDK>> as StorageMutate<Coins>>::insert(
                &mut storage,
                &utxo_id,
                &coin,
            )
                .expect("insert first utxo success");
            utxo_id_opt = Some(utxo_id);
            log_deposit(sdk, &last_owner,balance_left, utxo_id.tx_id(), utxo_id.output_index(), asset_id);
        }
        for utxo_id in &utxos_to_spend {
            log_withdraw(sdk, &last_owner, utxo_id.tx_id(), utxo_id.output_index());
        }

        // top up evm balance
        let withdraw_amount_wei = withdraw_amount as u128 * 1e9 as u128;
        sdk.call(
            caller,
            U256::from(withdraw_amount_wei),
            &[],
            10_000,
        );

        ExitCode::Ok
    }

    pub fn dry_run(sdk: &mut SDK, msg: &[u8]) -> ExitCode {
        let raw_tx_bytes: Bytes = msg.to_vec().into();
        let result = _exec_fuel_tx(sdk, u64::MAX, false, raw_tx_bytes);
        result.exit_code.into()
    }

    pub fn exec(sdk: &mut SDK, msg: &[u8]) -> ExitCode {
        let raw_tx_bytes: Bytes = msg.to_vec().into();
        let result = _exec_fuel_tx(sdk, u64::MAX, true, raw_tx_bytes);
        ExitCode::from(result.exit_code)
    }

    pub fn main_inner(&mut self) -> ExitCode {
        let asset_id: AssetId = AssetId::from_str(FUEL_TESTNET_BASE_ASSET_ID).unwrap();
        let input = self.sdk.input();
        if input.as_ref().starts_with(FVM_DEPOSIT_SIG_BYTES.as_slice()) {
            return FvmLoader::deposit(&mut self.sdk, input.slice(FVM_DEPOSIT_SIG_BYTES.len()..).as_ref(), &asset_id);
        } else if input.as_ref().starts_with(FVM_WITHDRAW_SIG_BYTES.as_slice()) {
            return FvmLoader::withdraw(&mut self.sdk, input.slice(FVM_WITHDRAW_SIG_BYTES.len()..).as_ref(), &asset_id);
        } else if input.as_ref().starts_with(FVM_DRY_RUN_SIG_BYTES.as_slice()) {
            return FvmLoader::dry_run(&mut self.sdk, input.slice(FVM_DRY_RUN_SIG_BYTES.len()..).as_ref());
        } else if input.as_ref().starts_with(FVM_EXEC_SIG_BYTES.as_slice()) {
            return FvmLoader::exec(&mut self.sdk, input.slice(FVM_EXEC_SIG_BYTES.len()..).as_ref());
        }
        panic!("couldn't detect function selector signature")
    }
}

#[derive(Contract)]
pub struct FvmLoaderEntrypoint<SDK> {
    sdk: SDK,
}

pub trait RouterAPI {
    fn fvm_deposit(&mut self, msg: [u8; 32]);
    fn fvm_withdraw(&mut self, msg: &[u8]);
    fn fvm_dry_run(&mut self, msg: &[u8]);
    fn fvm_exec(&mut self, msg: &[u8]);
}

#[router(mode = "solidity")]
impl<SDK: SharedAPI> RouterAPI for FvmLoaderEntrypoint<SDK> {
    #[signature("fvm_deposit(uint8[32])")]
    fn fvm_deposit(&mut self, msg: [u8; 32]) {
        let asset_id: AssetId = AssetId::from_str(FUEL_TESTNET_BASE_ASSET_ID).unwrap();
        let exit_code = FvmLoader::deposit(&mut self.sdk, &msg, &asset_id);
        self.sdk.exit(exit_code.into_i32());
    }

    #[signature("fvm_withdraw(bytes)")]
    fn fvm_withdraw(&mut self, msg: &[u8]) {
        let asset_id: AssetId = AssetId::from_str(FUEL_TESTNET_BASE_ASSET_ID).unwrap();
        let exit_code = FvmLoader::withdraw(&mut self.sdk, msg, &asset_id);
        self.sdk.exit(exit_code.into_i32());
    }

    #[signature("fvm_dry_run(bytes)")]
    fn fvm_dry_run(&mut self, msg: &[u8]) {
        let exit_code = FvmLoader::dry_run(&mut self.sdk, msg);
        self.sdk.exit(exit_code.into_i32());
    }

    #[signature("fvm_exec(bytes)")]
    fn fvm_exec(&mut self, msg: &[u8]) {
        let exit_code = FvmLoader::exec(&mut self.sdk, msg);
        self.sdk.exit(exit_code.into_i32());
    }
}

impl<SDK: SharedAPI> FvmLoaderEntrypoint<SDK> {
    pub fn deploy(&mut self) {
        self.sdk.exit(ExitCode::Ok.into_i32());
    }
}

// basic_entrypoint!(FvmLoader);
basic_entrypoint!(FvmLoaderEntrypoint);
