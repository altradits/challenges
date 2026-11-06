package piscine

import (
	"math"
	"reflect"
	"testing"
)

func closeEnough(got, want float64) bool {
	return math.Abs(got-want) < 1e-9
}

func TestMedianSorted(t *testing.T) {
	cases := []struct {
		name  string
		input []float64
		want  float64
	}{
		{"odd", []float64{3, 1, 2}, 2},
		{"even", []float64{4, 1, 3, 2}, 2.5},
		{"single", []float64{7}, 7},
		{"empty", []float64{}, 0},
		{"nil", nil, 0},
		{"unsorted dupes", []float64{5, 5, 1, 9}, 5},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MedianSorted(tc.input); !closeEnough(got, tc.want) {
				t.Errorf("MedianSorted(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestTopN(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		n    int
		want []int
	}{
		{"two of five", []int{3, 1, 4, 1, 5}, 2, []int{5, 4}},
		{"zero", []int{3, 1}, 0, []int{}},
		{"negative", []int{3, 1}, -1, []int{}},
		{"more than len", []int{2, 1}, 5, []int{2, 1}},
		{"empty", []int{}, 3, []int{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := TopN(tc.nums, tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TopN(%v, %d) = %v, want %v", tc.nums, tc.n, got, tc.want)
			}
		})
	}
}

func TestHypotenuse(t *testing.T) {
	if got := Hypotenuse(3, 4); !closeEnough(got, 5) {
		t.Errorf("Hypotenuse(3, 4) = %v, want 5", got)
	}
	if got := Hypotenuse(0, 0); got != 0 {
		t.Errorf("Hypotenuse(0, 0) = %v, want 0", got)
	}
}

func TestRoundToNearest(t *testing.T) {
	cases := []struct {
		input float64
		want  float64
	}{
		{3.2, 3},
		{3.3, 3.5},
		{3.7, 3.5},
		{3.8, 4},
		{-1.2, -1},
		{0, 0},
	}
	for _, tc := range cases {
		if got := RoundToNearest(tc.input); !closeEnough(got, tc.want) {
			t.Errorf("RoundToNearest(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
