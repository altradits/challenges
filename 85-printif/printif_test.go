package piscine

import "testing"

func TestPrintIf(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"long", "abcdefz", "G\n"},
		{"boundary", "abc", "G\n"},
		{"empty", "", "G\n"},
		{"short", "14", "Invalid Input\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PrintIf(tc.input); got != tc.want {
				t.Errorf("PrintIf(%q) = %q, want %q (len 0 or >= 3 -> G)", tc.input, got, tc.want)
			}
		})
	}
}
