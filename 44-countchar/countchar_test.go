package piscine

import "testing"

func TestCountChar(t *testing.T) {
	cases := []struct {
		name string
		s    string
		c    rune
		want int
	}{
		{"two", "hello", 'l', 2},
		{"zero", "hello", 'z', 0},
		{"empty", "", 'a', 0},
		{"all", "aaa", 'a', 3},
		{"unicode", "héllo", 'l', 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountChar(tc.s, tc.c); got != tc.want {
				t.Errorf("CountChar(%q, %q) = %d, want %d", tc.s, tc.c, got, tc.want)
			}
		})
	}
}
