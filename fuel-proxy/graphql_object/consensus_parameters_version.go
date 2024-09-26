package graphql_object

import "github.com/graphql-go/graphql"

type ConsensusParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum ConsensusParametersVersion {
//		V1,
//	 }
type ConsensusParametersVersionStruct struct {
}

func ConsensusParametersVersion() *ConsensusParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "ConsensusParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &ConsensusParametersVersionType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Type:   enum,
		},
	}
}
