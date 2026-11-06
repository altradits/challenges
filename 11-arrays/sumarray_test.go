package piscine

import "testing"

func TestSumArray(t *testing.T) {
	cases := []struct {
		name  string
		input [5]int
		want  int
	}{
		{"one to five", [5]int{1, 2, 3, 4, 5}, 15},
		{"zeros", [5]int{0, 0, 0, 0, 0}, 0},
		{"mixed", [5]int{-1, -2, 3, 4, 5}, 9},
		{"tens", [5]int{10, 20, 30, 40, 50}, 150},
		{"negatives", [5]int{-5, -5, -5, -5, -5}, -25},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumArray(tc.input); got != tc.want {
				t.Errorf("SumArray(%v) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
