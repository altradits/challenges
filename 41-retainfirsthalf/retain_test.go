package piscine

import "testing"

func TestRetainFirstHalf(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"even", "This is the 1st halfThis is the 2nd half", "This is the 1st half"},
		{"single", "A", "A"},
		{"empty", "", ""},
		{"odd drops middle", "Hello World", "Hello"},
		{"two chars", "ab", "a"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RetainFirstHalf(tc.input); got != tc.want {
				t.Errorf("RetainFirstHalf(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
