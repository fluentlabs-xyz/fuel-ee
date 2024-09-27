package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_scalars"
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
}

func MakeCoin() (*CoinType, error) {
	objectConfig := graphql.ObjectConfig{Name: "Coin", Fields: graphql.Fields{
		"amount": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"blockCreated": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"txCreatedIdx": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"owner": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
		"assetId": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
		"utxoId": &graphql.Field{
			Type: graphql_scalars.Bytes34Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
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
