package graphql_schemas

import "github.com/graphql-go/graphql"

type NodeInfoRequestType struct {
	SchemaFields SchemaFields
}

type NodeInfoRequestStruct struct {
}

func NodeInfoRequest(nodeInfoType *NodeInfoType) (*NodeInfoRequestType, error) {
	objectConfig := graphql.ObjectConfig{Name: "NodeInfoRequest", Fields: graphql.Fields{
		"nodeInfo": &graphql.Field{
			Type: nodeInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &NodeInfoStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &NodeInfoRequestType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
