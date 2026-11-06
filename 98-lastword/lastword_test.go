package piscine

import "testing"

func TestLastWord(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"sparta", "this        ...       is sparta, then again, maybe    not", "not\n"},
		{"comma kept", " lorem,ipsum ", "lorem,ipsum\n"},
		{"only spaces", " ", "\n"},
		{"empty", "", "\n"},
		{"single", "word", "word\n"},
		{"trailing spaces", "hello world   ", "world\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LastWord(tc.input); got != tc.want {
				t.Errorf("LastWord(%q) = %q, want %q (last space-delimited word + newline)", tc.input, got, tc.want)
			}
		})
	}
}
