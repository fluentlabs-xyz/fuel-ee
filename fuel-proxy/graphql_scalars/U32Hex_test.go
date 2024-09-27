package graphql_scalars

import (
	"encoding/binary"
	"encoding/hex"
	"testing"
)

func TestU32HexEncodeDecode(t *testing.T) {
	var v uint32 = 12345
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, v)
	bufHex := hex.EncodeToString(buf)
	if bufHex != "00003039" {
		t.Fatalf("failed to encode v. bufHex: %s", bufHex)
	}

	vDecoded := binary.BigEndian.Uint32(buf)
	if vDecoded != v {
		t.Fatalf("v!=vDecoded: %d!=%d", v, vDecoded)
	}

	//
	u32Val := NewU32Hex(12345)
	u32ValString := u32Val.String()
	if u32ValString != "0x00003039" {
		t.Fatalf("failed to encode u32Val. u32ValString: %s", u32ValString)
	}
	u32ValDecoded, err := NewU32HexTryFromString(u32ValString)
	if err != nil {
		t.Fatalf("failed to decode u32ValString: %s", err)
	}
	if u32Val.value != u32ValDecoded.value {
		t.Fatalf("u32Val.value!=u32ValDecoded.value: %d!=%d", u32Val.value, u32ValDecoded.value)
	}
}
