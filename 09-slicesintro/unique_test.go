package piscine

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"mixed dupes", []int{1, 2, 2, 3, 1, 4}, []int{1, 2, 3, 4}},
		{"empty", []int{}, []int{}},
		{"nil", nil, []int{}},
		{"single", []int{5}, []int{5}},
		{"all same", []int{7, 7, 7}, []int{7}},
		{"order kept", []int{3, 1, 2, 1, 3}, []int{3, 1, 2}},
		{"negatives", []int{-1, 0, -1, 2}, []int{-1, 0, 2}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Unique(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Unique(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
