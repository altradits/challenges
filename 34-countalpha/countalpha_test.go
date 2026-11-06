package piscine

import "testing"

func TestCountAlpha(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"hello world", "Hello world", 10},
		{"spaced", "H e l l o", 5},
		{"digits mixed", "H1e2l3l4o", 5},
		{"empty", "", 0},
		{"no letters", "123 !@#", 0},
		{"all letters", "abcXYZ", 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountAlpha(tc.input); got != tc.want {
				t.Errorf("CountAlpha(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
