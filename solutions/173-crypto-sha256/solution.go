package piscine

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func SHA256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func SHA256d(b []byte) [32]byte {
	first := sha256.Sum256(b)
	return sha256.Sum256(first[:])
}

func HashReader(r io.Reader) ([32]byte, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return [32]byte{}, err
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out, nil
}

func Checksum4(b []byte) [4]byte {
	d := SHA256d(b)
	return [4]byte{d[0], d[1], d[2], d[3]}
}
