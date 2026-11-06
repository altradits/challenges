package piscine

import (
	"encoding/binary"
	"errors"
)

var (
	ErrShortVarint        = errors.New("short varint")
	ErrNonCanonicalVarint = errors.New("non-canonical varint")
)

func EncodeVarint(n uint64) []byte {
	switch {
	case n < 253:
		return []byte{byte(n)}
	case n <= 0xffff:
		b := make([]byte, 3)
		b[0] = 0xfd
		binary.LittleEndian.PutUint16(b[1:], uint16(n))
		return b
	case n <= 0xffffffff:
		b := make([]byte, 5)
		b[0] = 0xfe
		binary.LittleEndian.PutUint32(b[1:], uint32(n))
		return b
	default:
		b := make([]byte, 9)
		b[0] = 0xff
		binary.LittleEndian.PutUint64(b[1:], n)
		return b
	}
}

func DecodeVarint(b []byte) (uint64, int, error) {
	if len(b) == 0 {
		return 0, 0, ErrShortVarint
	}
	switch b[0] {
	case 0xfd:
		if len(b) < 3 {
			return 0, 0, ErrShortVarint
		}
		v := binary.LittleEndian.Uint16(b[1:3])
		if v < 253 {
			return 0, 0, ErrNonCanonicalVarint
		}
		return uint64(v), 3, nil
	case 0xfe:
		if len(b) < 5 {
			return 0, 0, ErrShortVarint
		}
		v := binary.LittleEndian.Uint32(b[1:5])
		if v < 0x10000 {
			return 0, 0, ErrNonCanonicalVarint
		}
		return uint64(v), 5, nil
	case 0xff:
		if len(b) < 9 {
			return 0, 0, ErrShortVarint
		}
		v := binary.LittleEndian.Uint64(b[1:9])
		if v < 0x100000000 {
			return 0, 0, ErrNonCanonicalVarint
		}
		return v, 9, nil
	default:
		return uint64(b[0]), 1, nil
	}
}
