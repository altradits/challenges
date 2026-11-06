package piscine

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	cases := []struct {
		name string
		nbrs []int
		want []string
	}{
		{"from 1", []int{1}, []string{"algorithm", "ascii", "package", "golang"}},
		{"2 to 4", []int{2, 4}, []string{"ascii", "package"}},
		{"last 3", []int{-3}, []string{"ascii", "package", "golang"}},
		{"neg range", []int{-2, -1}, []string{"package"}},
		{"end before start", []int{2, 0}, nil},
		{"none", nil, nil},
		{"full", []int{0}, a},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Slice(a, tc.nbrs...)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Slice(a, %v) = %v, want %v", tc.nbrs, got, tc.want)
			}
		})
	}
}
