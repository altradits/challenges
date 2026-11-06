package piscine

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	cases := []struct {
		target int
		want   int
	}{
		{1, 0}, {9, 4}, {5, 2}, {0, -1}, {10, -1}, {4, -1},
	}
	for _, tc := range cases {
		if got := BinarySearch(nums, tc.target); got != tc.want {
			t.Errorf("BinarySearch(target %d) = %d, want %d", tc.target, got, tc.want)
		}
	}
	if got := BinarySearch(nil, 1); got != -1 {
		t.Errorf("BinarySearch(nil) = %d, want -1", got)
	}
}

func TestSearchInsertPosition(t *testing.T) {
	nums := []int{1, 3, 5, 7}
	cases := []struct {
		target int
		want   int
	}{
		{1, 0}, {7, 3}, {0, 0}, {8, 4}, {4, 2}, {2, 1},
	}
	for _, tc := range cases {
		if got := SearchInsertPosition(nums, tc.target); got != tc.want {
			t.Errorf("SearchInsertPosition(%d) = %d, want %d", tc.target, got, tc.want)
		}
	}
}

func TestFirstLastOccurrence(t *testing.T) {
	nums := []int{1, 2, 2, 2, 3}
	if got := FindFirstOccurrence(nums, 2); got != 1 {
		t.Errorf("First(2) = %d, want 1", got)
	}
	if got := FindLastOccurrence(nums, 2); got != 3 {
		t.Errorf("Last(2) = %d, want 3", got)
	}
	if got := FindFirstOccurrence(nums, 9); got != -1 {
		t.Errorf("First(9) = %d, want -1", got)
	}
	if got := FindLastOccurrence(nums, 9); got != -1 {
		t.Errorf("Last(9) = %d, want -1", got)
	}
	if got := FindFirstOccurrence(nil, 1); got != -1 {
		t.Errorf("First(nil) = %d, want -1", got)
	}
}
