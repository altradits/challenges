package piscine

import "testing"

func TestToLower(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"upper", "HELLO", "hello"},
		{"mixed", "Go Lang", "go lang"},
		{"empty", "", ""},
		{"already", "abc", "abc"},
		{"digits", "A1B2", "a1b2"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToLower(tc.input); got != tc.want {
				t.Errorf("ToLower(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
