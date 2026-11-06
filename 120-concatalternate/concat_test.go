package piscine

import (
	"reflect"
	"testing"
)

func TestConcatAlternate(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 []int
		want   []int
	}{
		{"equal", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 4, 2, 5, 3, 6}},
		{"second longer", []int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{"tail kept", []int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}, []int{4, 1, 5, 2, 6, 3, 7, 8, 9}},
		{"empty second", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"both empty", []int{}, []int{}, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConcatAlternate(tc.s1, tc.s2)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ConcatAlternate(%v, %v) = %v, want %v (alternate from the longer slice)", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
