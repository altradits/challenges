# Skills for 122-stringprefix

## What You'll Learn

**Previous:** [121-stringcount](../121-stringcount/skills.md) | **Next:** [123-stringsuffix](../123-stringsuffix/skills.md)

**Challenge:** Implement `HasPrefix(s, prefix string) bool` without using `strings.HasPrefix`.

## Core Concept: `strings.HasPrefix` and `strings.TrimPrefix`

### The Built-in Functions

```go
import "strings"

strings.HasPrefix("hello world", "hello")  // true
strings.HasPrefix("hello world", "world")  // false
strings.HasPrefix("hello", "")             // true (empty prefix always matches)
strings.HasPrefix("", "hello")             // false

strings.TrimPrefix("hello world", "hello") // " world"
strings.TrimPrefix("hello world", "xyz")   // "hello world" (unchanged if no prefix)
```

### How to Implement `HasPrefix` Manually

```go
func HasPrefix(s, prefix string) bool {
    if len(prefix) == 0 {
        return true  // empty prefix is always present
    }
    if len(s) < len(prefix) {
        return false  // s is too short
    }
    return s[:len(prefix)] == prefix
}
```

Logic: extract the first `len(prefix)` bytes of `s` and compare.

### When to Use These

- `HasPrefix` — validate format: `"BTC:..."`, `"0x..."`, `"https://..."`
- `TrimPrefix` — strip a known leader: `strings.TrimPrefix(s, "0x")` → hex digits only

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Checking `s[0] == prefix[0]` | Only checks one byte | Compare full slice `s[:len(prefix)]` |
| Not checking `len(s) < len(prefix)` | Panics on short strings | Guard first |
| Returning false for empty prefix | Contract says true | Check `len(prefix) == 0` first |

## Algorithm

1. If `len(prefix) == 0`, return `true`
2. If `len(s) < len(prefix)`, return `false`
3. Return `s[:len(prefix)] == prefix`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [123-stringsuffix](../123-stringsuffix/skills.md)
