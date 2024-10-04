package graphql_scalars

import (
	"testing"
)

func TestU32EncodeDecodeV2(t *testing.T) {
	//
	u32Val := NewU32(12345)
	u32ValString := u32Val.String()
	if u32ValString != "12345" {
		t.Fatalf("failed to encode u32Val. u32ValString: %s", u32ValString)
	}
	u32ValDecoded, err := NewU32TryFromString(u32ValString)
	if err != nil {
		t.Fatalf("failed to decode u32ValString: %s", err)
	}
	if u32Val.value != u32ValDecoded.value {
		t.Fatalf("u32Val.value!=u32ValDecoded.value: %d!=%d", u32Val.value, u32ValDecoded.value)
	}
}
