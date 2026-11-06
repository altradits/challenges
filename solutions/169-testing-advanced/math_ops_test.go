package piscine

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
		{"large positive", 1000000, 1000000},
		{"large negative", -1000000, 1000000},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Abs(tc.input); got != tc.want {
				t.Errorf("Abs(%d) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exp      int
		want     int
	}{
		{"2^0", 2, 0, 1},
		{"2^1", 2, 1, 2},
		{"2^10", 2, 10, 1024},
		{"3^3", 3, 3, 27},
		{"5^0", 5, 0, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Pow(tc.base, tc.exp); got != tc.want {
				t.Errorf("Pow(%d,%d) = %d, want %d", tc.base, tc.exp, got, tc.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"two", 2, true},
		{"three", 3, true},
		{"four", 4, false},
		{"one", 1, false},
		{"zero", 0, false},
		{"negative", -7, false},
		{"large prime", 9999991, true},
		{"large composite", 9999994, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPrime(tc.input); got != tc.want {
				t.Errorf("IsPrime(%d) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
		{20, 6765},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := Fibonacci(tc.n); got != tc.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for b.Loop() {
		IsPrime(9999991)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for b.Loop() {
		Fibonacci(40)
	}
}
