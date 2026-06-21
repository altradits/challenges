# Skills for 92-countchar

## What You'll Learn

**Previous:** [91-findchar](../91-findchar/skills.md) | **Next:** [93-findlastchar](../93-findlastchar/skills.md)

**Challenge:** Write a function `CountChar(s string, c rune) int` that returns the number of times character `c` appears in string `s`.

## Core Concept: The Full-Scan Accumulator for a Specific Target

### FindChar vs CountChar

| Function | Strategy | Returns |
|----------|----------|---------|
| `FindChar` | Stop at first match | `int` (index or -1) |
| `CountChar` | Scan ALL, count every match | `int` (count) |

`CountChar` must **not** return early — you need to find every occurrence, not just the first.

### The Implementation

```go
func CountChar(s string, c rune) int {
    count := 0
    for _, ch := range s {
        if ch == c {
            count++
        }
    }
    return count
}
```

This is the accumulator pattern from [76-stringlength](../76-stringlength/skills.md), filtered by a specific character comparison. Notice that the index `i` is not needed — use `_` to discard it.

### Two-Parameter Functions with a `rune` Argument

You saw this signature in [91-findchar](../91-findchar/skills.md):

```go
func CountChar(s string, c rune) int
```

Called with:

```go
CountChar("Hello World", 'l')   // returns 3
CountChar("   ", ' ')           // returns 3
```

### Case Sensitivity

The comparison `ch == c` is case-sensitive. `'a'` and `'A'` are different runes:

```go
CountChar("Hello", 'h')   // 0 — no lowercase 'h'
CountChar("Hello", 'H')   // 1
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `return count` inside the loop | Returns after counting only the first match | Return only after the loop |
| `return 1` on first match | Only counts existence, not all occurrences | Use `count++` and return after loop |
| Forgetting `count := 0` | `count` is undeclared | Declare before the loop |

## Solving This Challenge

### Algorithm

1. `count := 0`
2. For each rune `ch` in `s`:
   - If `ch == c`: `count++`
3. Return `count`

### Trace Through an Example

Input: `CountChar("Hello World", 'l')`

| ch | ch=='l'? | count |
|----|---------|-------|
| H | No | 0 |
| e | No | 0 |
| l | Yes | 1 |
| l | Yes | 2 |
| o | No | 2 |
| ` ` | No | 2 |
| W | No | 2 |
| o | No | 2 |
| r | No | 2 |
| l | Yes | 3 |
| d | No | 3 |

Result: `3`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [93-findlastchar](../93-findlastchar/skills.md)
