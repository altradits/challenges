package piscine

import "testing"

func TestHashCode(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"A", "A", "B"},
		{"AB", "AB", "CD"},
		{"BAC", "BAC", "EDF"},
		{"hello", "Hello World", "Spwwz+bz}wo"},
		{"empty", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := HashCode(tc.input); got != tc.want {
				t.Errorf("HashCode(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
