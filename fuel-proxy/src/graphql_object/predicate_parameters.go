package graphql_object

import "github.com/graphql-go/graphql"

type PredicateParametersType struct {
	SchemaFields SchemaFields
}

//	pub struct PredicateParameters {
//	   pub version: PredicateParametersVersion,
//	   pub max_predicate_length: U64,
//	   pub max_predicate_data_length: U64,
//	   pub max_message_data_length: U64,
//	   pub max_gas_per_predicate: U64,
//	}
type PredicateParametersStruct struct {
}

func PredicateParameters(predicateParametersVersionType *PredicateParametersVersionType) (*PredicateParametersType, error) {
	objectConfig := graphql.ObjectConfig{Name: "PredicateParameters", Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: predicateParametersVersionType.SchemaFields.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		"chainId": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"maxPredicateLength": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102400, nil
			},
		},
		"maxPredicateDataLength": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102400, nil
			},
		},
		"maxGasPerPredicate": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 30000000, nil
			},
		},
		"maxMessageDataLength": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102400, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &PredicateParametersType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
