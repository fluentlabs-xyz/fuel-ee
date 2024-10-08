use alloc::vec::Vec;
use alloy_primitives::B256;
use alloy_sol_types::SolValue;
use fluentbase_sdk::SharedAPI;
use fuel_core_executor::executor::{
    BlockExecutor,
    ExecutionData,
    ExecutionOptions,
    TxStorageTransaction,
};
use fuel_core_executor::ports::RelayerPort;
use fuel_core_storage::{
    column::Column,
    kv_store::{KeyValueInspect, KeyValueMutate, WriteOperation},
    structured_storage::StructuredStorage,
    transactional::{Changes, ConflictPolicy, InMemoryTransaction, IntoTransaction},
};
use fuel_core_types::fuel_tx::field::{Inputs, Outputs};
use fuel_core_types::{
    blockchain::header::PartialBlockHeader,
    fuel_tx::{Cacheable, ContractId, Receipt, Word},
    fuel_vm::{
        checked_transaction::{Checked, IntoChecked},
        interpreter::{CheckedMetadata, ExecutableTransaction, MemoryInstance},
        ProgramState,
    },
    services::executor::Result,
};
use fuel_core_types::fuel_tx::{Address, AssetId, TxId};
use crate::fvm::types::{FVM_DEPOSIT_SIG_BYTES, FVM_WITHDRAW_SIG_BYTES};

#[derive(Debug, Clone)]
pub struct FvmTransactResult<Tx> {
    pub reverted: bool,
    pub program_state: ProgramState,
    pub tx: Tx,
    pub receipts: Vec<Receipt>,
    pub changes: Changes,
}

pub fn fvm_transact<'a, Tx, T, R>(
    storage: &mut T,
    checked_tx: Checked<Tx>,
    header: &'a PartialBlockHeader,
    coinbase_contract_id: ContractId,
    gas_price: Word,
    memory: &'a mut MemoryInstance,
    execution_options: ExecutionOptions,
    block_executor: BlockExecutor<R>,
    execution_data: &mut ExecutionData,
) -> Result<FvmTransactResult<Tx>>
where
    Tx: ExecutableTransaction + Cacheable + Send + Sync + 'static,
    <Tx as IntoChecked>::Metadata: CheckedMetadata + Send + Sync,
    T: KeyValueInspect<Column=Column>,
    R: RelayerPort
{
    execute_chargeable_transaction(
        storage,
        checked_tx,
        header,
        coinbase_contract_id,
        gas_price,
        memory,
        execution_options,
        execution_data,
        block_executor,
    )
    // match checked_tx {
    //     MaybeCheckedTransaction::CheckedTransaction(tx, _) => {
    //         match tx {
    //             CheckedTransaction::Script(checked_tx) => {
    //                 execute_chargeable_transaction(
    //                     storage,
    //                     checked_tx,
    //                     header,
    //                     coinbase_contract_id,
    //                     gas_price,
    //                     memory,
    //                     execution_options,
    //                     execution_data,
    //                     block_executor,
    //                 )
    //             }
    //             CheckedTransaction::Create(checked_tx) => {
    //                 execute_chargeable_transaction(
    //                     storage,
    //                     checked_tx,
    //                     header,
    //                     coinbase_contract_id,
    //                     gas_price,
    //                     memory,
    //                     execution_options,
    //                     execution_data,
    //                     block_executor,
    //                 )
    //             }
    //             CheckedTransaction::Upgrade(checked_tx) => {
    //                 execute_chargeable_transaction(
    //                     storage,
    //                     checked_tx,
    //                     header,
    //                     coinbase_contract_id,
    //                     gas_price,
    //                     memory,
    //                     execution_options,
    //                     execution_data,
    //                     block_executor,
    //                 )
    //             }
    //             CheckedTransaction::Upload(checked_tx) => {
    //                 execute_chargeable_transaction(
    //                     storage,
    //                     checked_tx,
    //                     header,
    //                     coinbase_contract_id,
    //                     gas_price,
    //                     memory,
    //                     execution_options,
    //                     execution_data,
    //                     block_executor,
    //                 )
    //             }
    //             CheckedTransaction::Mint(_) => {
    //                 panic!("mint transaction not supported")
    //             }
    //         }
    //     }
    //     MaybeCheckedTransaction::Transaction(tx) => {
    //         panic!("pure tx not supported yet");
    //         // let block_height = *header.height();
    //         // let checked_tx = tx
    //         //     .into_checked_basic(block_height, &consensus_params)
    //         //     .into();
    //         // execute_chargeable_transaction(
    //         //     storage,
    //         //     checked_tx,
    //         //     header,
    //         //     coinbase_contract_id,
    //         //     gas_price,
    //         //     memory,
    //         //     consensus_params,
    //         //     extra_tx_checks,
    //         //     execution_data,
    //         // )
    //     }
    // }
}

fn execute_chargeable_transaction<'a, Tx, T, R>(
    storage: &mut T,
    checked_tx: Checked<Tx>,
    header: &'a PartialBlockHeader,
    coinbase_contract_id: ContractId,
    gas_price: Word,
    memory: &'a mut MemoryInstance,
    execution_options: ExecutionOptions,
    execution_data: &mut ExecutionData,
    block_executor: BlockExecutor<R>,
) -> Result<FvmTransactResult<Tx>>
where
    Tx: ExecutableTransaction + Cacheable + Send + Sync + 'static,
    <Tx as IntoChecked>::Metadata: CheckedMetadata + Send + Sync,
    T: KeyValueInspect<Column=Column>,
    R: RelayerPort
{
    let structured_storage = StructuredStorage::new(storage);
    let mut structured_storage = structured_storage.into_transaction();
    let in_memory_transaction = InMemoryTransaction::new(
        Changes::new(),
        ConflictPolicy::Overwrite,
        &mut structured_storage,
    );
    let tx_transaction = &mut TxStorageTransaction::new(in_memory_transaction);

    let tx_id = checked_tx.id();

    let mut checked_tx = checked_tx;
    if execution_options.extra_tx_checks {
        checked_tx = block_executor.extra_tx_checks(checked_tx, header, tx_transaction, memory)?;
    }

    let (reverted, program_state, tx, receipts) = block_executor.attempt_tx_execution_with_vm(
        checked_tx,
        header,
        coinbase_contract_id,
        gas_price,
        tx_transaction,
        memory,
    )?;

    block_executor.spend_input_utxos(tx.inputs(), tx_transaction, reverted, execution_data)?;

    block_executor.persist_output_utxos(
        *header.height(),
        execution_data,
        &tx_id,
        tx_transaction,
        tx.inputs(),
        tx.outputs(),
    )?;

    block_executor.update_execution_data(
        &tx,
        execution_data,
        receipts.clone(),
        gas_price,
        reverted,
        program_state,
        tx_id,
    )?;

    Ok(FvmTransactResult {
        reverted,
        program_state,
        tx,
        receipts,
        changes: tx_transaction.changes().clone(),
    })
}

pub fn fvm_transact_commit<Tx, T, R>(
    storage: &mut T,
    checked_tx: Checked<Tx>,
    header: &PartialBlockHeader,
    coinbase_contract_id: ContractId,
    gas_price: Word,
    execution_options: ExecutionOptions,
    block_executor: BlockExecutor<R>,
    execution_data: &mut ExecutionData,
) -> Result<FvmTransactResult<Tx>>
where
    Tx: ExecutableTransaction + Cacheable + Send + Sync + 'static,
    <Tx as IntoChecked>::Metadata: CheckedMetadata + Send + Sync,
    T: KeyValueMutate<Column=Column>,
    R: RelayerPort
{
    // debug_log!("ecl(fvm_transact_commit): start");

    let mut memory = MemoryInstance::new();

    let result = fvm_transact(
        storage,
        checked_tx,
        header,
        coinbase_contract_id,
        gas_price,
        &mut memory,
        execution_options,
        block_executor,
        execution_data,
    )?;

    for (col_num, changes) in &result.changes {
        let column: Column = col_num.clone().try_into().expect("valid column number");
        match column {
            Column::Metadata
            | Column::ContractsRawCode
            | Column::ContractsState
            | Column::ContractsLatestUtxo
            | Column::ContractsAssets
            | Column::ContractsAssetsMerkleData
            | Column::ContractsAssetsMerkleMetadata
            | Column::ContractsStateMerkleData
            | Column::ContractsStateMerkleMetadata
            | Column::Coins => {
                for (key, op) in changes {
                    match op {
                        WriteOperation::Insert(v) => {
                            storage.write(key, column, v)?;
                        }
                        WriteOperation::Remove => {
                            storage.delete(key, column)?;
                        }
                    }
                }
            }

            Column::Transactions
            | Column::FuelBlocks
            | Column::FuelBlockMerkleData
            | Column::FuelBlockMerkleMetadata
            | Column::Messages
            | Column::ProcessedTransactions
            | Column::FuelBlockConsensus
            | Column::ConsensusParametersVersions
            | Column::StateTransitionBytecodeVersions
            | Column::UploadedBytecodes
            | Column::GenesisMetadata => {
                panic!("unsupported column {:?} operation", column)
            }
        }
    }

    Ok(result)
}

pub fn log_deposit<SDK: SharedAPI>(sdk: &mut SDK, owner_address: &Address, coin_amount: u64, tx_id: &TxId, tx_output_index: u16, asset_id: &AssetId) {
    let log_data = (owner_address.0, coin_amount, tx_id.0, tx_output_index, asset_id.0).abi_encode();
    let topics = [B256::left_padding_from(FVM_DEPOSIT_SIG_BYTES.as_slice())];
    sdk.emit_log(log_data.into(), &topics);
}

pub fn log_withdraw<SDK: SharedAPI>(sdk: &mut SDK, owner_address: &Address, tx_id: &TxId, tx_output_index: u16) {
    let log_data = (owner_address.0, tx_id.0, tx_output_index).abi_encode();
    let topics = [B256::left_padding_from(FVM_WITHDRAW_SIG_BYTES.as_slice())];
    sdk.emit_log(log_data.into(), &topics);
}
