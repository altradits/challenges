package piscine

import "testing"

func TestDescribeAny(t *testing.T) {
	cases := []struct {
		name  string
		input any
		want  string
	}{
		{"int", 42, "integer: 42"},
		{"negative int", -7, "integer: -7"},
		{"string", "hello", "string: hello"},
		{"bool true", true, "bool: true"},
		{"bool false", false, "bool: false"},
		{"int slice", []int{1, 2, 3}, "int slice of length 3"},
		{"empty slice", []int{}, "int slice of length 0"},
		{"float", 1.5, "unknown type"},
		{"nil", nil, "unknown type"},
		{"struct", struct{}{}, "unknown type"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Describe(tc.input); got != tc.want {
				t.Errorf("Describe(%#v) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestSumAny(t *testing.T) {
	cases := []struct {
		name  string
		input []any
		want  float64
	}{
		{"mix", []any{1, 2.5, "skip", true, 3}, 6.5},
		{"empty", []any{}, 0},
		{"nil", nil, 0},
		{"ignored only", []any{"x", false}, 0},
		{"negative", []any{-2, -0.5}, -2.5},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sum(tc.input); got != tc.want {
				t.Errorf("Sum(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
