# 183-dynamic-programming — Dynamic Programming

## Challenge

Implement in `package piscine`:

```go
// ClimbStairs: given n stairs, count ways to climb if you can take 1 or 2 steps.
func ClimbStairs(n int) int

// CoinChange: given coin denominations and a target amount,
// return minimum coins needed. Return -1 if impossible.
func CoinChange(coins []int, amount int) int

// LongestCommonSubsequence returns the length of the LCS of strings a and b.
func LongestCommonSubsequence(a, b string) int

// Knapsack01: given weights and values (both len n) and capacity W,
// return the maximum value achievable.
func Knapsack01(weights, values []int, capacity int) int
```

## Examples

```
ClimbStairs(5)                              →  8
ClimbStairs(1)                              →  1
ClimbStairs(2)                              →  2

CoinChange([]int{1,5,10,25}, 36)           →  3  (25+10+1)
CoinChange([]int{2}, 3)                    →  -1

LongestCommonSubsequence("abcde", "ace")   →  3  ("ace")
LongestCommonSubsequence("abc", "abc")     →  3
LongestCommonSubsequence("abc", "def")     →  0

Knapsack01([]int{2,3,4,5}, []int{3,4,5,6}, 8)  →  10
```

Read `skills.md` before you start.
