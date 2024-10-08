package helpers

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

func RequireNoError(err error) {
	if err != nil {
		log.Fatalf("failed: %s", err)
	}
}

func RequireNotNil(e interface{}, errText string, errTextArgs ...interface{}) {
	if e == nil {
		log.Fatalf(errText, errTextArgs)
	}
}

func BytesToHexNumberString(v []byte) string {
	res := strings.TrimLeft(hex.EncodeToString(v), "0")
	if res == "" {
		return "00"
	}
	if len(res)%2 != 0 {
		return "0" + res
	}
	return res
}

func BytesToHexString(v []byte) string {
	if len(v)%2 != 0 {
		return "0" + hex.EncodeToString(v)
	}
	return hex.EncodeToString(v)
}

func BytesToHexNumberStringPrefixed(v []byte) string {
	return "0x" + BytesToHexNumberString(v)
}

func BytesToHexStringPrefixed(v []byte) string {
	return "0x" + BytesToHexString(v)
}

func Uint64ToBytesBE(v uint64, len int32) ([]byte, error) {
	if len < 8 {
		return nil, errors.New("len must be GE 8")
	}
	res := make([]byte, len)
	binary.BigEndian.PutUint64(res[len-8:], v)
	return res, nil
}

func Uint64ToBytesBEMust(v uint64, len int32) []byte {
	res, err := Uint64ToBytesBE(v, len)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func Uint32ToBytesBE(v uint32, len int32) ([]byte, error) {
	if len < 4 {
		return nil, errors.New("len must be GE 4")
	}
	res := make([]byte, len)
	binary.BigEndian.PutUint32(res[len-4:], v)
	return res, nil
}

func Uint32ToBytesBEMust(v uint32, len int32) []byte {
	res, err := Uint32ToBytesBE(v, len)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func HexStringToBinInt(vHex string) (*big.Int, error) {
	vHex = strings.TrimPrefix(vHex, "0x")
	res := new(big.Int)
	res, ok := res.SetString(vHex, 16)
	if !ok {
		return nil, errors.New(fmt.Sprintf("failed converting hex res '%s' to big int", vHex))
	}
	return res, nil
}
