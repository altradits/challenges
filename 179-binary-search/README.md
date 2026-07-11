# 179-binary-search — Binary Search

## Challenge

Implement in `package piscine`:

```go
// BinarySearch returns the index of target in sorted slice nums,
// or -1 if not found. Slice must be sorted in ascending order.
func BinarySearch(nums []int, target int) int

// SearchInsertPosition returns the index where target should be inserted
// to keep nums sorted. If target exists, returns its index.
func SearchInsertPosition(nums []int, target int) int

// FindFirstOccurrence returns the index of the FIRST occurrence of target,
// or -1 if not found.
func FindFirstOccurrence(nums []int, target int) int

// FindLastOccurrence returns the index of the LAST occurrence of target,
// or -1 if not found.
func FindLastOccurrence(nums []int, target int) int
```

## Examples

```
BinarySearch([]int{1,3,5,7,9}, 5)   →  2
BinarySearch([]int{1,3,5,7,9}, 6)   →  -1
BinarySearch([]int{}, 5)             →  -1

SearchInsertPosition([]int{1,3,5,6}, 5)  →  2
SearchInsertPosition([]int{1,3,5,6}, 2)  →  1
SearchInsertPosition([]int{1,3,5,6}, 7)  →  4

FindFirstOccurrence([]int{1,2,2,2,3}, 2)  →  1
FindLastOccurrence([]int{1,2,2,2,3}, 2)   →  3
```

Read `skills.md` before you start.
