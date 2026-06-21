# Skills for fromto

## What You'll Learn

**Previous:** [../38-findprevprime/skills.md](../38-findprevprime/skills.md) | **Next:** [../40-iscapitalized/README.md](../40-iscapitalized/README.md)

**Challenge:** Return a string showing all numbers from `from` to `to`, comma-separated, zero-padded to two digits, with validation for the 0–99 range.

## Core Concept: Building a Range String With fmt.Sprintf and Padding

### What Is It?
Loop over the range (forward or backward depending on which is larger), format each number with leading zero if needed, join with `", "`, and validate that both arguments are in 0–99.

```go
import (
    "fmt"
    "strings"
)

func FromTo(from int, to int) string {
    if from > 99 || from < 0 || to > 99 || to < 0 {
        return "Invalid\n"
    }
    parts := []string{}
    if from <= to {
        for i := from; i <= to; i++ {
            parts = append(parts, fmt.Sprintf("%02d", i))
        }
    } else {
        for i := from; i >= to; i-- {
            parts = append(parts, fmt.Sprintf("%02d", i))
        }
    }
    return strings.Join(parts, ", ") + "\n"
}
```

### Zero-Padding With fmt.Sprintf("%02d", n)
`%02d` formats an integer with at least 2 digits, padding with a leading zero if needed:
```go
fmt.Sprintf("%02d", 5)  // "05"
fmt.Sprintf("%02d", 10) // "10"
fmt.Sprintf("%02d", 0)  // "00"
```

The `0` after `%` means "pad with zeros". The `2` is the minimum width.

### Building the List
Collect formatted numbers into a slice, then join:
```go
parts := []string{}
for i := from; i <= to; i++ {
    parts = append(parts, fmt.Sprintf("%02d", i))
}
return strings.Join(parts, ", ") + "\n"
```

### Handling Reverse Range (from > to)
When `from > to`, count downward:
```go
for i := from; i >= to; i-- {
    parts = append(parts, fmt.Sprintf("%02d", i))
}
```

### When from == to
Both loops produce exactly one element: the number itself. `strings.Join([]string{"10"}, ", ")` = `"10"`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `%2d` without `0` | Pads with spaces, not zeros | Use `%02d` |
| Only checking upper bound | `from = -1` passes through | Check both `< 0` and `> 99` for both args |
| Forgetting `\n` at the end | Output missing newline | Add `+ "\n"` after the join |

## Solving This Challenge

### Algorithm
1. If either argument is `< 0` or `> 99`, return `"Invalid\n"`.
2. Build `parts` slice: if `from <= to`, count up; else count down.
3. Return `strings.Join(parts, ", ") + "\n"`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../40-iscapitalized/README.md](../40-iscapitalized/README.md)
