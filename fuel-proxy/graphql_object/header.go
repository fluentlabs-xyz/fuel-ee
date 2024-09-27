package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type HeaderType struct {
	SchemaFields SchemaFields
}

type HeaderStruct struct {
	Time uint64 `json:"time"`
}

//	pub struct Header {
//	   pub id: BlockId,
//	   pub da_height: U64,
//	   pub version: HeaderVersion,
//	   pub consensus_parameters_version: U32,
//	   pub state_transition_bytecode_version: U32,
//	   pub transactions_count: U16,
//	   pub message_receipt_count: U32,
//	   pub transactions_root: Bytes32,
//	   pub message_outbox_root: Bytes32,
//	   pub event_inbox_root: Bytes32,
//	   pub height: U32,
//	   pub prev_root: Bytes32,
//	   pub time: Tai64Timestamp,
//	   pub application_hash: Bytes32,
//	}
func Header() (*HeaderType, error) {

	objectConfig := graphql.ObjectConfig{Name: "Header", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_scalars.NewBytes32TryFromStringOrPanic("1212121212121212121212121212121212121212121212121212121212121212"), nil
			},
		},
		"daHeight": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 4611686018427387914, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &HeaderType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
