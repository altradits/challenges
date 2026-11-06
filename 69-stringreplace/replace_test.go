package piscine

import "testing"

func TestReplace(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		old, new string
		want     string
	}{
		{"chars", "hello world", "l", "L", "heLLo worLd"},
		{"words", "foo bar foo", "foo", "baz", "baz bar baz"},
		{"missing", "abc", "x", "y", "abc"},
		{"empty s", "", "a", "b", ""},
		{"empty old", "abc", "", "x", "abc"},
		{"delete", "banana", "an", "", "ba"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Replace(tc.s, tc.old, tc.new); got != tc.want {
				t.Errorf("Replace(%q, %q, %q) = %q, want %q", tc.s, tc.old, tc.new, got, tc.want)
			}
		})
	}
}
