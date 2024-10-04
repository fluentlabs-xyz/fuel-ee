package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type SpendQueryElementInputType struct {
	SchemaFields SchemaFields
}

//	pub struct SpendQueryElementInput {
//	   pub assetId: AssetId,
//	   pub amount: U64,
//	   pub max: Option<U32>,
//	}

var SpendQueryElementInput = graphql.NewInputObject(graphql.InputObjectConfig{Name: "SpendQueryElementInput", Fields: graphql.InputObjectConfigFieldMap{
	"assetId": &graphql.InputObjectFieldConfig{
		Type: graphql_scalars.Bytes32Type,
	},
	"amount": &graphql.InputObjectFieldConfig{
		Type: graphql.Int,
	},
	"max": &graphql.InputObjectFieldConfig{
		Type: graphql.Int,
	},
}})
