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

---

## Time limit

**30 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 183
```

or run the tests directly:

```bash
go test ./183-dynamic-programming/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 183` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
