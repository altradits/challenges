package piscine

import "testing"

func TestIsCapitalized(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"lowercase word", "Hello! How are you?", false},
		{"all caps", "Hello How Are You", true},
		{"digits skipped", "Whats 4this 100K?", true},
		{"trailing digit", "Whatsthis4", true},
		{"leading marks", "!!!!Whatsthis4", true},
		{"empty", "", false},
		{"single lower", "hello", false},
		{"single upper", "Hello", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsCapitalized(tc.input); got != tc.want {
				t.Errorf("IsCapitalized(%q) = %v, want %v (every letter-led word must start uppercase)", tc.input, got, tc.want)
			}
		})
	}
}
