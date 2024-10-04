package common

import (
	// "fmt"
	// "math/big"
	// "strings"

	// proto "github.com/Ankr-network/multirpc-proto-contract"
	"github.com/pkg/errors"
	// "github.com/shopspring/decimal"
)

func PanicOnError(err error, msg string) {
	if err != nil {
		panic(errors.WithMessage(err, msg))
	}
}

// func HumanReadableValue(value string, decimals int32) string {
// 	bigValue, _ := new(big.Float).SetPrec(200).SetMode(big.ToZero).SetString(value)
// 	bigPower := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
// 	bigHumanValue := new(big.Float).SetPrec(200).SetMode(big.ToZero).Quo(bigValue, new(big.Float).SetInt(bigPower))
// 	textValue := bigHumanValue.Text('f', 18)
// 	return strings.TrimRight(strings.TrimRight(textValue, "0"), ".")
// }
