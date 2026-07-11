# 185-sliding-window — Sliding Window Technique

## Challenge

Implement in `package piscine`:

```go
// MaxSubarraySum returns the maximum sum of any contiguous subarray of length k.
func MaxSubarraySum(nums []int, k int) int

// LongestUniqueSubstring returns the length of the longest substring
// with no repeating characters.
func LongestUniqueSubstring(s string) int

// SmallestSubarrayWithSum returns the minimum length of a contiguous
// subarray whose sum >= target. Returns 0 if no such subarray exists.
func SmallestSubarrayWithSum(nums []int, target int) int

// AnagramCount returns the number of anagrams of pattern in s.
func AnagramCount(s, pattern string) int
```

## Examples

```
MaxSubarraySum([]int{2,3,4,1,5}, 3)  →  9  (3+4+1 or 2+3+4?)
// max window of k=3: [2,3,4]=9, [3,4,1]=8, [4,1,5]=10 → 10

LongestUniqueSubstring("abcabcbb")   →  3  ("abc")
LongestUniqueSubstring("bbbbb")      →  1  ("b")
LongestUniqueSubstring("pwwkew")     →  3  ("wke")

SmallestSubarrayWithSum([]int{2,3,1,2,4,3}, 7)  →  2  ([4,3])
SmallestSubarrayWithSum([]int{1,1,1,1}, 100)     →  0

AnagramCount("cbaebabacd", "abc")  →  2  ("cba" at 0, "abc" at 6)
```

Read `skills.md` before you start.
