# Skills for 179-binary-search

## What You'll Learn

**Previous:** [178-queue](../178-queue/skills.md) | **Next:** [180-sorting-algorithms](../180-sorting-algorithms/skills.md)

**Challenge:** Master binary search and its variations — the most important search algorithm.

## Binary Search

Binary search finds a target in a **sorted** slice by repeatedly halving the search space.

```
Array: [1, 3, 5, 7, 9]  target: 7
       lo=0, hi=4, mid=2 → nums[2]=5 < 7 → lo = 3
       lo=3, hi=4, mid=3 → nums[3]=7 = 7 → found at index 3
```

Time: O(log n) — a sorted array of 1 billion elements takes at most 30 comparisons.

```go
func BinarySearch(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2  // avoids overflow vs (lo+hi)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return -1
}
```

**Why `lo + (hi-lo)/2` instead of `(lo+hi)/2`?**
If `lo` and `hi` are both large ints, `lo+hi` can overflow. The subtraction form avoids this.

## Variations

### Find Insert Position

When the target is not found, `lo` holds the correct insert position:

```go
func SearchInsertPosition(nums []int, target int) int {
    lo, hi := 0, len(nums)
    for lo < hi {
        mid := lo + (hi-lo)/2
        if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid
        }
    }
    return lo
}
```

### First Occurrence (Left Boundary)

When duplicates exist, continue searching left after finding the target:

```go
func FindFirstOccurrence(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    result := -1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if nums[mid] == target {
            result = mid
            hi = mid - 1  // keep searching left
        } else if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return result
}
```

### Last Occurrence (Right Boundary)

```go
func FindLastOccurrence(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    result := -1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if nums[mid] == target {
            result = mid
            lo = mid + 1  // keep searching right
        } else if nums[mid] < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return result
}
```

## Standard Library

Go's `sort` package provides binary search utilities:

```go
import "sort"

idx := sort.SearchInts([]int{1,3,5,7,9}, 7)  // returns 3
idx := sort.Search(n, func(i int) bool { return nums[i] >= target })
```

`sort.Search` is the most general — it finds the first index where `f(i)` is true, given a monotone function.

## Documentation

- [sort package](https://pkg.go.dev/sort)
- [sort.Search](https://pkg.go.dev/sort#Search)
- [Go Blog — Binary Search](https://research.swtch.com/go-lessons)

**Next:** [180-sorting-algorithms](../180-sorting-algorithms/README.md)
