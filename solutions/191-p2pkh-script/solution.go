package piscine

import (
	"bytes"
	"fmt"
)

func BuildP2PKH(pubkeyHash []byte) ([]byte, error) {
	if len(pubkeyHash) != 20 {
		return nil, fmt.Errorf("pubkey hash must be 20 bytes, got %d", len(pubkeyHash))
	}
	out := []byte{0x76, 0xa9, 0x14}
	out = append(out, pubkeyHash...)
	return append(out, 0x88, 0xac), nil
}

func ParseP2PKH(script []byte) ([]byte, error) {
	if !IsP2PKH(script) {
		return nil, fmt.Errorf("not a P2PKH script")
	}
	return bytes.Clone(script[3:23]), nil
}

func IsP2PKH(script []byte) bool {
	return len(script) == 25 &&
		script[0] == 0x76 && script[1] == 0xa9 && script[2] == 0x14 &&
		script[23] == 0x88 && script[24] == 0xac
}
