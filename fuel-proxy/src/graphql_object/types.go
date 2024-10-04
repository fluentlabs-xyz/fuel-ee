package graphql_object

import (
	"github.com/graphql-go/graphql"
)

type SchemaFields struct {
	Schema       *graphql.Schema
	ObjectConfig *graphql.ObjectConfig
	Object       *graphql.Object
	SchemaConfig *graphql.SchemaConfig
}

type InputObjectFields struct {
	ObjectConfig *graphql.InputObjectConfig
	Object       *graphql.InputObject
}

type EnumFields struct {
	Config *graphql.EnumConfig
	Type   *graphql.Enum
}

type UnionFields struct {
	Config *graphql.UnionConfig
	Type   *graphql.Union
}
