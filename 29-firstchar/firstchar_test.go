package piscine

import "testing"

func TestFirstChar(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"hello", "Hello", "H"},
		{"single", "a", "a"},
		{"empty", "", ""},
		{"unicode", "世界", "世"},
		{"leading space", " hi", " "},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstChar(tc.input); got != tc.want {
				t.Errorf("FirstChar(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
