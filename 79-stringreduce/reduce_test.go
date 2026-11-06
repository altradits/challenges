package piscine

import "testing"

func maxRune(a, b rune) rune {
	if a > b {
		return a
	}
	return b
}

func minRune(a, b rune) rune {
	if a < b {
		return a
	}
	return b
}

func TestReduce(t *testing.T) {
	sum := func(a, b rune) rune { return a + b }
	diff := func(a, b rune) rune { return a - b }

	cases := []struct {
		name string
		s    string
		f    func(rune, rune) rune
		want rune
	}{
		{"max", "abc", maxRune, 'c'},
		{"min", "hello", minRune, 'e'},
		{"sum", "123", sum, rune(49 + 50 + 51)},
		{"order matters", "bcd", diff, rune((('b' - 'c') - 'd'))},
		{"single", "x", sum, 'x'},
		{"empty", "", sum, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reduce(tc.s, tc.f); got != tc.want {
				t.Errorf("Reduce(%q) = %q, want %q (fold left-to-right from first rune)", tc.s, got, tc.want)
			}
		})
	}
}
