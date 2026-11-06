package piscine

import "testing"

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"word", "hello", false},
		{"space", " ", false},
		{"single", "x", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsEmpty(tc.input); got != tc.want {
				t.Errorf("IsEmpty(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
