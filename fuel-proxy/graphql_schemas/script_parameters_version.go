package graphql_schemas

import "github.com/graphql-go/graphql"

type ScriptParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum ScriptParametersVersion {
//	   V1,
//	}
type ScriptParametersVersionStruct struct {
}

func ScriptParametersVersion() *ScriptParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "ScriptParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &ScriptParametersVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Entity: enum,
		},
	}
}
