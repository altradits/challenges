package piscine

import "testing"

func TestGcd(t *testing.T) {
	cases := []struct {
		name string
		a, b uint
		want uint
	}{
		{"42 10", 42, 10, 2},
		{"42 12", 42, 12, 6},
		{"14 77", 14, 77, 7},
		{"coprime", 17, 3, 1},
		{"zero a", 0, 5, 0},
		{"zero b", 5, 0, 0},
		{"equal", 6, 6, 6},
		{"one", 1, 99, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Gcd(tc.a, tc.b); got != tc.want {
				t.Errorf("Gcd(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
