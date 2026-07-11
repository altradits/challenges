# Skills for 180-sorting-algorithms

## What You'll Learn

**Previous:** [179-binary-search](../179-binary-search/skills.md) | **Next:** [181-binary-tree](../181-binary-tree/skills.md)

**Challenge:** Implement five sorting algorithms, understand their complexity, and know when to use each.

## Algorithm Comparison

| Algorithm | Best | Average | Worst | Space | Stable? |
|-----------|------|---------|-------|-------|---------|
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Insertion Sort | O(n) | O(n²) | O(n²) | O(1) | Yes |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) | No |
| Counting Sort | O(n+k) | O(n+k) | O(n+k) | O(k) | Yes |

(k = range of values)

## Bubble Sort

Compare adjacent elements; swap if out of order. Repeat until no swaps:

```
[5,3,1,4,2]
[3,5,1,4,2] ← swap 5,3
[3,1,5,4,2] ← swap 5,1
[3,1,4,5,2] ← swap 5,4
[3,1,4,2,5] ← swap 5,2  (5 is now in place)
...repeat...
```

The `swapped` early exit makes it O(n) on already-sorted input.

## Insertion Sort

Build a sorted subarray from left to right, inserting each element in the correct position:

```
[5,3,1,4,2]
 ↑ sorted
[3,5,1,4,2]  ← 3 inserted before 5
[1,3,5,4,2]  ← 1 inserted before 3
[1,3,4,5,2]  ← 4 inserted before 5
[1,2,3,4,5]  ← 2 inserted after 1
```

Excellent for small arrays and nearly-sorted data (what Go's `sort.Slice` uses for small sub-arrays).

## Merge Sort

Divide and conquer: split in half, sort each half, merge:

```
[5,3,1,4,2]
[5,3] [1,4,2]
[5][3] [1][4,2]
       [4][2]
[3,5] [1][2,4]
[3,5] [1,2,4]
[1,2,3,4,5]
```

Guaranteed O(n log n), stable, but uses O(n) extra memory.

## Quick Sort

Partition around a pivot — elements < pivot go left, > pivot go right, then recurse:

```
pivot = 2 (last element of [5,3,1,4,2])
scan:  5>2 skip, 3>2 skip, 1<2 swap to front, 4>2 skip
result: [1, 3, 5, 4, 2]
        [1] | [2] | [3, 5, 4]
```

Average O(n log n), worst case O(n²) (sorted input with bad pivot choice). In practice the fastest for general data.

## Counting Sort

Count occurrences, then reconstruct. Only works for non-negative integers with a known range:

```
[3,1,2,0,3]
counts: [1,1,1,2]  (0→1, 1→1, 2→1, 3→2)
output: [0,1,2,3,3]
```

O(n + max_value) — beats comparison-based O(n log n) when range is small.

## Standard Library

For production code, always use the standard library:

```go
import "sort"
import "slices"  // Go 1.21+

sort.Ints([]int{3,1,2})          // sorts in place
slices.Sort([]int{3,1,2})        // same, generic
sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })  // custom comparator
slices.SortFunc(s, func(a, b int) int { return cmp.Compare(a, b) })
```

These use pattern-defeating quicksort (pdqsort) — hybrid of quicksort, heapsort, and insertion sort for worst-case O(n log n) with excellent constants.

## Documentation

- [sort package](https://pkg.go.dev/sort)
- [slices package (Go 1.21+)](https://pkg.go.dev/slices)
- [Go Blog — Sorting](https://go.dev/blog/sort)

**Next:** [181-binary-tree](../181-binary-tree/README.md)
