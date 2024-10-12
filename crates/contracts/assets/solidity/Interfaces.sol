pragma solidity ^0.8.26;

struct UtxoIdSol {
    uint8[32] tx_id;
    uint16 output_index;
}

struct FvmWithdrawSol {
    uint64 withdraw_amount;
    bytes[] utxo_ids;
}

interface IFuelEE {
    function fvm_deposit(uint8[32] calldata address32) external payable;

    function fvm_withdraw(bytes calldata data) external;

    function fvm_dry_run(uint8[] calldata data) external;

    function fvm_exec(uint8[] calldata data) external;

    function _stub_1(FvmWithdrawSol memory data) external;
}

contract FuelEE {
    event Result(uint8[34], bytes32, bytes32, uint, uint, uint);
}