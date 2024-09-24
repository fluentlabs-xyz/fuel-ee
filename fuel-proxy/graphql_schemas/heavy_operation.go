package graphql_schemas

import "github.com/graphql-go/graphql"

type HeavyOperationType struct {
	SchemaFields SchemaFields
}

//	pub struct HeavyOperation {
//	   pub base: U64,
//	   pub gas_per_unit: U64,
//	}
type HeavyOperationStruct struct {
	Base       uint64
	GasPerUnit uint64
}

func HeavyOperation() (*HeavyOperationType, error) {
	objectConfig := graphql.ObjectConfig{Name: "HeavyOperation", Fields: graphql.Fields{
		"base": &graphql.Field{
			Type: graphql.String,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	return 1, nil
			//},
		},
		"gasPerUnit": &graphql.Field{
			Type: graphql.String,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	return 1, nil
			//},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &HeavyOperationType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
