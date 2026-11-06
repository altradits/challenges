package piscine

import "testing"

func TestItoa(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  string
	}{
		{"positive", 12345, "12345"},
		{"zero", 0, "0"},
		{"negative", -1234, "-1234"},
		{"large", 987654321, "987654321"},
		{"single", 7, "7"},
		{"minus one", -1, "-1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Itoa(tc.input); got != tc.want {
				t.Errorf("Itoa(%d) = %q, want %q (no strconv)", tc.input, got, tc.want)
			}
		})
	}
}
