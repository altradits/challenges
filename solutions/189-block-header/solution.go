package piscine

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type Header struct {
	Version    int32
	PrevBlock  [32]byte
	MerkleRoot [32]byte
	Timestamp  uint32
	Bits       uint32
	Nonce      uint32
}

func ParseHeader(b []byte) (Header, error) {
	var h Header
	if len(b) != 80 {
		return h, fmt.Errorf("header must be 80 bytes, got %d", len(b))
	}
	h.Version = int32(binary.LittleEndian.Uint32(b[0:4]))
	copy(h.PrevBlock[:], b[4:36])
	copy(h.MerkleRoot[:], b[36:68])
	h.Timestamp = binary.LittleEndian.Uint32(b[68:72])
	h.Bits = binary.LittleEndian.Uint32(b[72:76])
	h.Nonce = binary.LittleEndian.Uint32(b[76:80])
	return h, nil
}

func SerializeHeader(h Header) []byte {
	b := make([]byte, 80)
	binary.LittleEndian.PutUint32(b[0:4], uint32(h.Version))
	copy(b[4:36], h.PrevBlock[:])
	copy(b[36:68], h.MerkleRoot[:])
	binary.LittleEndian.PutUint32(b[68:72], h.Timestamp)
	binary.LittleEndian.PutUint32(b[72:76], h.Bits)
	binary.LittleEndian.PutUint32(b[76:80], h.Nonce)
	return b
}

func HeaderHash(h Header) [32]byte {
	ser := SerializeHeader(h)
	first := sha256.Sum256(ser)
	return sha256.Sum256(first[:])
}
