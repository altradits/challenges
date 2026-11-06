package piscine

import "testing"

func TestWordCount(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"two", "Hello World", 2},
		{"three", "Go is fun!", 3},
		{"empty", "", 0},
		{"spaces", "   ", 0},
		{"multi space", "One  two   three", 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WordCount(tc.input); got != tc.want {
				t.Errorf("WordCount(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
