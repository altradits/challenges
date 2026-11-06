package piscine

import "testing"

func TestPrintIfNot(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"long", "abcdefz", "Invalid Input\n"},
		{"boundary", "abc", "Invalid Input\n"},
		{"empty", "", "G\n"},
		{"short", "14", "G\n"},
		{"single", "x", "G\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PrintIfNot(tc.input); got != tc.want {
				t.Errorf("PrintIfNot(%q) = %q, want %q (len < 3 -> G)", tc.input, got, tc.want)
			}
		})
	}
}
