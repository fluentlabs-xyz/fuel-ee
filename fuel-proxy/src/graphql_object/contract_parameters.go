package graphql_object

import "github.com/graphql-go/graphql"

type ContractParametersType struct {
	SchemaFields SchemaFields
}

//	pub struct ContractParameters {
//	   pub version: ContractParametersVersion,
//	   pub contract_max_size: U64,
//	   pub max_storage_slots: U64,
//	}
type ContractParametersStruct struct {
	ContractMaxSize uint64 `json:"contractMaxSize"`
	MaxStorageSlots uint64 `json:"maxStorageSlots"`
}

func ContractParameters(versionType *ContractParametersVersionType) (*ContractParametersType, error) {
	objectConfig := graphql.ObjectConfig{Name: "ContractParameters", Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: versionType.SchemaFields.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		"contractMaxSize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102400, nil
			},
		},
		"maxStorageSlots": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1760, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ContractParametersType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
