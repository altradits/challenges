package piscine

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
)

const MaxPayload = 32 << 20

func doubleSHA190(b []byte) [32]byte {
	first := sha256.Sum256(b)
	return sha256.Sum256(first[:])
}

func EncodeMessage(magic [4]byte, command string, payload []byte) ([]byte, error) {
	if len(command) == 0 || len(command) > 12 {
		return nil, fmt.Errorf("command must be 1-12 chars, got %d", len(command))
	}
	for i := 0; i < len(command); i++ {
		if command[i] < 0x20 || command[i] > 0x7e {
			return nil, fmt.Errorf("command must be printable ASCII")
		}
	}
	if len(payload) > MaxPayload {
		return nil, fmt.Errorf("payload %d bytes exceeds %d cap", len(payload), MaxPayload)
	}
	out := make([]byte, 0, 24+len(payload))
	out = append(out, magic[:]...)
	cmd := make([]byte, 12)
	copy(cmd, command)
	out = append(out, cmd...)
	var ln [4]byte
	binary.LittleEndian.PutUint32(ln[:], uint32(len(payload)))
	out = append(out, ln[:]...)
	sum := doubleSHA190(payload)
	out = append(out, sum[0], sum[1], sum[2], sum[3])
	out = append(out, payload...)
	return out, nil
}

func DecodeMessage(r io.Reader, magic [4]byte) (string, []byte, error) {
	var hdr [24]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return "", nil, fmt.Errorf("reading header: %w", err)
	}
	if !bytes.Equal(hdr[0:4], magic[:]) {
		return "", nil, fmt.Errorf("bad magic %x, want %x", hdr[0:4], magic)
	}
	length := binary.LittleEndian.Uint32(hdr[16:20])
	if length > MaxPayload {
		return "", nil, fmt.Errorf("payload length %d exceeds %d cap", length, MaxPayload)
	}
	payload := make([]byte, length)
	if _, err := io.ReadFull(r, payload); err != nil {
		return "", nil, fmt.Errorf("reading payload: %w", err)
	}
	sum := doubleSHA190(payload)
	if !bytes.Equal(hdr[20:24], sum[0:4]) {
		return "", nil, fmt.Errorf("bad checksum")
	}
	cmd := hdr[4:16]
	if i := bytes.IndexByte(cmd, 0); i >= 0 {
		cmd = cmd[:i]
	}
	if len(cmd) == 0 {
		return "", nil, fmt.Errorf("empty command")
	}
	return string(cmd), payload, nil
}
