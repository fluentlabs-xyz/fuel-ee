use alloc::{format, vec::Vec};
use alloy_sol_types::{SolType, SolValue};
use core::str::FromStr;
use fluentbase_sdk::Bytes;
use fluentbase_sdk::{basic_entrypoint, derive::Contract, ExitCode, SharedAPI, B256, U256};
use fuel_core_storage::{
    structured_storage::StructuredStorage, tables::Coins, StorageInspect,
    StorageMutate,
};
use fuel_core_types::{
    entities::coins::coin::{CompressedCoin, CompressedCoinV1},
    fuel_types::AssetId,
};
use fuel_ee_core::fvm::exec::_exec_fuel_tx;
use fuel_ee_core::fvm::types::{FVM_DEPOSIT_SIG_BYTES, FVM_DRY_RUN_SIG_BYTES, FVM_WITHDRAW_SIG_BYTES};
use fuel_ee_core::fvm::{
    helpers::FUEL_TESTNET_BASE_ASSET_ID,
    types::{
        FvmDepositInput, FvmWithdrawInput, WasmStorage,
    },
};
use fuel_tx::{TxId, UtxoId};

#[derive(Contract)]
pub struct FvmLoaderEntrypoint<SDK> {
    sdk: SDK,
}

impl<SDK: SharedAPI> FvmLoaderEntrypoint<SDK> {
    pub fn deploy(&mut self) {
        self.sdk.exit(ExitCode::Ok.into_i32());
    }

    pub fn main(&mut self) {
        let exit_code = self.main_inner();
        self.sdk.exit(exit_code.into_i32());
    }

    pub fn main_inner(&mut self) -> ExitCode {
        let base_asset_id: AssetId = AssetId::from_str(FUEL_TESTNET_BASE_ASSET_ID).unwrap();
        let asset_id = base_asset_id;
        let input = self.sdk.input();
        if input.as_ref().starts_with(FVM_DEPOSIT_SIG_BYTES.as_slice()) {
            let deposit_input: FvmDepositInput =
                <FvmDepositInput as SolType>::abi_decode(&input.slice(FVM_DEPOSIT_SIG_BYTES.len()..).as_ref(), true)
                    .expect("valid fvm deposit input");
            let recipient_address = fuel_core_types::fuel_types::Address::new(deposit_input.address.0);

            let contract_ctx = self.sdk.contract_context();
            let caller = contract_ctx.caller;
            let value = contract_ctx.value;

            let evm_balance = self.sdk.balance(&caller);
            if evm_balance < value {
                return ExitCode::InsufficientBalance;
            }
            if value == U256::default() {
                panic!("value must be greater 0 and is used as a deposit amount");
            }
            let value_gwei = value / U256::from(1_000_000_000);
            if value != value_gwei * U256::from(1_000_000_000) {
                panic!("can not convert deposit value into gwei without cutting least significant part");
            };

            let mut wasm_storage = WasmStorage { sdk: &mut self.sdk };
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
            coin.set_asset_id(asset_id);

            <StructuredStorage<WasmStorage<'_, SDK>> as StorageMutate<Coins>>::insert(
                &mut storage,
                &utxo_id,
                &coin,
            )
                .expect("failed to save deposit utxo");

            // UtxoId:       *graphql_scalars.NewBytes34TryFromStringOrPanic("0x00000000000000000000000000000000000000000000000000000000000000000000"),
            // Amount:       1000,
            // Owner:        *graphql_scalars.NewBytes32TryFromStringOrPanic("0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e"),
            // AssetId:      *graphql_scalars.NewBytes32TryFromStringOrPanic(types.FuelBaseAssetId),
            // BlockCreated: 1,
            // TxCreatedIdx: 0,
            let log_data = (coin_amount, tx_id.0, tx_output_index, recipient_address.0, asset_id.0).abi_encode();
            let topics = [B256::left_padding_from(FVM_DEPOSIT_SIG_BYTES.as_slice())];
            self.sdk.emit_log(log_data.into(), &topics);

            return ExitCode::Ok;
        } else if input.as_ref().starts_with(FVM_WITHDRAW_SIG_BYTES.as_slice()) {
            let contract_ctx = self.sdk.contract_context();
            let caller = contract_ctx.caller;
            let utxo_ids: FvmWithdrawInput =
                <FvmWithdrawInput as SolType>::abi_decode(input.slice(FVM_WITHDRAW_SIG_BYTES.len()..).as_ref(), true)
                    .expect("valid fvm withdraw input");
            // panic!("utxo_ids {:?}", &utxo_ids);
            let FvmWithdrawInput {
                utxos,
                withdraw_amount,
            } = utxo_ids;
            let mut utxos_total_balance = 0;
            let withdraw_amount = withdraw_amount.as_limbs()[0];
            let utxos: Vec<UtxoId> = utxos
                .iter()
                .map(|v| {
                    UtxoId::new(
                        TxId::new(v.tx_id.0),
                        v.output_index.as_limbs()[0]
                            .try_into()
                            .expect("output index is a valid u16 number"),
                    )
                })
                .collect();
            if utxos.len() <= 0 {
                panic!("provide utxos when withdrawing funds")
            }
            let mut last_owner: Option<fuel_core_types::fuel_types::Address> = None;
            for utxo_id in &utxos {
                let wasm_storage = WasmStorage { sdk: &mut self.sdk };
                let mut storage = StructuredStorage::new(wasm_storage);
                let coin = <StructuredStorage<WasmStorage<'_, SDK>> as StorageInspect<Coins>>::get(
                    &mut storage,
                    &utxo_id,
                )
                    .expect(&format!("got error when fetching utxo: {}", &utxo_id))
                    .expect(&format!("utxo {} doesnt exist", &utxo_id));
                utxos_total_balance += coin.amount();
                if coin.asset_id() != &base_asset_id {
                    panic!(
                        "utxo {} asset id doesn't match base asset id {}",
                        &utxo_id, &base_asset_id
                    )
                }
                if let Some(last_owner) = last_owner {
                    if &last_owner != coin.owner() {
                        panic!("all utxo owners must be the same")
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

            let mut wasm_storage = WasmStorage { sdk: &mut self.sdk };
            let deposit_withdraw_tx_index =
                wasm_storage.deposit_withdraw_tx_next_index().to_be_bytes();
            let mut storage = StructuredStorage::new(wasm_storage);

            // spend utxos (just delete them)
            for utxo in &utxos {
                <StructuredStorage<WasmStorage<'_, SDK>> as StorageMutate<Coins>>::remove(
                    &mut storage,
                    &utxo,
                )
                    .expect(&format!("failed to remove spent utxo: {}", utxo));
            }
            let balance_left = utxos_total_balance - withdraw_amount;
            let last_owner = last_owner.expect("utxo owner not found");
            let mut utxo_id_opt: Option<UtxoId> = None;
            if balance_left > 0 {
                // if there is fvm balance left - create utxo based on balance
                let mut coin = CompressedCoin::V1(CompressedCoinV1::default());
                coin.set_owner(last_owner);
                coin.set_amount(balance_left);
                coin.set_asset_id(base_asset_id);
                // TODO need counter to form TxId dynamically and without collisions
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
            }

            // top up evm balance
            let withdraw_amount_wei = withdraw_amount as u128 * 1e9 as u128;
            self.sdk.call(
                caller,
                U256::from(withdraw_amount_wei),
                &[],
                10_000,
            );

            let log_data = if let Some(utxo_id) = utxo_id_opt {
                (withdraw_amount_wei, utxo_id.tx_id().0, utxo_id.output_index()).abi_encode()
            } else {
                (withdraw_amount_wei).abi_encode()
            };
            let topics = [B256::left_padding_from(FVM_WITHDRAW_SIG_BYTES.as_slice()), last_owner.0.into()];
            self.sdk.emit_log(log_data.into(), &topics);

            return ExitCode::Ok;
        } else if input.as_ref().starts_with(FVM_DRY_RUN_SIG_BYTES.as_slice()) {
            let raw_tx_bytes: Bytes = input.slice(FVM_DRY_RUN_SIG_BYTES.len()..).into();
            let result = _exec_fuel_tx(&mut self.sdk, u64::MAX, false, raw_tx_bytes);
            return result.exit_code.into()
        }

        let result = _exec_fuel_tx(&mut self.sdk, u64::MAX, true, input);
        result.exit_code.into()
    }
}

basic_entrypoint!(FvmLoaderEntrypoint);
