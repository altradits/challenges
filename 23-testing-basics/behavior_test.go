package piscine

import "testing"

// These hidden tests check behavior. Your own solution_test.go (table-driven
// tests plus a benchmark, as the README requires) runs alongside them.

func TestAbs(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Abs(tc.input); got != tc.want {
				t.Errorf("Abs(%d) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestPalindrome(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"odd", "racecar", true},
		{"even", "abba", true},
		{"single", "x", true},
		{"empty", "", true},
		{"not", "hello", false},
		{"case sensitive", "Abba", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Palindrome(tc.input); got != tc.want {
				t.Errorf("Palindrome(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  string
	}{
		{"three", 3, "Fizz"},
		{"five", 5, "Buzz"},
		{"fifteen", 15, "FizzBuzz"},
		{"plain", 7, "7"},
		{"one", 1, "1"},
		{"thirty", 30, "FizzBuzz"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FizzBuzz(tc.input); got != tc.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
