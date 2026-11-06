package piscine

import "testing"

func TestCountChar(t *testing.T) {
	cases := []struct {
		name string
		s    string
		c    rune
		want int
	}{
		{"three l", "Hello World", 'l', 3},
		{"control rune", "5  balloons", 5, 0},
		{"spaces", "   ", ' ', 3},
		{"digit", "The 7 deadly sins", '7', 1},
		{"missing", "abc", 'z', 0},
		{"empty", "", 'a', 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountChar(tc.s, tc.c); got != tc.want {
				t.Errorf("CountChar(%q, %q) = %d, want %d", tc.s, tc.c, got, tc.want)
			}
		})
	}
}
