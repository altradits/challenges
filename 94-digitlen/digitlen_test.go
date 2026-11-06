package piscine

import "testing"

func TestDigitLen(t *testing.T) {
	cases := []struct {
		name string
		n    int
		base int
		want int
	}{
		{"decimal", 100, 10, 3},
		{"binary", 100, 2, 7},
		{"hex negative", -100, 16, 2},
		{"bad base", 100, -1, -1},
		{"base one", 100, 1, -1},
		{"base 37", 100, 37, -1},
		{"zero", 0, 10, 1},
		{"single", 5, 10, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := DigitLen(tc.n, tc.base); got != tc.want {
				t.Errorf("DigitLen(%d, %d) = %d, want %d", tc.n, tc.base, got, tc.want)
			}
		})
	}
}
