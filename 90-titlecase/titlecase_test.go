package piscine

import "testing"

func TestTitleCase(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"basic", "hello world", "Hello World"},
		{"punctuation", "Go is fun!", "Go Is Fun!"},
		{"empty", "", ""},
		{"spacing kept", "multiple   spaces   here", "Multiple   Spaces   Here"},
		{"all caps", "ALL CAPS", "All Caps"},
		{"single", "a", "A"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TitleCase(tc.input); got != tc.want {
				t.Errorf("TitleCase(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
