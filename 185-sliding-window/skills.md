# Skills for 185-sliding-window

## What You'll Learn

**Previous:** [184-two-pointers](../184-two-pointers/skills.md) | **Next:** [frameworks/01-gin](../frameworks/01-gin/README.md)

**Challenge:** Use a moving window over arrays and strings to avoid recomputing from scratch on each step.

## The Sliding Window Pattern

A sliding window maintains a range `[left, right]` and slides it over the input, adding from the right and removing from the left. This converts many O(n²) problems to O(n).

### Fixed-Size Window

```go
// Sum of every window of size k
windowSum := 0
for i := 0; i < k; i++ { windowSum += nums[i] }

for i := k; i < len(nums); i++ {
    windowSum += nums[i]       // add new right element
    windowSum -= nums[i-k]     // remove old left element
    // windowSum now covers [i-k+1, i]
}
```

### Variable-Size Window

```go
left := 0
for right := 0; right < len(nums); right++ {
    // expand window: include nums[right]
    
    for window_is_invalid {
        // shrink window: exclude nums[left]
        left++
    }
    
    // record result for valid window [left, right]
}
```

## LongestUniqueSubstring

Track the last seen index of each character. If we see a repeat inside the window, move `left` past the previous occurrence:

```go
charIndex := map[rune]int{}
left := 0
maxLen := 0

for right, ch := range s {
    if idx, ok := charIndex[ch]; ok && idx >= left {
        left = idx + 1  // shrink past the duplicate
    }
    charIndex[ch] = right
    if right-left+1 > maxLen {
        maxLen = right - left + 1
    }
}
```

## SmallestSubarrayWithSum

Expand right until sum >= target, then shrink from left as long as the condition holds:

```go
for right := 0; right < len(nums); right++ {
    windowSum += nums[right]
    for windowSum >= target {
        minLen = min(minLen, right-left+1)
        windowSum -= nums[left]
        left++
    }
}
```

## When to Use Sliding Window

| Pattern | When |
|---------|------|
| Fixed window | Subarray/substring of exactly k elements |
| Variable window (expand-contract) | Longest/smallest subarray meeting a condition |
| Two frequency maps | Anagram / permutation counting in string |

## Documentation

- [Go Spec — For statements with range clause](https://go.dev/ref/spec#For_range)

**Next:** [frameworks/01-gin](../frameworks/01-gin/README.md)
