package piscine

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 4, 2, 10, 2, 3, 1, 0, 20}, 4, 24},
		{[]int{1, 2, 3}, 3, 6},
		{[]int{1, 2}, 5, 0},
		{[]int{-1, -2, -3}, 2, -3},
	}
	for _, tc := range cases {
		if got := MaxSubarraySum(tc.nums, tc.k); got != tc.want {
			t.Errorf("MaxSubarraySum(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
		}
	}
}

func TestLongestUniqueSubstring(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
	}
	for _, tc := range cases {
		if got := LongestUniqueSubstring(tc.input); got != tc.want {
			t.Errorf("LongestUniqueSubstring(%q) = %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestSmallestSubarrayWithSum(t *testing.T) {
	if got := SmallestSubarrayWithSum([]int{2, 3, 1, 2, 4, 3}, 7); got != 2 {
		t.Errorf("got %d, want 2", got)
	}
	if got := SmallestSubarrayWithSum([]int{1, 1, 1}, 10); got != 0 {
		t.Errorf("impossible = %d, want 0", got)
	}
	if got := SmallestSubarrayWithSum([]int{1, 4, 4}, 4); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}

func TestAnagramCount(t *testing.T) {
	cases := []struct {
		s, pattern string
		want       int
	}{
		{"cbaebabacd", "abc", 2},
		{"abab", "ab", 3},
		{"abc", "abcd", 0},
		{"", "a", 0},
	}
	for _, tc := range cases {
		if got := AnagramCount(tc.s, tc.pattern); got != tc.want {
			t.Errorf("AnagramCount(%q, %q) = %d, want %d", tc.s, tc.pattern, got, tc.want)
		}
	}
}
