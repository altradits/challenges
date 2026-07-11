package piscine

import "strings"

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("Max requires at least one argument")
	}
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func Join(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

func Wrap(fn func(int) int, nums ...int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}
