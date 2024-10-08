package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type CoinType struct {
	SchemaFields SchemaFields
}

//	pub struct Coin {
//		   pub amount: U64,
//		   pub block_created: U32,
//		   pub tx_created_idx: U16,
//		   pub assetId: AssetId,
//		   pub utxo_id: UtxoId,
//	       pub owner: Address,
//	}
type CoinStruct struct {
	Amount       uint64                   `json:"amount"`
	BlockCreated uint64                   `json:"blockCreated"`
	TxCreatedIdx uint64                   `json:"txCreatedIdx"`
	AssetId      *graphql_scalars.Bytes32 `json:"assetId"`
	Owner        *graphql_scalars.Bytes32 `json:"owner"`
	UtxoId       *graphql_scalars.Bytes34 `json:"utxoId"`
}

func MakeCoin() (*CoinType, error) {
	// {
	//  "data": {
	//    "coinsToSpend": [
	//      [
	//        {
	//          "type": "Coin",
	//          "utxoId": "0xa9d5261a68ec08433015f7747d88d0541ced59213224fb96e5ba33e303314afb0001",
	//          "owner": "0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e",
	//          "amount": "1152921504606846975",
	//          "assetId": "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07",
	//          "blockCreated": "1",
	//          "txCreatedIdx": "0"
	//        }
	//      ]
	//    ]
	//  }
	// }
	objectConfig := graphql.ObjectConfig{Name: "Coin", Fields: graphql.Fields{
		"amount": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*CoinStruct)
				if ok {
					return coin.Amount, nil
				}
				return 0, nil
			},
		},
		"blockCreated": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*CoinStruct)
				if ok {
					return coin.BlockCreated, nil
				}
				return 0, nil
			},
		},
		"txCreatedIdx": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*CoinStruct)
				if ok {
					return coin.TxCreatedIdx, nil
				}
				return 0, nil
			},
		},
		"owner": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*CoinStruct)
				if ok {
					return coin.Owner, nil
				}
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
		"assetId": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*CoinStruct)
				if ok {
					return coin.AssetId, nil
				}
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
		"utxoId": &graphql.Field{
			Type: graphql_scalars.Bytes34Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*CoinStruct)
				if ok {
					return coin.UtxoId, nil
				}
				return graphql_scalars.NewBytes34TryFromStringOrPanic("0x00000000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &CoinType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
