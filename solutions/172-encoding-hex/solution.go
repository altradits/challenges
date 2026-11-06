package piscine

import "encoding/hex"

func ToHex(b []byte) string {
	return hex.EncodeToString(b)
}

func FromHex(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func IsHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

func ReverseHex(s string) (string, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return hex.EncodeToString(b), nil
}
