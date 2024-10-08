package graphql_scalars

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/helpers"
	"testing"
)

func TestHexToBigInt(t *testing.T) {
	valHexStr := "9"
	valExpected := uint64(9)
	amount, err := helpers.HexStringToBinInt(valHexStr)
	if err != nil {
		t.Fatalf("failed to convert hex to big int: %s", err)
	}
	if amount.Uint64() != valExpected {
		t.Fatalf("value doesnt match: %d!=%d", amount.Uint64(), valExpected)
	}

	valHexStr = "09"
	valExpected = uint64(9)
	amount, err = helpers.HexStringToBinInt(valHexStr)
	if err != nil {
		t.Fatalf("failed to convert hex to big int: %s", err)
	}
	if amount.Uint64() != valExpected {
		t.Fatalf("value doesnt match: %d!=%d", amount.Uint64(), valExpected)
	}

	valHexStr = "0x9"
	valExpected = uint64(9)
	amount, err = helpers.HexStringToBinInt(valHexStr)
	if err != nil {
		t.Fatalf("failed to convert hex to big int: %s", err)
	}
	if amount.Uint64() != valExpected {
		t.Fatalf("value doesnt match: %d!=%d", amount.Uint64(), valExpected)
	}

	valHexStr = "0x09"
	valExpected = uint64(9)
	amount, err = helpers.HexStringToBinInt(valHexStr)
	if err != nil {
		t.Fatalf("failed to convert hex to big int: %s", err)
	}
	if amount.Uint64() != valExpected {
		t.Fatalf("value doesnt match: %d!=%d", amount.Uint64(), valExpected)
	}
}
