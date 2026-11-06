package piscine

import (
	"reflect"
	"testing"
)

func TestTwoSum184(t *testing.T) {
	lo, hi := TwoSum([]int{1, 2, 3, 4, 6}, 6)
	if lo != 1 || hi != 3 {
		t.Errorf("TwoSum = (%d, %d), want (1, 3)", lo, hi)
	}
	lo, hi = TwoSum([]int{2, 7, 11, 15}, 9)
	if lo != 0 || hi != 1 {
		t.Errorf("TwoSum = (%d, %d), want (0, 1)", lo, hi)
	}
	lo, hi = TwoSum([]int{1, 2, 3}, 100)
	if lo != -1 || hi != -1 {
		t.Errorf("TwoSum missing = (%d, %d), want (-1, -1)", lo, hi)
	}
}

func TestRemoveDuplicates184(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3}
	n := RemoveDuplicates(nums)
	if n != 3 {
		t.Fatalf("RemoveDuplicates length = %d, want 3", n)
	}
	if !reflect.DeepEqual(nums[:n], []int{1, 2, 3}) {
		t.Errorf("nums[:%d] = %v, want [1 2 3]", n, nums[:n])
	}
	if n := RemoveDuplicates(nil); n != 0 {
		t.Errorf("RemoveDuplicates(nil) = %d, want 0", n)
	}
	if n := RemoveDuplicates([]int{5}); n != 1 {
		t.Errorf("RemoveDuplicates([5]) = %d, want 1", n)
	}
}

func TestReverseInPlace(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ReverseInPlace(nums)
	if !reflect.DeepEqual(nums, []int{5, 4, 3, 2, 1}) {
		t.Errorf("reversed = %v", nums)
	}
	even := []int{1, 2, 3, 4}
	ReverseInPlace(even)
	if !reflect.DeepEqual(even, []int{4, 3, 2, 1}) {
		t.Errorf("reversed = %v", even)
	}
}

func TestContainerWithMostWater(t *testing.T) {
	if got := ContainerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); got != 49 {
		t.Errorf("got %d, want 49", got)
	}
	if got := ContainerWithMostWater([]int{1, 1}); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	if got := ContainerWithMostWater([]int{4, 3, 2, 1, 4}); got != 16 {
		t.Errorf("got %d, want 16", got)
	}
}
