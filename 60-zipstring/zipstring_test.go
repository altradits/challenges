package piscine

import "testing"

func TestZipString(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"mixed", "YouuungFellllas", "1Y1o3u1n1g1F1e4l1a1s"},
		{"spaces", "Thee quuick", "1T1h2e1 1q2u1i1c1k"},
		{"simple", "Helloo", "1H1e2l2o"},
		{"empty", "", ""},
		{"single", "a", "1a"},
		{"long run", "aaaaaaaaaa", "10a"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ZipString(tc.input); got != tc.want {
				t.Errorf("ZipString(%q) = %q, want %q (count then char per run)", tc.input, got, tc.want)
			}
		})
	}
}
