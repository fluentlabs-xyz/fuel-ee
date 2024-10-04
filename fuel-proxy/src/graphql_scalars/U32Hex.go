package graphql_scalars

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"strings"
)

type U32Hex struct {
	value uint32
}

func (b *U32Hex) String() string {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, b.value)
	return "0x" + hex.EncodeToString(buf)
}

func NewU32Hex(v uint32) *U32Hex {
	return &U32Hex{
		value: v,
	}
}

func NewU32HexTryFromString(v string) (*U32Hex, error) {
	if strings.HasPrefix(v, "0x") {
		v = v[2:]
	}
	output, err := hex.DecodeString(v)
	if err != nil {
		return nil, err
	}
	if len(output) != 4 {
		return nil, fmt.Errorf("output must be 32 bytes len")
	}
	r := binary.BigEndian.Uint32(output)
	return NewU32Hex(r), nil
}

var U32HexType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "U32Hex",
	Description: "The `U32Hex` holds unsigned 32 bit value",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case U32Hex:
			return value.String()
		case *U32Hex:
			v := *value
			return v.String()
		default:
			return nil
		}
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, err := NewU32HexTryFromString(value)
			if err != nil {
				return nil
			}
			return v
		case *string:
			v, err := NewU32HexTryFromString(*value)
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
			v, err := NewU32HexTryFromString(valueAST.Value)
			if err != nil {
				return nil
			}
			return v
		default:
			return nil
		}
	},
})
