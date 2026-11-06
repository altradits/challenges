package piscine

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		name string
		s    string
		sub  string
		want int
	}{
		{"chars", "hello hello", "l", 4},
		{"pairs", "hello hello", "ll", 2},
		{"missing", "hello hello", "xyz", 0},
		{"empty s", "", "a", 0},
		{"empty sub", "abc", "", 0},
		{"overlapping", "aaa", "aa", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.s, tc.sub); got != tc.want {
				t.Errorf("Count(%q, %q) = %d, want %d (overlapping matches count)", tc.s, tc.sub, got, tc.want)
			}
		})
	}
}
