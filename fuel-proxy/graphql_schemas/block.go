package graphql_schemas

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_types"
	"github.com/graphql-go/graphql"
)

type BlockType struct {
	SchemaFields SchemaFields
}

type BlockStruct struct {
	Id     [32]byte `json:"id"`
	Height uint32   `json:"height"`
	// pub version: BlockVersion,
	// pub header: Header,
	// pub consensus: Consensus,
	// pub transaction_ids: Vec<TransactionId>,
}

//	pub struct Block {
//	   pub version: BlockVersion,
//	   pub id: BlockId,
//	   pub header: Header,
//	   pub consensus: Consensus,
//	   pub transaction_ids: Vec<TransactionId>,
//	}
func Block(headerType *HeaderType, transactionType *TransactionType) (*BlockType, error) {

	objectConfig := graphql.ObjectConfig{Name: "Block", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql_types.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_types.NewBytes32TryFromStringOrPanic("1212121212121212121212121212121212121212121212121212121212121212"), nil
			},
		},
		"height": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 111, nil
			},
		},
		"header": &graphql.Field{
			Type: headerType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "block header here", nil
			},
		},
		"transactions": &graphql.Field{
			Type: graphql.NewList(transactionType.SchemaFields.Object),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "block header here", nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &BlockType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
