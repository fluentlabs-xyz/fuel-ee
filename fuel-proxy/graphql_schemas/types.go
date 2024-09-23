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
	EnumConfig *graphql.EnumConfig
	Enum       *graphql.Enum
}
