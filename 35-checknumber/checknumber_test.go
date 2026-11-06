package piscine

import "testing"

func TestCheckNumber(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"no digits", "Hello", false},
		{"one digit", "Hello1", true},
		{"all digits", "123", true},
		{"empty", "", false},
		{"symbols only", "abc!@#", false},
		{"zero", "0", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CheckNumber(tc.input); got != tc.want {
				t.Errorf("CheckNumber(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
