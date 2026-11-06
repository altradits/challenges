package piscine

import "testing"

func TestRemoveSpaces(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"sentence", "hello world", "helloworld"},
		{"empty", "", ""},
		{"none", "abc", "abc"},
		{"only spaces", "   ", ""},
		{"edges", "  hi  ", "hi"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RemoveSpaces(tc.input); got != tc.want {
				t.Errorf("RemoveSpaces(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
