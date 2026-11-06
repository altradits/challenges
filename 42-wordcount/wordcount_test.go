package piscine

import "testing"

func TestWordCount(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"two words", "hello world", 2},
		{"empty", "", 0},
		{"single", "go", 1},
		{"extra spaces", "  many   spaces  here ", 3},
		{"tabs", "a\tb\nc", 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WordCount(tc.input); got != tc.want {
				t.Errorf("WordCount(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
