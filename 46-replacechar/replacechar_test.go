package piscine

import "testing"

func TestReplaceChar(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		old, new rune
		want     string
	}{
		{"two l", "Hello", 'l', 'L', "HeLLo"},
		{"no match", "Hello", 'z', 'X', "Hello"},
		{"empty", "", 'a', 'b', ""},
		{"all", "aaaa", 'a', 'b', "bbbb"},
		{"unicode", "héllo", 'l', 'r', "hérro"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ReplaceChar(tc.s, tc.old, tc.new); got != tc.want {
				t.Errorf("ReplaceChar(%q, %q, %q) = %q, want %q", tc.s, tc.old, tc.new, got, tc.want)
			}
		})
	}
}
