package graphql_scalars

import (
	"encoding/hex"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strings"
)

type Address struct {
	value [32]byte
}

func (b *Address) String() string {
	return "0x" + hex.EncodeToString(b.value[:])
}

func NewAddress(v [32]byte) *Address {
	return &Address{
		value: v,
	}
}

func NewAddressTryFromString(v string) (*Address, error) {
	if strings.HasPrefix(v, "0x") {
		v = v[2:]
	}
	output, err := hex.DecodeString(v)
	if err != nil {
		return nil, err
	}
	if len(output) != 32 {
		return nil, fmt.Errorf("output must be 32 bytes len")
	}
	return NewAddress(([32]byte)(output)), nil
}

func NewAddressTryFromStringOrPanic(v string) *Address {
	output, err := NewAddressTryFromString(v)
	if err != nil {
		panic(fmt.Sprintf("failed to convert string into Address: %s", err))
	}
	return output
}

var AddressType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Address",
	Description: "The `AddressType` holds fixed 32 bytes array",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case Address:
			return value.String()
		case *Address:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewAddressTryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewAddressTryFromString(*value)
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
			v, err := NewAddressTryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
