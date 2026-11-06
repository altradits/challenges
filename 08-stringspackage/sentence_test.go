package piscine

import "testing"

func TestSentenceStats(t *testing.T) {
	cases := []struct {
		name         string
		input        string
		words, chars int
		hasHello     bool
	}{
		{"hello world", "Hello world", 2, 10, true},
		{"shouted hello", "say HELLO to everyone", 4, 18, true},
		{"no greeting", "no greeting here", 3, 14, false},
		{"empty", "", 0, 0, false},
		{"tabs and newlines", "hi\tthere\nyou", 3, 10, false},
		{"hello among words", "well hello there", 3, 14, true},
		{"only spaces", "   ", 0, 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			words, chars, hasHello := SentenceStats(tc.input)
			if words != tc.words || chars != tc.chars || hasHello != tc.hasHello {
				t.Errorf("SentenceStats(%q) = (%d, %d, %v), want (%d, %d, %v)",
					tc.input, words, chars, hasHello, tc.words, tc.chars, tc.hasHello)
			}
		})
	}
}
