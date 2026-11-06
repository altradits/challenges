package piscine

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"racecar", "racecar", true},
		{"hello", "Hello", false},
		{"famous", "A man a plan a canal Panama", true},
		{"empty", "", true},
		{"single", "x", true},
		{"even", "abba", true},
		{"digits", "12321", true},
		{"not quite", "abca", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.input); got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
