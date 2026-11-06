package piscine

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		name  string
		input []uint
		want  bool
	}{
		{"reachable", []uint{2, 3, 1, 1, 4}, true},
		{"blocked", []uint{3, 2, 1, 0, 4}, false},
		{"single", []uint{0}, true},
		{"zero start", []uint{0, 1}, false},
		{"hop", []uint{1, 0, 2}, false},
		{"long jump", []uint{5, 0, 0, 0, 0, 0}, true},
		{"empty", []uint{}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanJump(tc.input); got != tc.want {
				t.Errorf("CanJump(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
