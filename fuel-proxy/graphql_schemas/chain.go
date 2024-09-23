package graphql_schemas

import (
	"github.com/graphql-go/graphql"
)

type ChainType struct {
	SchemaFields SchemaFields
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
func Chain(chainInfoType *ChainInfoType) (*ChainType, error) {

	objectConfig := graphql.ObjectConfig{Name: "Chain", Fields: graphql.Fields{
		"chain": &graphql.Field{
			Type: chainInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "123", nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ChainType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
