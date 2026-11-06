package piscine

import (
	"bytes"
	"testing"
)

var testMagic = [4]byte{0xf9, 0xbe, 0xb4, 0xd9}

func TestMessageRoundTrip(t *testing.T) {
	payload := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	msg, err := EncodeMessage(testMagic, "ping", payload)
	if err != nil {
		t.Fatalf("EncodeMessage error: %v", err)
	}
	if len(msg) != 24+len(payload) {
		t.Fatalf("message length = %d, want %d", len(msg), 24+len(payload))
	}
	cmd, got, err := DecodeMessage(bytes.NewReader(msg), testMagic)
	if err != nil {
		t.Fatalf("DecodeMessage error: %v", err)
	}
	if cmd != "ping" {
		t.Errorf("command = %q, want ping", cmd)
	}
	if !bytes.Equal(got, payload) {
		t.Errorf("payload = %x, want %x", got, payload)
	}
}

func TestMessageHeaderLayout(t *testing.T) {
	msg, err := EncodeMessage(testMagic, "version", []byte{0xaa})
	if err != nil {
		t.Fatalf("EncodeMessage error: %v", err)
	}
	if !bytes.Equal(msg[0:4], testMagic[:]) {
		t.Errorf("magic = %x", msg[0:4])
	}
	if !bytes.Equal(msg[4:16], []byte("version\x00\x00\x00\x00\x00")) {
		t.Errorf("command field = %q", msg[4:16])
	}
	if !bytes.Equal(msg[16:20], []byte{1, 0, 0, 0}) {
		t.Errorf("length = %x, want 01000000", msg[16:20])
	}
}

func TestEncodeMessageRejects(t *testing.T) {
	if _, err := EncodeMessage(testMagic, "", []byte{1}); err == nil {
		t.Error("empty command should error")
	}
	if _, err := EncodeMessage(testMagic, "1234567890123", []byte{1}); err == nil {
		t.Error("13-char command should error")
	}
	big := make([]byte, MaxPayload+1)
	if _, err := EncodeMessage(testMagic, "tx", big); err == nil {
		t.Error("oversize payload should error")
	}
}

func TestDecodeMessageRejects(t *testing.T) {
	good, err := EncodeMessage(testMagic, "ping", []byte{9, 9, 9})
	if err != nil {
		t.Fatalf("EncodeMessage error: %v", err)
	}
	corrupt := func(i int) []byte {
		b := bytes.Clone(good)
		b[i] ^= 0xff
		return b
	}
	cases := []struct {
		name string
		msg  []byte
	}{
		{"bad magic", corrupt(0)},
		{"bad checksum", corrupt(21)},
		{"corrupt payload", corrupt(25)},
		{"truncated header", good[:10]},
		{"truncated payload", good[:26]},
		{"empty", nil},
	}
	for _, tc := range cases {
		if _, _, err := DecodeMessage(bytes.NewReader(tc.msg), testMagic); err == nil {
			t.Errorf("%s: expected error", tc.name)
		}
	}
	// Wrong network magic.
	if _, _, err := DecodeMessage(bytes.NewReader(good), [4]byte{1, 2, 3, 4}); err == nil {
		t.Error("wrong magic: expected error")
	}
	// Lying length must error without hanging or huge alloc.
	var hdr bytes.Buffer
	hdr.Write(testMagic[:])
	hdr.Write([]byte("block\x00\x00\x00\x00\x00\x00\x00"))
	hdr.Write([]byte{0xff, 0xff, 0xff, 0xff})
	hdr.Write([]byte{0, 0, 0, 0})
	if _, _, err := DecodeMessage(&hdr, testMagic); err == nil {
		t.Error("oversize length: expected error")
	}
}
