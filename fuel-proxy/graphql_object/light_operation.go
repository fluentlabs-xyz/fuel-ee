package graphql_object

import "github.com/graphql-go/graphql"

type LightOperationType struct {
	SchemaFields SchemaFields
}

//	pub struct LightOperation {
//	   pub base: U64,
//	   pub units_per_gas: U64,
//	}
type LightOperationStruct struct {
	Base        uint64
	UnitsPerGas uint64
}

func LightOperation() (*LightOperationType, error) {
	objectConfig := graphql.ObjectConfig{Name: "LightOperation", Fields: graphql.Fields{
		"base": &graphql.Field{
			Type: graphql.String,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	return 1, nil
			//},
		},
		"unitsPerGas": &graphql.Field{
			Type: graphql.String,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	return 1, nil
			//},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &LightOperationType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
