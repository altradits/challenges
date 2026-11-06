package piscine

import "testing"

func TestCountRepeats(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"ll group", "hello", 1},
		{"two groups", "helloo", 2},
		{"three groups", "heelloo", 3},
		{"none", "abc", 0},
		{"empty", "", 0},
		{"triple", "aaa", 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountRepeats(tc.input); got != tc.want {
				t.Errorf("CountRepeats(%q) = %d, want %d (count groups of adjacent repeats)", tc.input, got, tc.want)
			}
		})
	}
}
