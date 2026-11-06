package piscine

import (
	"bytes"
	"testing"
)

var testHash = []byte{
	0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99,
	0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33,
}

func TestBuildP2PKH(t *testing.T) {
	script, err := BuildP2PKH(testHash)
	if err != nil {
		t.Fatalf("BuildP2PKH error: %v", err)
	}
	want := append([]byte{0x76, 0xa9, 0x14}, testHash...)
	want = append(want, 0x88, 0xac)
	if !bytes.Equal(script, want) {
		t.Errorf("BuildP2PKH = %x\nwant    %x", script, want)
	}
	if _, err := BuildP2PKH(testHash[:19]); err == nil {
		t.Error("19-byte hash should error")
	}
	if _, err := BuildP2PKH(nil); err == nil {
		t.Error("nil hash should error")
	}
}

func TestParseP2PKH(t *testing.T) {
	script, err := BuildP2PKH(testHash)
	if err != nil {
		t.Fatalf("BuildP2PKH error: %v", err)
	}
	got, err := ParseP2PKH(script)
	if err != nil {
		t.Fatalf("ParseP2PKH error: %v", err)
	}
	if !bytes.Equal(got, testHash) {
		t.Errorf("ParseP2PKH = %x", got)
	}
	// Must be a copy: mutating it leaves the script intact.
	got[0] ^= 0xff
	if script[3] != testHash[0] {
		t.Error("ParseP2PKH returned an alias, not a copy")
	}
	bad := append(bytes.Clone(script), 0x00)
	for _, b := range [][]byte{nil, {}, script[:24], bad} {
		if _, err := ParseP2PKH(b); err == nil {
			t.Errorf("ParseP2PKH(%x) should error", b)
		}
	}
	tampered := bytes.Clone(script)
	tampered[24] = 0x00
	if _, err := ParseP2PKH(tampered); err == nil {
		t.Error("tampered opcode should error")
	}
}

func TestIsP2PKH(t *testing.T) {
	script, _ := BuildP2PKH(testHash)
	if !IsP2PKH(script) {
		t.Error("IsP2PKH(valid) = false")
	}
	if IsP2PKH(nil) || IsP2PKH([]byte{0x76}) {
		t.Error("IsP2PKH should reject short input")
	}
	not := bytes.Clone(script)
	not[0] = 0x00
	if IsP2PKH(not) {
		t.Error("IsP2PKH should reject wrong opcode")
	}
}
