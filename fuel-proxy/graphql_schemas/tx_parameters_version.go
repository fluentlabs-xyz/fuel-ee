package graphql_schemas

import "github.com/graphql-go/graphql"

type TxParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum TxParametersVersion {
//	   V1,
//	}
type TxParametersVersionStruct struct {
}

func TxParametersVersion() *TxParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "TxParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &TxParametersVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Entity: enum,
		},
	}
}
