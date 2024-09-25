package graphql_types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strconv"
)

type U64 struct {
	value uint64
}

func (b *U64) String() string {
	return strconv.FormatUint(b.value, 10)
}

func NewU64(v uint64) *U64 {
	return &U64{
		value: v,
	}
}

func NewU64TryFromString(v string) (*U64, error) {
	out, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return nil, err
	}
	return NewU64(out), nil
}

var U64Type = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "U64",
	Description: "The `U64` holds unsigned 64 bit value",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case U64:
			return value.String()
		case *U64:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewU64TryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewU64TryFromString(*value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			v, err := NewU64TryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
