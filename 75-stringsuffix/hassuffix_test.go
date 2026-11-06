package piscine

import "testing"

func TestHasSuffix(t *testing.T) {
	cases := []struct {
		name string
		s    string
		suf  string
		want bool
	}{
		{"yes", "hello world", "world", true},
		{"no", "hello world", "hello", false},
		{"both empty", "", "", true},
		{"longer suffix", "abc", "abcd", false},
		{"empty suffix", "abc", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasSuffix(tc.s, tc.suf); got != tc.want {
				t.Errorf("HasSuffix(%q, %q) = %v, want %v", tc.s, tc.suf, got, tc.want)
			}
		})
	}
}
