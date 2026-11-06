package piscine

import (
	"bytes"
	"testing"
)

func TestUint32LE(t *testing.T) {
	got := EncodeUint32LE(0x01020304)
	want := []byte{0x04, 0x03, 0x02, 0x01}
	if !bytes.Equal(got, want) {
		t.Errorf("EncodeUint32LE = %x, want %x", got, want)
	}
	if v := DecodeUint32LE(got); v != 0x01020304 {
		t.Errorf("DecodeUint32LE = %#x, want 0x01020304", v)
	}
	if v := DecodeUint32LE([]byte{0x01}); v != 1 {
		t.Errorf("DecodeUint32LE short = %d, want 1 (zero-padded)", v)
	}
	if v := DecodeUint32LE(nil); v != 0 {
		t.Errorf("DecodeUint32LE(nil) = %d, want 0", v)
	}
}

func TestUint64BE(t *testing.T) {
	got := EncodeUint64BE(1)
	want := []byte{0, 0, 0, 0, 0, 0, 0, 1}
	if !bytes.Equal(got, want) {
		t.Errorf("EncodeUint64BE = %x, want %x", got, want)
	}
	if v := DecodeUint64BE(got); v != 1 {
		t.Errorf("DecodeUint64BE = %d, want 1", v)
	}
	if v := DecodeUint64BE([]byte{0x01}); v != 0x0100000000000000 {
		t.Errorf("DecodeUint64BE short = %#x, want zero-padded", v)
	}
}

func TestPackUnpackRecord(t *testing.T) {
	got := PackRecord(0x01020304, 0x0506)
	want := []byte{0x04, 0x03, 0x02, 0x01, 0x06, 0x05}
	if !bytes.Equal(got, want) {
		t.Fatalf("PackRecord = %x, want %x", got, want)
	}
	id, score, err := UnpackRecord(got)
	if err != nil {
		t.Fatalf("UnpackRecord error: %v", err)
	}
	if id != 0x01020304 || score != 0x0506 {
		t.Errorf("UnpackRecord = (%#x, %#x)", id, score)
	}
	if _, _, err := UnpackRecord([]byte{1, 2, 3}); err == nil {
		t.Error("UnpackRecord short input should error")
	}
}
