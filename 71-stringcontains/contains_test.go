package piscine

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		name string
		s    string
		sub  string
		want bool
	}{
		{"found", "hello world", "world", true},
		{"missing", "hello world", "xyz", false},
		{"empty s", "", "a", false},
		{"empty sub", "abc", "", true},
		{"both empty", "", "", true},
		{"whole", "abc", "abc", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.s, tc.sub); got != tc.want {
				t.Errorf("Contains(%q, %q) = %v, want %v", tc.s, tc.sub, got, tc.want)
			}
		})
	}
}
