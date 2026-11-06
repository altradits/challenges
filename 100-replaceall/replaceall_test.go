package piscine

import "testing"

func TestReplaceAll(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		old, new string
		want     string
	}{
		{"chars", "Hello World", "l", "L", "HeLLo WorLd"},
		{"words", "banana", "na", "NA", "baNANA"},
		{"empty old", "Hello World", "", "x", "Hello World"},
		{"full", "abcabc", "abc", "xyz", "xyzxyz"},
		{"missing", "no match here", "z", "Z", "no match here"},
		{"delete", "aabaa", "a", "", "b"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ReplaceAll(tc.text, tc.old, tc.new); got != tc.want {
				t.Errorf("ReplaceAll(%q, %q, %q) = %q, want %q", tc.text, tc.old, tc.new, got, tc.want)
			}
		})
	}
}
