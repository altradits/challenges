package piscine

import (
	"reflect"
	"testing"
)

var sortCases = [][]int{
	{},
	{1},
	{3, 1, 2},
	{5, 4, 3, 2, 1},
	{1, 2, 3, 4, 5},
	{3, 1, 2, 1, 3},
	{-5, 0, -1, 3, 2},
}

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func TestInPlaceSorts(t *testing.T) {
	sorts := map[string]func([]int){
		"BubbleSort":    BubbleSort,
		"InsertionSort": InsertionSort,
		"QuickSort":     QuickSort,
	}
	for name, sortFn := range sorts {
		t.Run(name, func(t *testing.T) {
			for _, tc := range sortCases {
				in := append([]int(nil), tc...)
				sortFn(in)
				if !isSorted(in) {
					t.Errorf("%s(%v) = %v, not sorted", name, tc, in)
				}
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range sortCases {
		in := append([]int(nil), tc...)
		orig := append([]int(nil), tc...)
		got := MergeSort(in)
		if !isSorted(got) {
			t.Errorf("MergeSort(%v) = %v", tc, got)
		}
		if !reflect.DeepEqual(in, orig) {
			t.Errorf("MergeSort must not modify input: %v -> %v", orig, in)
		}
	}
}

func TestCountingSort(t *testing.T) {
	cases := [][]int{{}, {1}, {3, 1, 2}, {5, 5, 5}, {9, 0, 3, 7, 3}}
	for _, tc := range cases {
		in := append([]int(nil), tc...)
		CountingSort(in)
		if !isSorted(in) {
			t.Errorf("CountingSort(%v) = %v", tc, in)
		}
	}
}
