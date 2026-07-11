# 180-sorting-algorithms — Implement Classic Sorting Algorithms

## Challenge

Implement in `package piscine` (do NOT use `sort.Slice` — implement from scratch):

```go
// BubbleSort sorts nums in ascending order in-place. O(n²)
func BubbleSort(nums []int)

// InsertionSort sorts nums in ascending order in-place. O(n²) worst, O(n) best.
func InsertionSort(nums []int)

// MergeSort returns a new sorted slice using divide-and-conquer. O(n log n)
func MergeSort(nums []int) []int

// QuickSort sorts nums in-place using partition. O(n log n) average.
func QuickSort(nums []int)

// CountingSort sorts nums (non-negative integers) in-place. O(n + max)
func CountingSort(nums []int)
```

## Examples

```
BubbleSort(&[5,3,1,4,2])     →  [1,2,3,4,5]
InsertionSort(&[5,3,1,4,2])  →  [1,2,3,4,5]
MergeSort([5,3,1,4,2])       →  [1,2,3,4,5]  (new slice)
QuickSort(&[5,3,1,4,2])      →  [1,2,3,4,5]
CountingSort(&[3,1,2,0,3])   →  [0,1,2,3,3]
```

Read `skills.md` before you start.
