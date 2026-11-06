package piscine

import "testing"

func TestReverseStrCap(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"sentence", "First SMALL TesT", "firsT smalL tesT"},
		{"hello", "Hello", "hellO"},
		{"empty", "", ""},
		{"single char", "a", "A"},
		{"extra spaces collapse", "a  b", "A B"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ReverseStrCap(tc.input); got != tc.want {
				t.Errorf("ReverseStrCap(%q) = %q, want %q (lowercase word, uppercase last letter)", tc.input, got, tc.want)
			}
		})
	}
}
