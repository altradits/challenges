package piscine

import (
	"encoding/binary"
	"fmt"
)

func EncodeUint32LE(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

// DecodeUint32LE reads the first 4 bytes; missing bytes count as zero.
func DecodeUint32LE(b []byte) uint32 {
	var tmp [4]byte
	copy(tmp[:], b)
	return binary.LittleEndian.Uint32(tmp[:])
}

func EncodeUint64BE(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// DecodeUint64BE reads the first 8 bytes; missing bytes count as zero.
func DecodeUint64BE(b []byte) uint64 {
	var tmp [8]byte
	copy(tmp[:], b)
	return binary.BigEndian.Uint64(tmp[:])
}

func PackRecord(id uint32, score uint16) []byte {
	b := make([]byte, 6)
	binary.LittleEndian.PutUint32(b[0:4], id)
	binary.LittleEndian.PutUint16(b[4:6], score)
	return b
}

func UnpackRecord(b []byte) (uint32, uint16, error) {
	if len(b) != 6 {
		return 0, 0, fmt.Errorf("record must be 6 bytes, got %d", len(b))
	}
	return binary.LittleEndian.Uint32(b[0:4]),
		binary.LittleEndian.Uint16(b[4:6]), nil
}
