package piscine

import (
	"math"
	"sort"
)

func MedianSorted(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sorted := append([]float64(nil), nums...)
	sort.Float64s(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 1 {
		return sorted[mid]
	}
	return (sorted[mid-1] + sorted[mid]) / 2
}

func TopN(nums []int, n int) []int {
	if n <= 0 {
		return []int{}
	}
	sorted := append([]int(nil), nums...)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	if n > len(sorted) {
		n = len(sorted)
	}
	out := sorted[:n]
	if out == nil {
		return []int{}
	}
	return out
}

func Hypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func RoundToNearest(f float64) float64 {
	return math.Round(f*2) / 2
}
