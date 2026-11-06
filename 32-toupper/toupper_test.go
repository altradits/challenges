package piscine

import "testing"

func TestToUpper(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"lower", "hello", "HELLO"},
		{"mixed", "Go Lang", "GO LANG"},
		{"empty", "", ""},
		{"already", "ABC", "ABC"},
		{"digits", "a1b2", "A1B2"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToUpper(tc.input); got != tc.want {
				t.Errorf("ToUpper(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
