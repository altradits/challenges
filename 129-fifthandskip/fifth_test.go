package piscine

import "testing"

func TestFifthAndSkip(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"alpha", "abcdefghijklmnopqrstuwxyz", "abcde ghijk mnopq stuwx z\n"},
		{"sentence", "This is a short sentence", "Thisi ashor sente ce\n"},
		{"short", "1234", "Invalid Input\n"},
		{"empty", "", "\n"},
		{"exact five", "abcde", "abcde\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FifthAndSkip(tc.input); got != tc.want {
				t.Errorf("FifthAndSkip(%q) = %q, want %q (strip spaces, take 5 skip 1)", tc.input, got, tc.want)
			}
		})
	}
}
