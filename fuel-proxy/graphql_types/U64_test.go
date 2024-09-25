package graphql_types

import (
	"testing"
)

func TestU64EncodeDecodeV2(t *testing.T) {
	//
	u64Val := NewU64(18446744073709551615)
	u64ValString := u64Val.String()
	if u64ValString != "18446744073709551615" {
		t.Fatalf("failed to encode u64Val. u64ValString: %s", u64ValString)
	}
	u64ValDecoded, err := NewU64TryFromString(u64ValString)
	if err != nil {
		t.Fatalf("failed to decode u64ValString: %s", err)
	}
	if u64Val.value != u64ValDecoded.value {
		t.Fatalf("u64Val.value!=u64ValDecoded.value: %d!=%d", u64Val.value, u64ValDecoded.value)
	}
}
