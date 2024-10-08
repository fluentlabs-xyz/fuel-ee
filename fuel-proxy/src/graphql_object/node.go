package graphql_object

import "github.com/graphql-go/graphql"

type NodeType struct {
	SchemaFields SchemaFields
}

type NodeStruct struct {
	Node interface{} `json:"node"`
}

func NewNodeType(subtype *graphql.Object) (*NodeType, error) {
	objectConfig := graphql.ObjectConfig{Name: "Node", Fields: graphql.Fields{
		"node": &graphql.Field{
			Type: subtype,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	return true, nil
			//},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &NodeType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
