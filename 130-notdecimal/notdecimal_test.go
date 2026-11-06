package piscine

import "testing"

func TestNotDecimal(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "0.1", "1\n"},
		{"int part", "174.2", "1742\n"},
		{"long frac", "0.1255", "1255\n"},
		{"longer", "1.20525856", "120525856\n"},
		{"letters kept", "-0.0f00d00", "-0.0f00d00\n"},
		{"empty", "", "\n"},
		{"negative", "-19.525856", "-19525856\n"},
		{"no dot", "1952", "1952\n"},
		{"zero frac", "1.0", "1.0\n"},
		{"not a number", "abc", "abc\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NotDecimal(tc.input); got != tc.want {
				t.Errorf("NotDecimal(%q) = %q, want %q (valid decimals lose the dot)", tc.input, got, tc.want)
			}
		})
	}
}
