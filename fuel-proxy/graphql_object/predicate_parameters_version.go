package graphql_object

import "github.com/graphql-go/graphql"

type PredicateParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum PredicateParametersVersion {
//	   V1,
//	}
type PredicateParametersVersionStruct struct {
}

func PredicateParametersVersion() *PredicateParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "PredicateParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &PredicateParametersVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Type:   enum,
		},
	}
}
