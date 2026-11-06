package piscine

import (
	"errors"
	"strings"
	"testing"
)

func TestSHA256Hex(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
	}
	for _, tc := range cases {
		if got := SHA256Hex(tc.input); got != tc.want {
			t.Errorf("SHA256Hex(%q) =\n%s\nwant\n%s", tc.input, got, tc.want)
		}
	}
}

func TestSHA256d(t *testing.T) {
	got := SHA256d([]byte("hello"))
	want := "9595c9df90075148eb06860365df33584b75bff782a510c6cd4883a419833d50"
	if SHA256HexBytes(got[:]) != want {
		t.Errorf("SHA256d(hello) = %x, want %s", got, want)
	}
	empty := SHA256d(nil)
	if SHA256HexBytes(empty[:]) != "5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456" {
		t.Errorf("SHA256d(empty) = %x", empty)
	}
}

func SHA256HexBytes(b []byte) string {
	const digits = "0123456789abcdef"
	out := make([]byte, 0, len(b)*2)
	for _, c := range b {
		out = append(out, digits[c>>4], digits[c&0xf])
	}
	return string(out)
}

func TestHashReader(t *testing.T) {
	got, err := HashReader(strings.NewReader("abc"))
	if err != nil {
		t.Fatalf("HashReader error: %v", err)
	}
	if SHA256HexBytes(got[:]) != "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" {
		t.Errorf("HashReader(abc) = %x", got)
	}
	if _, err := HashReader(failReader{}); err == nil {
		t.Error("HashReader should propagate read errors")
	}
}

type failReader struct{}

func (failReader) Read([]byte) (int, error) { return 0, errors.New("boom") }

func TestChecksum4(t *testing.T) {
	if got := Checksum4([]byte("hello")); got != [4]byte{0x95, 0x95, 0xc9, 0xdf} {
		t.Errorf("Checksum4 = %x", got)
	}
}
