package piscine

import "testing"

func TestItoaBase(t *testing.T) {
	cases := []struct {
		name  string
		value int
		base  int
		want  string
	}{
		{"hex A", 10, 16, "A"},
		{"hex FF", 255, 16, "FF"},
		{"hex negative", -255, 16, "-FF"},
		{"binary", 5, 2, "101"},
		{"zero", 0, 10, "0"},
		{"octal", 8, 8, "10"},
		{"decimal", 42, 10, "42"},
		{"base 4", 7, 4, "13"},
		{"negative binary", -5, 2, "-101"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ItoaBase(tc.value, tc.base); got != tc.want {
				t.Errorf("ItoaBase(%d, %d) = %q, want %q", tc.value, tc.base, got, tc.want)
			}
		})
	}
}
