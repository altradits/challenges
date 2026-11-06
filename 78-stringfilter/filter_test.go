package piscine

import (
	"testing"
	"unicode"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func TestFilter(t *testing.T) {
	cases := []struct {
		name string
		s    string
		f    func(rune) bool
		want string
	}{
		{"vowels", "hello", isVowel, "eo"},
		{"digits", "abc123", unicode.IsDigit, "123"},
		{"upper", "Go Lang", unicode.IsUpper, "GL"},
		{"empty", "", isVowel, ""},
		{"none match", "abc", unicode.IsDigit, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Filter(tc.s, tc.f); got != tc.want {
				t.Errorf("Filter(%q) = %q, want %q (keep chars where f is true)", tc.s, got, tc.want)
			}
		})
	}
}
