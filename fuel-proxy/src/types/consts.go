package types

import "github.com/fluentlabs-xyz/fuel-ee/src/helpers"

const FuelTxOwnerAddress = "0x369f74918912b80c9947d6A174c0C6e2c95fAe1D"
const FuelBaseAssetId = "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07"
const EthGenesisAccount1Address = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const FuelRelayerAccountAddress = EthGenesisAccount1Address
const EthFuelVMPrecompileAddress = "0x0000000000000000000000000000000000005250"

const FvmDepositSig = 3934318243 // _fvm_deposit(uint64)
var FvmDepositSigBytes = helpers.Uint32ToBytesBEMust(FvmDepositSig, 4)
var FvmDepositSig32Bytes = helpers.Uint32ToBytesBEMust(FvmDepositSig, 32)

const FvmWithdrawSig = 2866282671 // _fvm_withdraw(uint64)
var FvmWithdrawSigBytes = helpers.Uint32ToBytesBEMust(FvmWithdrawSig, 4)
var FvmWithdrawSig32Bytes = helpers.Uint32ToBytesBEMust(FvmWithdrawSig, 32)

const FvmDryRunSig = 567857912 // _fvm_dry_run(uint64)
var FvmDryRunSigBytes = helpers.Uint32ToBytesBEMust(FvmDryRunSig, 4)
var FvmDryRunSig32Bytes = helpers.Uint32ToBytesBEMust(FvmDryRunSig, 32)

const FvmExecSig = 1692387067 // _fvm_exec(uint64)
var FvmExecSigBytes = helpers.Uint32ToBytesBEMust(FvmExecSig, 4)
var FvmExecSig32Bytes = helpers.Uint32ToBytesBEMust(FvmExecSig, 32)
