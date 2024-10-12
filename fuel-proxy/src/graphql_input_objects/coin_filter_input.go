package graphql_input_objects

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type CoinFilterInputType struct {
	SchemaFields graphql_object.InputObjectFields
}

//	pub struct CoinFilterInput {
//	   pub owner: Address,
//	   pub asset_id: Option<AssetId>,
//	}
type CoinFilterInputStruct struct {
}

var CoinFilterInput = graphql.NewInputObject(graphql.InputObjectConfig{Name: "CoinFilterInput", Fields: graphql.InputObjectConfigFieldMap{
	"owner": &graphql.InputObjectFieldConfig{
		Type: graphql_scalars.Bytes32Type,
	},
	"assetId": &graphql.InputObjectFieldConfig{
		Type: graphql_scalars.Bytes32Type,
	},
}})
