package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/graphql-go/graphql"
)

type GetNodeInfoEntry struct {
	SchemaFields graphql_object.SchemaFields
}

type GetNodeInfoStruct struct {
}

func MakeGetNodeInfoEntry(nodeInfoType *graphql_object.NodeInfoType) (*GetNodeInfoEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "GetNodeInfoEntry", Fields: graphql.Fields{
		"nodeInfo": &graphql.Field{
			Type: nodeInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &graphql_object.NodeInfoStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetNodeInfoEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
