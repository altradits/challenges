package piscine

import "testing"

func TestHasPrefix(t *testing.T) {
	cases := []struct {
		name string
		s    string
		pre  string
		want bool
	}{
		{"yes", "hello world", "hello", true},
		{"no", "hello world", "world", false},
		{"both empty", "", "", true},
		{"longer prefix", "abc", "abcd", false},
		{"empty prefix", "abc", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasPrefix(tc.s, tc.pre); got != tc.want {
				t.Errorf("HasPrefix(%q, %q) = %v, want %v", tc.s, tc.pre, got, tc.want)
			}
		})
	}
}
