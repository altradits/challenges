# Skills for 183-dynamic-programming

## What You'll Learn

**Previous:** [182-graph-bfs-dfs](../182-graph-bfs-dfs/skills.md) | **Next:** [184-two-pointers](../184-two-pointers/skills.md)

**Challenge:** Solve four classic DP problems — stairs, coin change, LCS, and 0/1 knapsack.

## What Is Dynamic Programming?

Dynamic programming (DP) is **memoised recursion** or **tabulation**. You break a problem into overlapping subproblems, solve each once, and store the result.

Two approaches:
- **Top-down (memoisation):** Recursive solution + cache results
- **Bottom-up (tabulation):** Fill a table iteratively from base cases

Both have the same time complexity. Bottom-up is usually faster in practice (no recursion overhead).

## Pattern: 1D DP Table

### ClimbStairs

```
n stairs: ways to reach stair n = ways to reach n-1 + ways to reach n-2
(Fibonacci with different base cases)

n=1: 1 way
n=2: 2 ways
n=3: dp[2]+dp[1] = 3 ways
n=4: dp[3]+dp[2] = 5 ways
n=5: dp[4]+dp[3] = 8 ways
```

```go
func ClimbStairs(n int) int {
    if n <= 2 { return n }
    a, b := 1, 2
    for i := 3; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```

### CoinChange

```
dp[i] = min coins to make amount i

dp[0] = 0
dp[i] = min over all coins c where c <= i: dp[i-c] + 1
```

```go
func CoinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        dp[i] = amount + 1  // "infinity"
        for _, coin := range coins {
            if coin <= i && dp[i-coin]+1 < dp[i] {
                dp[i] = dp[i-coin] + 1
            }
        }
    }
    if dp[amount] > amount { return -1 }
    return dp[amount]
}
```

## Pattern: 2D DP Table

### Longest Common Subsequence

```
dp[i][j] = LCS length of a[:i] and b[:j]

if a[i-1] == b[j-1]: dp[i][j] = dp[i-1][j-1] + 1
else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
```

Table for a="abcde", b="ace":

```
    ""  a  c  e
""   0  0  0  0
a    0  1  1  1
b    0  1  1  1
c    0  1  2  2
d    0  1  2  2
e    0  1  2  3   ← LCS = 3
```

### 0/1 Knapsack

```
dp[i][w] = max value using first i items with capacity w

if weights[i-1] > w: dp[i][w] = dp[i-1][w]  (can't include item)
else: dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]] + values[i-1])
```

## DP Problem Recognition

A problem is likely DP when:
- It asks for "how many ways", "minimum", "maximum", "longest", "shortest"
- Choices at each step affect future choices
- The problem has **optimal substructure** (optimal solution uses optimal subsolutions)
- The problem has **overlapping subproblems** (naive recursion recomputes the same values)

## Documentation

- [Go Blog — A simple, secure coding problem (DP)](https://go.dev/blog/slices-intro)
- [LeetCode DP patterns](https://leetcode.com/tag/dynamic-programming/)

**Next:** [184-two-pointers](../184-two-pointers/README.md)
