package graphql_scalars

import (
	"encoding/hex"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strings"
)

type HexString struct {
	value []byte
}

func (b *HexString) String() string {
	return hex.EncodeToString(b.value)
}

func NewHexString(v []byte) *HexString {
	return &HexString{
		value: v,
	}
}

func NewHexStringTryFromString(v string) (*HexString, error) {
	if strings.HasPrefix(v, "0x") {
		v = v[2:]
	}
	r, err := hex.DecodeString(v)
	if err != nil {
		return nil, err
	}
	return &HexString{value: r}, nil
}

var HexStringType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "HexString",
	Description: "The `HexString` holds hex string representation",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case HexString:
			return value.String()
		case *HexString:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewHexStringTryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewHexStringTryFromString(*value)
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
			v, err := NewHexStringTryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
