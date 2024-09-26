package graphql_scalars

import (
	"encoding/hex"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strings"
)

type Bytes34 struct {
	value [34]byte
}

func (b *Bytes34) String() string {
	return "0x" + hex.EncodeToString(b.value[:])
}

func NewBytes34(v [34]byte) *Bytes34 {
	return &Bytes34{
		value: v,
	}
}

func NewBytes34TryFromString(v string) (*Bytes34, error) {
	if strings.HasPrefix(v, "0x") {
		v = v[2:]
	}
	output, err := hex.DecodeString(v)
	if err != nil {
		return nil, err
	}
	if len(output) != 34 {
		return nil, fmt.Errorf("output must be 32 bytes len")
	}
	return NewBytes34(([34]byte)(output)), nil
}

func NewBytes34TryFromStringOrPanic(v string) *Bytes34 {
	output, err := NewBytes34TryFromString(v)
	if err != nil {
		panic(fmt.Sprintf("failed to convert string into Bytes34: %s", err))
	}
	return output
}

var Bytes34Type = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Bytes34",
	Description: "The `Bytes34Type` holds fixed 32 bytes array",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case Bytes34:
			return value.String()
		case *Bytes34:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewBytes34TryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewBytes34TryFromString(*value)
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
			v, err := NewBytes34TryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
