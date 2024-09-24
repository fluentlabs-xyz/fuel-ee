package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_schemas"
	"github.com/graphql-go/graphql"
)

type GetNodeInfoType struct {
	SchemaFields graphql_schemas.SchemaFields
}

type GetNodeInfoStruct struct {
}

func GetNodeInfo(nodeInfoType *graphql_schemas.NodeInfoType) (*GetNodeInfoType, error) {
	objectConfig := graphql.ObjectConfig{Name: "GetNodeInfo", Fields: graphql.Fields{
		"nodeInfo": &graphql.Field{
			Type: nodeInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &graphql_schemas.NodeInfoStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetNodeInfoType{
		SchemaFields: graphql_schemas.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
