package piscine

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"hello", "hello", "olleh"},
		{"empty", "", ""},
		{"single", "a", "a"},
		{"two", "ab", "ba"},
		{"sentence", "Go is fun", "nuf si oG"},
		{"unicode", "héllo", "olléh"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ReverseString(tc.input); got != tc.want {
				t.Errorf("ReverseString(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
