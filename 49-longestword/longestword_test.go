package piscine

import "testing"

func TestLongestWord(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"tie first wins", "Hello World", "Hello"},
		{"fun", "Go is fun", "fun"},
		{"empty", "", ""},
		{"growing", "a bb ccc dddd", "dddd"},
		{"spaces", "   ", ""},
		{"single", "word", "word"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongestWord(tc.input); got != tc.want {
				t.Errorf("LongestWord(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
