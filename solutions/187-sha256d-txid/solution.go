package piscine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func DoubleSHA(b []byte) [32]byte {
	first := sha256.Sum256(b)
	return sha256.Sum256(first[:])
}

func TxidDisplay(hash [32]byte) string {
	rev := hash
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return hex.EncodeToString(rev[:])
}

func TxidParse(s string) ([32]byte, error) {
	var out [32]byte
	if len(s) != 64 {
		return out, fmt.Errorf("txid must be 64 hex chars, got %d", len(s))
	}
	raw, err := hex.DecodeString(s)
	if err != nil {
		return out, err
	}
	copy(out[:], raw)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out, nil
}
