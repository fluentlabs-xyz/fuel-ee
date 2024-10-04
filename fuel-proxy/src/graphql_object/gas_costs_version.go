package graphql_object

import "github.com/graphql-go/graphql"

type GasCostsVersionType struct {
	SchemaFields EnumFields
}

//	pub enum GasCostsVersion {
//	   V1,
//	}
type GasCostsVersionStruct struct {
}

func GasCostsVersion() *GasCostsVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "GasCostsVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &GasCostsVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Type:   enum,
		},
	}
}
