package graphql_object

import "github.com/graphql-go/graphql"

type ContractParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum ContractParametersVersion {
//	   V1,
//	}
type ContractParametersVersionStruct struct {
}

func ContractParametersVersion() *ContractParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "ContractParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &ContractParametersVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Type:   enum,
		},
	}
}
