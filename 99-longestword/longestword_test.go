package piscine

import "testing"

func TestLongestWord(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"tie first", "hello world", "hello"},
		{"quick", "the quick brown fox", "quick"},
		{"ccc", "a bb ccc dd eee", "ccc"},
		{"empty", "", ""},
		{"spaces", "   ", ""},
		{"single", "single", "single"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongestWord(tc.input); got != tc.want {
				t.Errorf("LongestWord(%q) = %q, want %q (ties go to the first)", tc.input, got, tc.want)
			}
		})
	}
}
