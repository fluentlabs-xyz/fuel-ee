package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_object"
	"github.com/graphql-go/graphql"
)

type GetChainEntry struct {
	SchemaFields graphql_object.SchemaFields
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
func MakeGetChainEntry(chainInfoType *graphql_object.ChainInfoType) (*GetChainEntry, error) {

	objectConfig := graphql.ObjectConfig{Name: "GetChainEntry", Fields: graphql.Fields{
		"chain": &graphql.Field{
			Type: chainInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &graphql_object.ChainInfoStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetChainEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
