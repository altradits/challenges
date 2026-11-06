package piscine

import (
	"testing"
	"unicode"
)

func TestMap(t *testing.T) {
	upper := func(r rune) rune { return unicode.ToUpper(r) }
	lower := func(r rune) rune { return unicode.ToLower(r) }
	shift := func(r rune) rune { return r + 1 }

	cases := []struct {
		name string
		s    string
		f    func(rune) rune
		want string
	}{
		{"upper", "abc", upper, "ABC"},
		{"lower", "Hello", lower, "hello"},
		{"shift", "123", shift, "234"},
		{"empty", "", upper, ""},
		{"identity", "Go!", func(r rune) rune { return r }, "Go!"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Map(tc.s, tc.f); got != tc.want {
				t.Errorf("Map(%q) = %q, want %q", tc.s, got, tc.want)
			}
		})
	}
}
