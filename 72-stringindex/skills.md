# Skills for 72-stringindex

## What You'll Learn

**Previous:** [71-stringcontains](../71-stringcontains/skills.md) | **Next:** [73-stringcount](../73-stringcount/skills.md)

**Challenge:** Implement `Index(s, substr string) int` — find the byte position of the first occurrence of `substr` in `s`, or return -1.

## Core Concept: `strings.Index` and the Sliding Window Search

### The Built-in Function

```go
import "strings"

strings.Index("hello world", "world")  // 6
strings.Index("hello world", "xyz")    // -1
strings.Index("hello world", "")       // 0  (empty always found at 0)
strings.Index("", "a")                 // -1
```

**Key variants:**
```go
strings.Index(s, substr)       // first occurrence
strings.LastIndex(s, substr)   // last occurrence
strings.IndexRune(s, 'x')      // first occurrence of a rune
strings.IndexAny(s, "aeiou")   // first occurrence of any char in the set
strings.IndexByte(s, 'x')      // first occurrence of a byte
```

### How Manual Search Works (Sliding Window)

```go
func Index(s, substr string) int {
    if len(substr) == 0 {
        return 0
    }
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return i
        }
    }
    return -1
}
```

The loop bound `i <= len(s)-len(substr)` prevents checking positions where the substring can't fit. This is the **sliding window** pattern.

### Why the Loop Bound Matters

```
s = "hello"  (len=5)
substr = "lo" (len=2)

i can be: 0,1,2,3  (5-2=3, so i <= 3)
i=0: s[0:2] = "he" ≠ "lo"
i=1: s[1:3] = "el" ≠ "lo"
i=2: s[2:4] = "ll" ≠ "lo"
i=3: s[3:5] = "lo" = "lo" → return 3
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `i < len(s)` as bound | Panics on last window | Use `i <= len(s)-len(substr)` |
| Not handling empty substr | Should return 0 | Check `len(substr) == 0` first |
| `i < len(s)-len(substr)` (strict) | Misses last window | Must use `<=` |

## Algorithm

1. If `substr` is empty, return 0
2. Loop `i` from 0 to `len(s)-len(substr)` inclusive
3. If `s[i:i+len(substr)] == substr`, return `i`
4. Return -1

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [73-stringcount](../73-stringcount/README.md)
