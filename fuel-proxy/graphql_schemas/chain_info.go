package graphql_schemas

import "github.com/graphql-go/graphql"

type ChainInfoType struct {
	SchemaFields SchemaFields
}

//	pub struct ChainInfo {
//	   pub da_height: U64,
//	   pub name: String,
//	   pub latest_block: Block,
//	   pub consensus_parameters: ConsensusParameters,
//	}
func ChainInfo(blockType *BlockType, consensusParametersType *ConsensusParametersType) (*ChainInfoType, error) {

	objectConfig := graphql.ObjectConfig{Name: "ChainInfo", Fields: graphql.Fields{
		"daHeight": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 12, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "some chain info name here", nil
			},
		},
		"latestBlock": &graphql.Field{
			Type: blockType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return BlockStruct{
					Id:     [32]byte{1, 2, 3},
					Height: 321,
				}, nil
			},
		},
		"consensusParameters": &graphql.Field{
			Type: consensusParametersType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ConsensusParametersStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ChainInfoType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
