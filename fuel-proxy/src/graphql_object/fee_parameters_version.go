package graphql_object

import "github.com/graphql-go/graphql"

type FeeParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum FeeParametersVersion {
//	   V1,
//	}
type FeeParametersVersionStruct struct {
}

func FeeParametersVersion() *FeeParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "FeeParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &FeeParametersVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Type:   enum,
		},
	}
}
