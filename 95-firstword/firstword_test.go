package piscine

import "testing"

func TestFirstWord(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"basic", "hello there", "hello\n"},
		{"empty", "", "\n"},
		{"dots", "hello   .........  bye", "hello\n"},
		{"leading spaces", "   hello", "hello\n"},
		{"only spaces", "   ", "\n"},
		{"single", "word", "word\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstWord(tc.input); got != tc.want {
				t.Errorf("FirstWord(%q) = %q, want %q (first space-delimited word + newline)", tc.input, got, tc.want)
			}
		})
	}
}
