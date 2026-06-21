# Skills for 58-stringcount

## What You'll Learn

**Previous:** [57-stringindex](../57-stringindex/skills.md) | **Next:** [59-stringprefix](../59-stringprefix/skills.md)

**Challenge:** Implement `Count(s, substr string) int` — count all occurrences of `substr` in `s`, including overlapping ones.

## Core Concept: `strings.Count` and the Sliding Counter Loop

### The Built-in Function

```go
import "strings"

strings.Count("hello hello", "l")   // 4
strings.Count("hello hello", "ll")  // 2
strings.Count("aaa", "aa")          // 1  (non-overlapping!)
strings.Count("hello", "")          // 6  (one between each char + ends)
```

**Note**: `strings.Count` counts NON-overlapping occurrences. If your challenge says "count overlapping", you need the manual approach below.

### Non-overlapping Count (what `strings.Count` does)

```go
func Count(s, substr string) int {
    count := 0
    start := 0
    for {
        idx := strings.Index(s[start:], substr)
        if idx == -1 {
            break
        }
        count++
        start += idx + len(substr)  // skip past the match
    }
    return count
}
```

### Overlapping Count (advance by 1 each time)

```go
func CountOverlapping(s, substr string) int {
    count := 0
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            count++
            // do NOT skip — advance by 1 to catch overlaps
        }
    }
    return count
}
```

Example: `CountOverlapping("aaa", "aa")` → 2 (positions 0 and 1 both match)

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not handling empty substr | Infinite loop | Return 0 or special value |
| Advancing by len(substr) for overlapping count | Misses overlaps | Advance by 1 |
| Using `strings.Count` when overlap needed | Wrong answer | Use manual loop |

## Algorithm (non-overlapping)

1. If `substr` is empty, return 0
2. `start = 0`, `count = 0`
3. Loop: find `substr` in `s[start:]`; if -1 break; count++; start += idx+len(substr)
4. Return `count`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [59-stringprefix](../59-stringprefix/README.md)
