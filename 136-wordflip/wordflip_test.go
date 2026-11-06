package piscine

import "testing"

func TestWordFlip(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"basic", "First second last", "last second First\n"},
		{"empty", "", "Invalid Output\n"},
		{"spaces", "     ", "\n"},
		{"messy", " hello  all  of  you! ", "you! of all hello\n"},
		{"single", "solo", "solo\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WordFlip(tc.input); got != tc.want {
				t.Errorf("WordFlip(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
