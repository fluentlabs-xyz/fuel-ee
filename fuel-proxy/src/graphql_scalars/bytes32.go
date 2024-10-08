package graphql_scalars

import (
	"encoding/hex"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strings"
)

type Bytes32 struct {
	value [32]byte
}

func (b *Bytes32) String() string {
	return "0x" + hex.EncodeToString(b.value[:])
}

func (b *Bytes32) Val() [32]byte {
	return b.value
}

func NewBytes32(v [32]byte) *Bytes32 {
	return &Bytes32{
		value: v,
	}
}

func NewBytes32TryFromString(v string) (*Bytes32, error) {
	if strings.HasPrefix(v, "0x") {
		v = v[2:]
	}
	output, err := hex.DecodeString(v)
	if err != nil {
		return nil, err
	}
	if len(output) > 32 {
		return nil, fmt.Errorf("output must be <=32 bytes len")
	}
	if len(output) == 32 {
		return NewBytes32(([32]byte)(output)), nil
	}
	res := [32]byte{}
	copy(res[len(res)-len(output):], output)
	return NewBytes32(res), nil
}

func NewBytes32TryFromInterface(v interface{}) (*Bytes32, error) {
	res, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("can create bytes32 only from the string")
	}
	return NewBytes32TryFromString(res)
}

func NewBytes32TryFromStringOrPanic(v string) *Bytes32 {
	output, err := NewBytes32TryFromString(v)
	if err != nil {
		panic(fmt.Sprintf("failed to convert string into Bytes32: %s", err))
	}
	return output
}

var Bytes32Type = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Bytes32",
	Description: "The `Bytes32Type` holds fixed 32 bytes array",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case Bytes32:
			return value.String()
		case *Bytes32:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewBytes32TryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewBytes32TryFromString(*value)
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
			v, err := NewBytes32TryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
