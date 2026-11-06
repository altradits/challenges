package piscine

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 8}, {10, 89},
	}
	for _, tc := range cases {
		if got := ClimbStairs(tc.n); got != tc.want {
			t.Errorf("ClimbStairs(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 3, 4}, 6, 2},
	}
	for _, tc := range cases {
		if got := CoinChange(tc.coins, tc.amount); got != tc.want {
			t.Errorf("CoinChange(%v, %d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
		}
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
		{"bsbininm", "jmjkbkjkv", 1},
	}
	for _, tc := range cases {
		if got := LongestCommonSubsequence(tc.a, tc.b); got != tc.want {
			t.Errorf("LCS(%q, %q) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestKnapsack01(t *testing.T) {
	weights := []int{1, 2, 3}
	values := []int{6, 10, 12}
	if got := Knapsack01(weights, values, 5); got != 22 {
		t.Errorf("Knapsack01 = %d, want 22", got)
	}
	if got := Knapsack01(weights, values, 0); got != 0 {
		t.Errorf("Knapsack01(cap 0) = %d, want 0", got)
	}
	if got := Knapsack01(weights, values, 1); got != 6 {
		t.Errorf("Knapsack01(cap 1) = %d, want 6", got)
	}
}
