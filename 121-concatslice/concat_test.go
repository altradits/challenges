package piscine

import (
	"reflect"
	"testing"
)

func TestConcatSlice(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 []int
		want   []int
	}{
		{"basic", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"empty first", []int{}, []int{4, 5, 6, 7, 8, 9}, []int{4, 5, 6, 7, 8, 9}},
		{"empty second", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"both empty", []int{}, []int{}, []int{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ConcatSlice(tc.s1, tc.s2)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ConcatSlice(%v, %v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
