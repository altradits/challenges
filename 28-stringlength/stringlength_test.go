package piscine

import "testing"

func TestStringLength(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"hello", "Hello", 5},
		{"empty", "", 0},
		{"go", "Go", 2},
		{"single", "a", 1},
		{"sentence", "Go is fun", 9},
		{"unicode", "héllo", 5},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := StringLength(tc.input); got != tc.want {
				t.Errorf("StringLength(%q) = %d, want %d (count runes, no len())", tc.input, got, tc.want)
			}
		})
	}
}
