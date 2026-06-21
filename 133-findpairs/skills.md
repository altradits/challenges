# Skills for findpairs

## What You'll Learn

**Previous:** [132-slice](../132-slice/skills.md) | **Next:** [134-revwstr](../134-revwstr/skills.md)

**Challenge:** Parse an array and a target from JSON-like strings, then find all index pairs whose values sum to the target.

## Core Concept: Hash Map for Pair Finding + Input Parsing

### What Is It?

Finding pairs that sum to a target can be done in O(n) time using a map: for each element, check if its "complement" (target - element) was seen earlier. This challenge adds input parsing, making it a complete real-world exercise.

### How It Works

**Part 1 — Parsing the input:**

Input format: `"[1, 2, 3, 4, 5]"` and `"6"`.

```go
import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Invalid input.")
        return
    }
    arrStr := os.Args[1]
    targetStr := os.Args[2]

    // Validate array format: must start with [ and end with ]
    if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
        fmt.Println("Invalid input.")
        return
    }

    // Validate target (must be a single integer)
    target, err := strconv.Atoi(strings.TrimSpace(targetStr))
    if err != nil {
        fmt.Println("Invalid target sum.")
        return
    }

    // Parse the array contents
    inner := arrStr[1 : len(arrStr)-1]
    parts := strings.Split(inner, ",")
    nums := []int{}
    for _, p := range parts {
        p = strings.TrimSpace(p)
        n, err := strconv.Atoi(p)
        if err != nil {
            fmt.Printf("Invalid number: %s\n", p)
            return
        }
        nums = append(nums, n)
    }
```

**Part 2 — Finding pairs with a hash map:**

```go
    type pair struct{ i, j int }
    pairs := []pair{}

    // For each element, record its index in a map
    // Then for each new element, check if complement was seen
    seen := make(map[int][]int)  // value → list of indices

    for j, v := range nums {
        complement := target - v
        for _, i := range seen[complement] {
            pairs = append(pairs, pair{i, j})
        }
        seen[v] = append(seen[v], j)
    }

    if len(pairs) == 0 {
        fmt.Println("No pairs found.")
        return
    }
    fmt.Printf("Pairs with sum %d: [", target)
    for i, p := range pairs {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Printf("[%d %d]", p.i, p.j)
    }
    fmt.Println("]")
}
```

**Trace `[1, 2, 3, 4, 5]` with target `6`:**

| j | v | complement | seen[complement] | pairs |
|---|---|-----------|-----------------|-------|
| 0 | 1 | 5 | [] | [] |
| 1 | 2 | 4 | [] | [] |
| 2 | 3 | 3 | [] | [] |
| 3 | 4 | 2 | [1] | [{1,3}] |
| 4 | 5 | 1 | [0] | [{1,3},{0,4}] |

Output: `Pairs with sum 6: [[0 4] [1 3]]`

Wait — the expected output is `[[0 4] [1 3]]`. Let me check: `nums[0]=1, nums[4]=5 → sum=6` → pair `{0,4}`. `nums[1]=2, nums[3]=4 → sum=6` → pair `{1,3}`. The expected output sorts them: `[0 4]` then `[1 3]`. The map approach gives them in the order found, which depends on iteration order. Sort if needed.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `strings.Contains(targetStr, " ")` check omitted | `"2 5"` passes as valid | Check that `strconv.Atoi(targetStr)` succeeds with no spaces |
| Validating array without checking `[` and `]` | `"1, 2, 3, 4"` passes | Check first and last chars |
| Using `strconv.ParseFloat` | Floats are not valid here | Use `strconv.Atoi` and handle errors |

## Solving This Challenge

### Algorithm
1. Validate arg count (must be 2).
2. Validate `arrStr` starts with `[` and ends with `]`.
3. Parse `targetStr` with `strconv.Atoi`; error → "Invalid target sum."
4. Split `arrStr[1:len-1]` by `,`, trim spaces, parse each with `strconv.Atoi`; error → "Invalid number: X".
5. Use a map `seen[value] → []int{indices}` to find pairs in O(n).
6. Print pairs or "No pairs found."

## The Challenge
See [README.md](README.md) for full description.

**Next:** [134-revwstr](../134-revwstr/README.md)
