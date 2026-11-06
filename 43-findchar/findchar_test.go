package piscine

import "testing"

func TestFindChar(t *testing.T) {
	cases := []struct {
		name string
		s    string
		c    rune
		want int
	}{
		{"found", "hello", 'l', 2},
		{"first", "hello", 'h', 0},
		{"missing", "hello", 'z', -1},
		{"empty", "", 'a', -1},
		{"unicode byte index", "héllo", 'l', 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindChar(tc.s, tc.c); got != tc.want {
				t.Errorf("FindChar(%q, %q) = %d, want %d (byte index, -1 if missing)", tc.s, tc.c, got, tc.want)
			}
		})
	}
}
