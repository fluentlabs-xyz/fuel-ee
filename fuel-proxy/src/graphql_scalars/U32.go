package graphql_scalars

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strconv"
)

type U32 struct {
	value uint32
}

func (b *U32) String() string {
	return strconv.FormatUint(uint64(b.value), 10)
}

func NewU32(v uint32) *U32 {
	return &U32{
		value: v,
	}
}

func NewU32TryFromString(v string) (*U32, error) {
	out64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return nil, err
	}
	out32 := uint32(out64)
	return NewU32(out32), nil
}

//func NewU32TryFromStringOrPanic(v string) *U32 {
//	output, err := NewU32TryFromString(v)
//	if err != nil {
//		panic(fmt.Sprintf("failed to convert string into U32: %s", err))
//	}
//	return output
//}

var U32Type = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "U32",
	Description: "The `U32` holds unsigned 32 bit value",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case U32:
			return value.String()
		case *U32:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewU32TryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewU32TryFromString(*value)
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
			v, err := NewU32TryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
