use alloy_sol_types::SolValue;
use core::{
    mem::take,
    str::{from_utf8, FromStr},
};
use ethers::abi::AbiEncode;
use fluentbase_genesis::{
    devnet::{devnet_genesis_from_file, GENESIS_POSEIDON_HASH_SLOT},
    Genesis
    ,
};
use fluentbase_poseidon::poseidon_hash;
use fluentbase_runtime::{RuntimeContext};
use fluentbase_sdk::runtime::TestingContext;
use fluentbase_types::{calc_create_address, Account, Address, Bytes, NativeAPI, SharedAPI, DEVNET_CHAIN_ID, KECCAK_EMPTY, POSEIDON_EMPTY, PRECOMPILE_FVM, U256};
use fuel_core_types::{
    fuel_asm::op,
    fuel_crypto::{
        rand::{rngs::StdRng, SeedableRng},
        SecretKey,
    },
    fuel_types::{canonical::Serialize, AssetId, BlockHeight, ChainId},
};
use fuel_ee_core::fvm::types::{FvmWithdrawInput, UtxoIdSol, FVM_DRY_RUN_SIG, FVM_DRY_RUN_SIG_BYTES, FVM_EXEC_SIG, FVM_EXEC_SIG_BYTES};
use fuel_ee_core::fvm::{
    helpers::FUEL_TESTNET_BASE_ASSET_ID,
    types::{FVM_DEPOSIT_SIG, FVM_DEPOSIT_SIG_BYTES, FVM_WITHDRAW_SIG, FVM_WITHDRAW_SIG_BYTES},
};
use fuel_tx::field::Witnesses;
use fuel_tx::{ConsensusParameters, Input, TransactionBuilder, TransactionRepr, TxId, TxPointer, UtxoId};
use fuel_vm::{fuel_asm::RegId, storage::MemoryStorage};
use hashbrown::HashMap;
use hex::FromHex;
use revm::{
    primitives::{AccountInfo, Bytecode, Env, ExecutionResult, TransactTo},
    rwasm::RwasmDbWrapper,
    DatabaseCommit,
    Evm,
    InMemoryDB,
};
use rwasm::rwasm::{BinaryFormat, RwasmModule};
use crate::generated::i_fuel_ee::FvmDepositCall;

#[allow(dead_code)]
struct EvmTestingContext {
    sdk: TestingContext,
    genesis: Genesis,
    db: InMemoryDB,
}

impl Default for EvmTestingContext {
    fn default() -> Self {
        Self::load_from_genesis(devnet_genesis_from_file())
    }
}

#[allow(dead_code)]
impl EvmTestingContext {
    fn load_from_genesis(genesis: Genesis) -> Self {
        // create jzkt and put it into testing context
        let mut db = InMemoryDB::default();
        // convert all accounts from genesis into jzkt
        for (k, v) in genesis.alloc.iter() {
            let poseidon_hash = v
                .storage
                .as_ref()
                .and_then(|v| v.get(&GENESIS_POSEIDON_HASH_SLOT).cloned())
                .unwrap_or_else(|| {
                    v.code
                        .as_ref()
                        .map(|v| poseidon_hash(&v).into())
                        .unwrap_or(POSEIDON_EMPTY)
                });
            // let keccak_hash = v
            //     .storage
            //     .as_ref()
            //     .and_then(|v| v.get(&GENESIS_KECCAK_HASH_SLOT).cloned())
            //     .unwrap_or_else(|| {
            //         v.code
            //             .as_ref()
            //             .map(|v| keccak256(&v))
            //             .unwrap_or(KECCAK_EMPTY)
            //     });
            let account = Account {
                address: *k,
                balance: v.balance,
                nonce: v.nonce.unwrap_or_default(),
                // it makes not much sense to fill these fields, but it reduces hash calculation
                // time a bit
                // source_code_size: v.code.as_ref().map(|v| v.len() as u64).unwrap_or_default(),
                // source_code_hash: keccak_hash,
                rwasm_code_size: v.code.as_ref().map(|v| v.len() as u64).unwrap_or_default(),
                rwasm_code_hash: poseidon_hash,
                ..Default::default()
            };
            let mut info: AccountInfo = account.into();
            info.code = v.code.clone().map(Bytecode::new_raw);
            info.rwasm_code = v.code.clone().map(Bytecode::new_raw);
            db.insert_account_info(*k, info);
        }
        let mut testing_ctx = TestingContext::new(RuntimeContext::default());
        let mut res = Self {
            sdk: testing_ctx,
            genesis,
            db,
        };
        res
    }

    pub(crate) fn add_wasm_contract<I: Into<RwasmModule>>(
        &mut self,
        address: Address,
        rwasm_module: I,
    ) -> AccountInfo {
        let rwasm_binary = {
            let rwasm_module: RwasmModule = rwasm_module.into();
            let mut result = Vec::new();
            rwasm_module.write_binary_to_vec(&mut result).unwrap();
            result
        };
        let account = Account {
            address,
            balance: U256::ZERO,
            nonce: 0,
            // it makes not much sense to fill these fields, but it optimizes hash calculation a bit
            source_code_size: 0,
            source_code_hash: KECCAK_EMPTY,
            rwasm_code_size: rwasm_binary.len() as u64,
            rwasm_code_hash: poseidon_hash(&rwasm_binary).into(),
        };
        let mut info: AccountInfo = account.into();
        info.code = None;
        if !rwasm_binary.is_empty() {
            info.rwasm_code = Some(Bytecode::new_raw(rwasm_binary.into()));
        }
        self.db.insert_account_info(address, info.clone());
        info
    }

    pub(crate) fn get_balance(&mut self, address: Address) -> U256 {
        let account = self.db.load_account(address).unwrap();
        account.info.balance
    }

    pub(crate) fn add_balance(&mut self, address: Address, value: U256) {
        let account = self.db.load_account(address).unwrap();
        account.info.balance += value;
        let mut revm_account = revm::primitives::Account::from(account.info.clone());
        revm_account.mark_touch();
        self.db.commit(HashMap::from([(address, revm_account)]));
    }

    pub(crate) fn with_sdk<F>(&mut self, f: F)
    where
        F: Fn(
            RwasmDbWrapper<'_, fluentbase_sdk::runtime::RuntimeContextWrapper, &mut InMemoryDB>,
        ) -> (),
    {
        let mut evm = Evm::builder().with_db(&mut self.db).build();
        let runtime_context = RuntimeContext::default()
            .with_depth(0u32);
        let native_sdk = fluentbase_sdk::runtime::RuntimeContextWrapper::new(runtime_context);
        f(RwasmDbWrapper::new(&mut evm.context.evm, native_sdk))
    }
}

struct TxBuilder<'a> {
    pub(crate) ctx: &'a mut EvmTestingContext,
    pub(crate) env: Env,
}

#[allow(dead_code)]
impl<'a> TxBuilder<'a> {
    fn create(ctx: &'a mut EvmTestingContext, deployer: Address, init_code: Bytes) -> Self {
        let mut env = Env::default();
        env.tx.caller = deployer;
        env.tx.transact_to = TransactTo::Create;
        env.tx.data = init_code;
        env.tx.gas_limit = 300_000_000;
        Self { ctx, env }
    }

    fn call(
        ctx: &'a mut EvmTestingContext,
        caller: Address,
        callee: Address,
        value: Option<U256>,
    ) -> Self {
        let mut env = Env::default();
        if let Some(value) = value {
            env.tx.value = value;
        }
        env.tx.gas_price = U256::from(1);
        env.tx.caller = caller;
        env.tx.transact_to = TransactTo::Call(callee);
        env.tx.gas_limit = 10_000_000;
        Self { ctx, env }
    }

    fn input(mut self, input: Bytes) -> Self {
        self.env.tx.data = input;
        self
    }

    fn value(mut self, value: U256) -> Self {
        self.env.tx.value = value;
        self
    }

    fn chain_id(mut self, chain_id: u64) -> Self {
        self.env.tx.chain_id = Some(chain_id);
        self
    }

    fn gas_limit(mut self, gas_limit: u64) -> Self {
        self.env.tx.gas_limit = gas_limit;
        self
    }

    fn gas_price(mut self, gas_price: U256) -> Self {
        self.env.tx.gas_price = gas_price;
        self
    }

    fn exec(&mut self) -> ExecutionResult {
        let tx_chain_id = self.env.tx.chain_id;
        if tx_chain_id.is_some() {
            self.env.cfg.chain_id = tx_chain_id.unwrap();
        }
        let mut evm = Evm::builder()
            .with_env(Box::new(take(&mut self.env)))
            .with_ref_db(&mut self.ctx.db)
            .build();
        evm.transact_commit().unwrap()
    }
}

fn deploy_evm_tx(ctx: &mut EvmTestingContext, deployer: Address, init_bytecode: Bytes) -> Address {
    // let bytecode_type = BytecodeType::from_slice(init_bytecode.as_ref());
    // deploy greeting EVM contract
    let chain_id = ctx.genesis.config.chain_id;
    let result = TxBuilder::create(ctx, deployer, init_bytecode.clone().into()).chain_id(chain_id).exec();
    if !result.is_success() {
        println!("{:?}", result);
        println!(
            "{}",
            from_utf8(result.output().cloned().unwrap_or_default().as_ref()).unwrap_or("")
        );
    }
    assert!(result.is_success());
    let contract_address = calc_create_address::<TestingContext>(&deployer, 0);
    assert_eq!(contract_address, deployer.create(0));
    // let contract_account = ctx.db.accounts.get(&contract_address).unwrap();
    // if bytecode_type == BytecodeType::EVM {
    //     let source_bytecode = ctx
    //         .db
    //         .contracts
    //         .get(&contract_account.info.code_hash)
    //         .unwrap()
    //         .original_bytes()
    //         .to_vec();
    //     assert_eq!(contract_account.info.code_hash, keccak256(&source_bytecode));
    //     assert!(source_bytecode.len() > 0);
    // }
    // if bytecode_type == BytecodeType::WASM {
    //     let rwasm_bytecode = ctx
    //         .db
    //         .contracts
    //         .get(&contract_account.info.rwasm_code_hash)
    //         .unwrap()
    //         .bytes()
    //         .to_vec();
    //     assert_eq!(
    //         contract_account.info.rwasm_code_hash.0,
    //         poseidon_hash(&rwasm_bytecode)
    //     );
    //     let is_rwasm = rwasm_bytecode.get(0).cloned().unwrap() == 0xef;
    //     assert!(is_rwasm);
    // }
    contract_address
}

fn call_evm_tx_simple(
    ctx: &mut EvmTestingContext,
    caller: Address,
    callee: Address,
    input: Bytes,
    gas_limit: Option<u64>,
    value: Option<U256>,
) -> ExecutionResult {
    // call greeting EVM contract
    let chain_id = ctx.genesis.config.chain_id;
    let mut tx_builder = TxBuilder::call(ctx, caller, callee, value).chain_id(chain_id).input(input);
    if let Some(gas_limit) = gas_limit {
        tx_builder = tx_builder.gas_limit(gas_limit);
    }
    tx_builder.exec()
}

fn call_evm_tx(
    ctx: &mut EvmTestingContext,
    caller: Address,
    callee: Address,
    input: Bytes,
    gas_limit: Option<u64>,
    value: Option<U256>,
) -> ExecutionResult {
    ctx.add_balance(caller, U256::from(1e18));
    call_evm_tx_simple(ctx, caller, callee, input, gas_limit, value)
}


#[test]
fn test_check_signatures_for_collisions() {
    let values = [
        TransactionRepr::Create as u8,
        TransactionRepr::Mint as u8,
        TransactionRepr::Script as u8,
        TransactionRepr::Upgrade as u8,
        TransactionRepr::Upload as u8,
    ];
    assert!(!values.contains(&FVM_DEPOSIT_SIG_BYTES[0]));
    assert!(!values.contains(&FVM_WITHDRAW_SIG_BYTES[0]));
    assert!(!values.contains(&FVM_DRY_RUN_SIG_BYTES[0]));
    assert!(!values.contains(&FVM_EXEC_SIG_BYTES[0]));
}

#[test]
fn test_fvm_deposit_and_transfer_between_accounts_tx() {
    let base_asset_id = AssetId::from_str(FUEL_TESTNET_BASE_ASSET_ID).unwrap();
    let chain_id = DEVNET_CHAIN_ID;
    // let chain_id = 1337;

    let secret1 = "0x99e87b0e9158531eeeb503ff15266e2b23c2a2507b138c9d1b1f2ab458df2d61";
    let secret1_vec = revm::primitives::hex::decode(secret1).unwrap();
    let secret1_secret_key = SecretKey::try_from(secret1_vec.as_slice()).unwrap();
    let secret1_address = Input::owner(&secret1_secret_key.public_key());
    println!("secret1_address: {}", secret1_address);
    let secret1_address_as_evm = Address::from_slice(&secret1_address.as_slice()[12..]);

    let secret2 = "0xde97d8624a438121b86a1956544bd72ed68cd69f2c99555b08b1e8c51ffd511c";
    let secret2_vec = revm::primitives::hex::decode(secret2).unwrap();
    let secret2_secret_key = SecretKey::try_from(secret2_vec.as_slice()).unwrap();
    let secret2_address = Input::owner(&secret2_secret_key.public_key());
    println!("secret2_address: {}", secret2_address);
    let secret2_address_as_evm = Address::from_slice(&secret2_address.as_slice()[12..]);

    let initial_balance = 1000;
    let coins_sent = 10;

    let bytecode = core::iter::once(op::ret(RegId::ZERO)).collect();
    let mut consensus_params = ConsensusParameters::standard();
    consensus_params.set_chain_id(chain_id.into());
    let mut test_builder = fuel_vm::util::test_helpers::TestBuilder {
        rng: StdRng::seed_from_u64(1),
        gas_price: 0,
        max_fee_limit: 600,
        script_gas_limit: 100,
        builder: TransactionBuilder::script(bytecode, vec![]),
        storage: MemoryStorage::default(),
        block_height: Default::default(),
        consensus_params,
    };

    let mut ctx = EvmTestingContext::default();

    // deposit to FVM
    let mut input = Vec::<u8>::new();
    input.extend_from_slice(FVM_DEPOSIT_SIG_BYTES.as_slice());
    input.extend_from_slice(secret2_address.as_slice());
    let result = call_evm_tx(
        &mut ctx,
        secret1_address_as_evm.clone(),
        PRECOMPILE_FVM,
        input.into(),
        Some(100_000_000),
        Some(U256::from(initial_balance * 1_000_000_000)));
    println!("move coins evm->fvm: {:?}", result);
    let output = result.output().unwrap_or_default();
    println!("output: {}", from_utf8(output).unwrap_or_default());
    assert!(result.is_success());

    let tx_in_id: TxId =
        TxId::from_str("0x0000000000000000000000000000000000000000000000000000000000000000")
            .unwrap();
    let utxo_id = UtxoId::new(tx_in_id, 0);
    test_builder
        .with_chain_id(ChainId::new(chain_id))
        .base_asset_id(base_asset_id)
        .start_script([op::ret(0)].to_vec(), Vec::new())
        .builder
        .add_unsigned_coin_input(
            secret2_secret_key.clone(),
            utxo_id,
            initial_balance.clone(),
            base_asset_id,
            TxPointer::new(BlockHeight::new(0), 0),
        )
        .add_output(fuel_tx::Output::coin(
            secret1_address.clone(),
            coins_sent,
            base_asset_id,
        ))
        .add_output(fuel_tx::Output::change(
            secret2_address.clone(),
            0,
            base_asset_id,
        ));
    let mut script_tx1 = test_builder.build().transaction().clone();
    println!(
        "tx1: {:?}",
        &script_tx1
    );
    println!(
        "tx1.witnesses()[0].as_vec(): {:x?}",
        script_tx1.witnesses()[0].as_vec()
    );
    let tx1: fuel_tx::Transaction = fuel_tx::Transaction::Script(script_tx1.clone());
    let fuel_tx_bytes = Bytes::from(tx1.to_bytes());

    println!("\n\n\n");
    let result = call_evm_tx(
        &mut ctx,
        secret2_address_as_evm.clone(),
        PRECOMPILE_FVM,
        fuel_tx_bytes.into(),
        Some(1_000_000_000),
        None,
    );
    println!("call_evm_tx result: {:?}", result);
    let output = result.output().unwrap_or_default();
    println!("output: {}", from_utf8(output).unwrap_or_default());
    assert!(result.is_success());

    // try to use the same coins as input - must report error

    let tx1: fuel_tx::Transaction = fuel_tx::Transaction::Script(script_tx1.clone());
    let fuel_tx_bytes = Bytes::from(tx1.to_bytes());

    println!("\n\n\n");
    let result = call_evm_tx(
        &mut ctx,
        secret2_address_as_evm.clone(),
        PRECOMPILE_FVM,
        fuel_tx_bytes.into(),
        Some(1_000_000_000),
        None,
    );
    println!("call_evm_tx result: {:?}", result);
    let output = result.output().unwrap_or_default();
    println!("output: {}", from_utf8(output).unwrap_or_default());
    assert!(!result.is_success());
}

#[test]
fn test_fvm_deposit_then_withdraw() {
    let base_asset_id = AssetId::from_str(FUEL_TESTNET_BASE_ASSET_ID).unwrap();
    let chain_id = DEVNET_CHAIN_ID;
    let chain_id = 0x1;

    let secret1 = "0x99e87b0e9158531eeeb503ff15266e2b23c2a2507b138c9d1b1f2ab458df2d61";
    let secret1_vec = revm::primitives::hex::decode(secret1).unwrap();
    let secret1_secret_key = SecretKey::try_from(secret1_vec.as_slice()).unwrap();
    let secret1_address = Input::owner(&secret1_secret_key.public_key());
    println!("secret1_address: {}", secret1_address);
    let secret1_address_as_evm = Address::from_slice(&secret1_address.as_slice()[12..]);

    let secret2 = "0xde97d8624a438121b86a1956544bd72ed68cd69f2c99555b08b1e8c51ffd511c";
    let secret2_vec = revm::primitives::hex::decode(secret2).unwrap();
    let secret2_secret_key = SecretKey::try_from(secret2_vec.as_slice()).unwrap();
    let secret2_address = Input::owner(&secret2_secret_key.public_key());
    println!("secret2_address: {}", secret2_address);
    // let secret2_address_as_evm = Address::from_slice(&secret2_address.as_slice()[12..]);

    let coins_sent = 0x500;

    let mut consensus_params = ConsensusParameters::standard();
    consensus_params.set_chain_id(chain_id.into());

    let mut ctx = EvmTestingContext::default();

    ctx.add_balance(secret1_address_as_evm, U256::from(1e18));

    let balance_before_deposit_to_fvm = ctx.get_balance(secret1_address_as_evm);
    assert_eq!(balance_before_deposit_to_fvm, U256::from(1e18));

    // FVM deposit
    let mut input = Vec::<u8>::new();
    // input.extend_from_slice(FVM_DEPOSIT_SIG_BYTES.as_slice());
    let fvm_deposit_call = FvmDepositCall{ address_32: secret2_address.0 };
    input.extend_from_slice(fvm_deposit_call.encode().as_slice());
    let result = call_evm_tx_simple(
        &mut ctx,
        secret1_address_as_evm.clone(),
        PRECOMPILE_FVM,
        input.into(),
        Some(100_000_000),
        Some(U256::from(coins_sent * 1e9 as u64)));
    println!("move coins evm->fvm result: {:?}", result);
    let output = result.output().unwrap_or_default();
    println!("output: {}", from_utf8(output).unwrap_or_default());
    assert!(result.is_success());

    let balance_after_deposit_to_fvm = ctx.get_balance(secret1_address_as_evm);

    let tx_id: TxId =
        TxId::from_str("0000000000000000000000000000000000000000000000000000000000000000").unwrap();
    let output_index = 0x0;

    // FVM withdraw
    let mut input = Vec::<u8>::new();
    input.extend_from_slice(FVM_WITHDRAW_SIG_BYTES.as_slice());
    let utxo_ids: FvmWithdrawInput = FvmWithdrawInput {
        utxos: vec![UtxoIdSol {
            tx_id: tx_id.0.into(),
            output_index,
        }],
        withdraw_amount: 0x1,
    };
    input.extend_from_slice(&utxo_ids.abi_encode());
    let result = call_evm_tx_simple(
        &mut ctx,
        secret1_address_as_evm.clone(),
        PRECOMPILE_FVM,
        input.into(),
        Some(3_000_000_000),
        None);
    println!("move coins fvm->evm result: {:?}", result);
    let output = result.output().unwrap_or_default();
    println!("output: {}", from_utf8(output).unwrap_or_default());
    assert!(result.is_success());

    let balance_after_withdraw_from_fvm = ctx.get_balance(secret1_address_as_evm);
    assert!(balance_after_withdraw_from_fvm > balance_after_deposit_to_fvm);
}
