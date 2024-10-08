package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
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
func NewBlockType(headerType *HeaderType, transactionType *TransactionType) (*BlockType, error) {

	objectConfig := graphql.ObjectConfig{Name: "Block", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x6e024412c62b6d8158b6c6dd702c0affc760ab489ec735fee637a1a031010784"), nil
			},
		},
		"height": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"header": &graphql.Field{
			Type: headerType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &HeaderStruct{}, nil
			},
		},
		"transactions": &graphql.Field{
			Type: graphql.NewList(transactionType.SchemaFields.Object),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &[]TransactionStruct{}, nil
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
