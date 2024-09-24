package graphql_schemas

import "github.com/graphql-go/graphql"

type ScriptParametersType struct {
	SchemaFields SchemaFields
}

//	pub struct ScriptParameters {
//	   pub version: ScriptParametersVersion,
//	   pub max_script_length: U64,
//	   pub max_script_data_length: U64,
//	}
type ScriptParametersStruct struct {
	MaxScriptLength     uint64 `json:"maxScriptLength"`
	MaxScriptDataLength uint64 `json:"maxScriptDataLength"`
}

func ScriptParameters(versionType *ScriptParametersVersionType) (*ScriptParametersType, error) {
	objectConfig := graphql.ObjectConfig{Name: "ScriptParameters", Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: versionType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		"maxScriptLength": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102400, nil
			},
		},
		"maxScriptDataLength": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102400, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ScriptParametersType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
