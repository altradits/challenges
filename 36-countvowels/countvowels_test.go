package piscine

import "testing"

func TestCountVowels(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"hello", "hello", 2},
		{"empty", "", 0},
		{"upper", "AEIOU", 5},
		{"mixed case", "Go Language", 5},
		{"none", "rhythm", 0},
		{"all", "aeiouAEIOU", 10},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountVowels(tc.input); got != tc.want {
				t.Errorf("CountVowels(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
