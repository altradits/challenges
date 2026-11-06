package piscine

import "testing"

func TestSearchReplace(t *testing.T) {
	cases := []struct {
		name      string
		text, old string
		new       string
		want      string
	}{
		{"first only", "hello world hello", "hello", "hi", "hi world hello"},
		{"missing", "hello", "z", "q", "hello"},
		{"empty text", "", "a", "b", ""},
		{"whole", "aaa", "aaa", "b", "b"},
		{"grow", "ab", "b", "bcd", "abcd"},
		{"shrink", "abcd", "bc", "", "ad"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SearchReplace(tc.text, tc.old, tc.new); got != tc.want {
				t.Errorf("SearchReplace(%q, %q, %q) = %q, want %q (first occurrence only)",
					tc.text, tc.old, tc.new, got, tc.want)
			}
		})
	}
}
