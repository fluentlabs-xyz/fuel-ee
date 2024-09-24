package graphql_schemas

import "github.com/graphql-go/graphql"

type FeeParametersType struct {
	SchemaFields SchemaFields
}

//		pub struct FeeParameters {
//	   pub version: FeeParametersVersion,
//	   pub gas_price_factor: U64,
//	   pub gas_per_byte: U64,
//	}
type FeeParametersStruct struct {
}

func FeeParameters(versionType *FeeParametersVersionType) (*FeeParametersType, error) {
	objectConfig := graphql.ObjectConfig{Name: "FeeParameters", Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: versionType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		"gasPriceFactor": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "92000", nil
			},
		},
		"gasPerByte": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "63", nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &FeeParametersType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
