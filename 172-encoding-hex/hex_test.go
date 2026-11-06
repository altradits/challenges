package piscine

import (
	"bytes"
	"testing"
)

func TestToHexFromHex(t *testing.T) {
	if got := ToHex([]byte{0xde, 0xad}); got != "dead" {
		t.Errorf("ToHex = %q, want dead", got)
	}
	if got := ToHex(nil); got != "" {
		t.Errorf("ToHex(nil) = %q, want \"\"", got)
	}
	got, err := FromHex("dead")
	if err != nil {
		t.Fatalf("FromHex error: %v", err)
	}
	if !bytes.Equal(got, []byte{0xde, 0xad}) {
		t.Errorf("FromHex = %x", got)
	}
	if _, err := FromHex("zz"); err == nil {
		t.Error("FromHex(\"zz\") should error")
	}
	if _, err := FromHex("abc"); err == nil {
		t.Error("FromHex odd length should error")
	}
	got, err = FromHex("")
	if err != nil || len(got) != 0 {
		t.Errorf("FromHex(\"\") = %x, %v", got, err)
	}
}

func TestIsHex(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"0123abcdef", true},
		{"ABCDEF", true},
		{"", true},
		{"zz", false},
		{"abc", false},
		{"0x12", false},
	}
	for _, tc := range cases {
		if got := IsHex(tc.input); got != tc.want {
			t.Errorf("IsHex(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestReverseHex(t *testing.T) {
	got, err := ReverseHex("0102ff")
	if err != nil {
		t.Fatalf("ReverseHex error: %v", err)
	}
	if got != "ff0201" {
		t.Errorf("ReverseHex = %q, want ff0201", got)
	}
	if _, err := ReverseHex("zz"); err == nil {
		t.Error("ReverseHex invalid should error")
	}
	got, err = ReverseHex("")
	if err != nil || got != "" {
		t.Errorf("ReverseHex(\"\") = %q, %v", got, err)
	}
}
