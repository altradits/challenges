package piscine

import "testing"

func TestIndex(t *testing.T) {
	cases := []struct {
		name string
		s    string
		sub  string
		want int
	}{
		{"found", "hello world", "world", 6},
		{"missing", "hello world", "xyz", -1},
		{"early", "abcabc", "b", 1},
		{"empty s", "", "a", -1},
		{"empty sub", "abc", "", 0},
		{"whole", "abc", "abc", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Index(tc.s, tc.sub); got != tc.want {
				t.Errorf("Index(%q, %q) = %d, want %d", tc.s, tc.sub, got, tc.want)
			}
		})
	}
}
