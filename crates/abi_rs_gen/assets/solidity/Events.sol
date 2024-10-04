pragma solidity ^0.8.26;

contract FuelEE {
    // UtxoId:       *graphql_scalars.NewBytes34TryFromStringOrPanic("0x00000000000000000000000000000000000000000000000000000000000000000000"),
    // Owner:        *graphql_scalars.NewBytes32TryFromStringOrPanic("0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e"),
    // AssetId:      *graphql_scalars.NewBytes32TryFromStringOrPanic(types.FuelBaseAssetId),
    // Amount:       1000,
    // BlockCreated: 1,
    // TxCreatedIdx: 0,
    event Result(uint8[34], bytes32, bytes32, uint, uint, uint);
}