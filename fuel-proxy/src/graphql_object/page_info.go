package graphql_object

import "github.com/graphql-go/graphql"

type PageInfoType struct {
	SchemaFields SchemaFields
}

//	pub struct PageInfo {
//	   pub end_cursor: Option<String>,
//	   pub has_next_page: bool,
//	   pub has_previous_page: bool,
//	   pub start_cursor: Option<String>,
//	}
type PageInfoStruct struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

func PageInfo() (*PageInfoType, error) {
	objectConfig := graphql.ObjectConfig{Name: "PageInfo", Fields: graphql.Fields{
		"hasPreviousPage": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return false, nil
			},
		},
		"hasNextPage": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return false, nil
			},
		},
		"startCursor": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "", nil
			},
		},
		"endCursor": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "", nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &PageInfoType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
