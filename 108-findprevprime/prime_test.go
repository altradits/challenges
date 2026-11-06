package piscine

import "testing"

func TestFindPrevPrime(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"prime stays", 5, 5},
		{"four", 4, 3},
		{"two", 2, 2},
		{"one", 1, 0},
		{"zero", 0, 0},
		{"negative", -7, 0},
		{"ten", 10, 7},
		{"large", 100, 97},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindPrevPrime(tc.input); got != tc.want {
				t.Errorf("FindPrevPrime(%d) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
