package piscine

import "testing"

func TestLastChar(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"hello", "Hello", "o"},
		{"single", "a", "a"},
		{"empty", "", ""},
		{"unicode", "世界", "界"},
		{"trailing space", "hi ", " "},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LastChar(tc.input); got != tc.want {
				t.Errorf("LastChar(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
