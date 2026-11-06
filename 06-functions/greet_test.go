package piscine

import "testing"

func TestGreet(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"world", "World", "Hello, World!"},
		{"short", "Go", "Hello, Go!"},
		{"alice", "Alice", "Hello, Alice!"},
		{"empty", "", "Hello, !"},
		{"spaces", "Ada Lovelace", "Hello, Ada Lovelace!"},
		{"unicode", "世界", "Hello, 世界!"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Greet(tc.input); got != tc.want {
				t.Errorf("Greet(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
