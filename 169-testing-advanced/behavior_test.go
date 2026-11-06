package piscine

import "testing"

// Hidden behavior tests for the MathOps functions. Your own math_ops_test.go
// (table-driven tests plus benchmarks, as the README requires) runs alongside.

func TestAbsHidden(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{5, 5}, {-5, 5}, {0, 0},
	}
	for _, tc := range cases {
		if got := Abs(tc.in); got != tc.want {
			t.Errorf("Abs(%d) = %d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestPowHidden(t *testing.T) {
	cases := []struct {
		base, exp int
		want      int
	}{
		{2, 10, 1024}, {5, 0, 1}, {3, 3, 27}, {10, 2, 100},
	}
	for _, tc := range cases {
		if got := Pow(tc.base, tc.exp); got != tc.want {
			t.Errorf("Pow(%d, %d) = %d, want %d", tc.base, tc.exp, got, tc.want)
		}
	}
}

func TestIsPrimeHidden(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 97}
	for _, n := range primes {
		if !IsPrime(n) {
			t.Errorf("IsPrime(%d) = false, want true", n)
		}
	}
	nonPrimes := []int{0, 1, 4, 9, 15, 100, -7}
	for _, n := range nonPrimes {
		if IsPrime(n) {
			t.Errorf("IsPrime(%d) = true, want false", n)
		}
	}
}

func TestFibonacciHidden(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {10, 55}, {20, 6765},
	}
	for _, tc := range cases {
		if got := Fibonacci(tc.in); got != tc.want {
			t.Errorf("Fibonacci(%d) = %d, want %d", tc.in, got, tc.want)
		}
	}
}
