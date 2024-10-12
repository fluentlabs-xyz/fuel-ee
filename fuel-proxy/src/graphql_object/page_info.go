package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

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

func MakePageInfoType() (*PageInfoType, error) {
	objectConfig := graphql.ObjectConfig{Name: "PageInfo", Fields: graphql.Fields{
		"hasPreviousPage": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v, ok := p.Source.(bool)
				if ok {
					return v, nil
				}
				return false, nil
			},
		},
		"hasNextPage": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v, ok := p.Source.(bool)
				if ok {
					return v, nil
				}
				return false, nil
			},
		},
		"startCursor": &graphql.Field{
			Type: graphql_scalars.Bytes34Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v, ok := p.Source.(*graphql_scalars.Bytes34)
				if ok {
					return v, nil
				}
				return graphql_scalars.NewBytes34Zero(), nil
			},
		},
		"endCursor": &graphql.Field{
			Type: graphql_scalars.Bytes34Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				v, ok := p.Source.(*graphql_scalars.Bytes34)
				if ok {
					return v, nil
				}
				return graphql_scalars.NewBytes34Zero(), nil
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
