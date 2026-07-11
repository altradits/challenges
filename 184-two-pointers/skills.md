# Skills for 184-two-pointers

## What You'll Learn

**Previous:** [183-dynamic-programming](../183-dynamic-programming/skills.md) | **Next:** [185-sliding-window](../185-sliding-window/skills.md)

**Challenge:** Eliminate nested loops using two index pointers that move toward each other or in the same direction.

## The Two-Pointer Pattern

Two pointers are most useful when:
- The input is **sorted** (or can be sorted cheaply)
- You need to find **pairs** or **subarrays** with a specific property
- A naive O(n²) solution exists and you need O(n)

### Converging Pointers (move toward each other)

```go
lo, hi := 0, len(nums)-1
for lo < hi {
    // decide which pointer to move
    if condition_to_move_lo {
        lo++
    } else {
        hi--
    }
}
```

### Slow/Fast Pointers (move in same direction)

```go
slow := 0
for fast := 1; fast < len(nums); fast++ {
    if should_keep(nums[fast]) {
        slow++
        nums[slow] = nums[fast]
    }
}
// slow+1 is the new length
```

## TwoSum on Sorted Array

Since the array is sorted:
- If `sum < target`: increase `lo` to get a larger left value
- If `sum > target`: decrease `hi` to get a smaller right value
- If `sum == target`: found it

O(n) vs O(n²) for nested loops.

## Remove Duplicates In-Place

Slow pointer marks where the next unique element goes. Fast pointer scans ahead:

```
[1, 1, 2, 2, 3]
 s  f               nums[f] == nums[s] — skip
    s  f            nums[f] != nums[s] — copy: s++, nums[s]=nums[f]
       s     f      nums[f] == nums[s] — skip
          s  f      nums[f] != nums[s] — copy
             s
Result: [1, 2, 3, 2, 3], length = 3 (first 3 elements)
```

## Container With Most Water

Start with widest container (lo=0, hi=n-1). Move the shorter height inward — moving the taller one can only decrease width without increasing height.

```
heights: [1, 8, 6, 2, 5, 4, 8, 3, 7]
lo=0(h=1), hi=8(h=7): water=min(1,7)*8=8, move lo (shorter)
lo=1(h=8), hi=8(h=7): water=min(8,7)*7=49, move hi
lo=1(h=8), hi=7(h=3): water=min(8,3)*6=18, move hi
...max so far is 49
```

**Next:** [185-sliding-window](../185-sliding-window/README.md)
