package graphql_input_objects

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type ExcludeInputType struct {
	SchemaFields graphql_object.InputObjectFields
}

//	pub struct ExcludeInput {
//	   utxos: Vec<UtxoId>,
//	   messages: Vec<Nonce>,
//	}
type ExcludeInputStruct struct {
	utxos    [][34]byte
	messages []interface{}
}

var ExcludeInput = graphql.NewInputObject(graphql.InputObjectConfig{Name: "ExcludeInput", Fields: graphql.InputObjectConfigFieldMap{
	"utxos": &graphql.InputObjectFieldConfig{
		Type: graphql.NewList(graphql_scalars.Bytes34Type),
	},
	"messages": &graphql.InputObjectFieldConfig{
		Type: graphql.NewList(graphql_scalars.Bytes32Type),
	},
}})
