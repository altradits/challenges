package piscine

import "testing"

func TestFromTo(t *testing.T) {
	cases := []struct {
		name     string
		from, to int
		want     string
	}{
		{"up", 1, 10, "01, 02, 03, 04, 05, 06, 07, 08, 09, 10\n"},
		{"down", 10, 1, "10, 09, 08, 07, 06, 05, 04, 03, 02, 01\n"},
		{"same", 10, 10, "10\n"},
		{"out of range", 100, 10, "Invalid\n"},
		{"negative", -1, 5, "Invalid\n"},
		{"zero", 0, 0, "00\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FromTo(tc.from, tc.to); got != tc.want {
				t.Errorf("FromTo(%d, %d) = %q, want %q", tc.from, tc.to, got, tc.want)
			}
		})
	}
}
