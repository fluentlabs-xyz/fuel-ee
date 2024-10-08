package graphql_object

import "github.com/graphql-go/graphql"

type ProgramStateType struct {
	SchemaFields SchemaFields
}

//	pub struct ProgramState {
//		   pub return_type: ReturnType,
//		   pub data: HexString,
//		}
type ProgramStateStruct struct {
	ReturnType string `json:"returnType"`
	Data       string `json:"data"`
}

func MakeProgramStateType() (*ProgramStateType, error) {
	objectConfig := graphql.ObjectConfig{Name: "ProgramState", Fields: graphql.Fields{
		"returnType": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"data": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ProgramStateType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
