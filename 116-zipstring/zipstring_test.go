package piscine

import "testing"

func TestZipString(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"mixed", "YouuungFellllas", "1Y1o3u1n1g1F1e4l1a1s"},
		{"sentence", "Thee quuick browwn fox juumps over the laaazy dog", "1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g"},
		{"hello", "Helloo Therre!", "1H1e2l2o1 1T1h1e2r1e1!"},
		{"empty", "", ""},
		{"single", "a", "1a"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ZipString(tc.input); got != tc.want {
				t.Errorf("ZipString(%q) = %q, want %q (count then char per run)", tc.input, got, tc.want)
			}
		})
	}
}
