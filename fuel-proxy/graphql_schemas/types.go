package graphql_schemas

import (
	"github.com/graphql-go/graphql"
)

type SchemaFields struct {
	Schema       *graphql.Schema
	ObjectConfig *graphql.ObjectConfig
	Object       *graphql.Object
	SchemaConfig *graphql.SchemaConfig
}

type EnumFields struct {
	Config *graphql.EnumConfig
	Entity *graphql.Enum
}

type UnionFields struct {
	Config *graphql.UnionConfig
	Entity *graphql.Union
}
