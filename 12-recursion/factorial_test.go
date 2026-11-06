package piscine

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct {
		name  string
		input uint
		want  uint
	}{
		{"zero", 0, 1},
		{"one", 1, 1},
		{"two", 2, 2},
		{"five", 5, 120},
		{"ten", 10, 3628800},
		{"twelve", 12, 479001600},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Factorial(tc.input); got != tc.want {
				t.Errorf("Factorial(%d) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
