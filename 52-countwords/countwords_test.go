package piscine

import "testing"

func TestCountWordsAdvanced(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"punctuation", "Hello, World!", 2},
		{"spaces", "Go is fun", 3},
		{"digits in word", "hello2world", 1},
		{"empty", "", 0},
		{"only punctuation", "!!!", 0},
		{"underscores split", "a_b", 2},
		{"tabs", "one\ttwo", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountWordsAdvanced(tc.input); got != tc.want {
				t.Errorf("CountWordsAdvanced(%q) = %d, want %d (words = letter/digit runs)", tc.input, got, tc.want)
			}
		})
	}
}
