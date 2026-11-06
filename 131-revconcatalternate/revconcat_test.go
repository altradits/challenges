package piscine

import (
	"reflect"
	"testing"
)

func TestRevConcatAlternate(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 []int
		want   []int
	}{
		{"equal", []int{1, 2, 3}, []int{4, 5, 6}, []int{3, 6, 2, 5, 1, 4}},
		{"second longer", []int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 3, 6, 2, 5, 1, 4}},
		{"first longer", []int{1, 2, 3, 9, 8}, []int{4, 5}, []int{8, 9, 3, 2, 5, 1, 4}},
		{"empty second", []int{1, 2, 3}, []int{}, []int{3, 2, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RevConcatAlternate(tc.s1, tc.s2)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RevConcatAlternate(%v, %v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
