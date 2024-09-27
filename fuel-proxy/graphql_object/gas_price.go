package graphql_object

import "github.com/graphql-go/graphql"

type GasPriceType struct {
	SchemaFields SchemaFields
}

//	pub struct GasPrice {
//	   pub da_height: U64,
//	   pub name: String,
//	   pub latest_block: Block,
//	   pub consensus_parameters: ConsensusParameters,
//	}
type GasPriceStruct struct {
	GasPrice uint32 `json:"gasPrice"`
}

func GasPrice() (*GasPriceType, error) {
	objectConfig := graphql.ObjectConfig{Name: "GasPrice", Fields: graphql.Fields{
		"gasPrice": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GasPriceType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
