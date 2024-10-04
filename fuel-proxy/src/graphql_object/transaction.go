package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type TransactionType struct {
	SchemaFields SchemaFields
}

type TransactionStruct struct {
}

//	pub struct Header {
//	   pub version: HeaderVersion,
//	   pub id: BlockId,
//	   pub da_height: U64,
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
func Transaction() (*TransactionType, error) {

	objectConfig := graphql.ObjectConfig{Name: "Transaction", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_scalars.NewBytes32TryFromStringOrPanic("a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"), nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &TransactionType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
