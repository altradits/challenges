package piscine

import "testing"

func TestWeAreUnique(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 string
		want   int
	}{
		{"foo boo", "foo", "boo", 2},
		{"both empty", "", "", -1},
		{"disjoint", "abc", "def", 6},
		{"same", "abc", "abc", 0},
		{"one empty", "a", "", 1},
		{"subset", "abc", "ab", 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WeAreUnique(tc.s1, tc.s2); got != tc.want {
				t.Errorf("WeAreUnique(%q, %q) = %d, want %d (chars in exactly one string)", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
