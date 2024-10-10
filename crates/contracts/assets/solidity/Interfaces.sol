pragma solidity ^0.8.26;

interface IFuelEE {
    function fvmDeposit(uint8[32] calldata address32) external payable;
    function fvmWithdraw(bytes calldata data) external;
    function fvmDryRun(bytes calldata data) external;
    function fvmExec(bytes calldata data) external;
}

contract FuelEE {
    event Result(uint8[34], bytes32, bytes32, uint, uint, uint);
}