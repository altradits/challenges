package piscine

import "testing"

func TestFindLastChar(t *testing.T) {
	cases := []struct {
		name string
		s    string
		c    rune
		want int
	}{
		{"last l", "hello", 'l', 3},
		{"single", "hello", 'h', 0},
		{"missing", "hello", 'z', -1},
		{"empty", "", 'a', -1},
		{"unicode byte index", "héllo", 'l', 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindLastChar(tc.s, tc.c); got != tc.want {
				t.Errorf("FindLastChar(%q, %q) = %d, want %d (byte index, -1 if missing)", tc.s, tc.c, got, tc.want)
			}
		})
	}
}
