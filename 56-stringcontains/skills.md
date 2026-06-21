# Skills for 56-stringcontains

**Previous:** [55-stringtrim](../55-stringtrim/skills.md) | **Next:** [57-stringindex](../57-stringindex/skills.md)

**Challenge:** Implement `Contains(s, substr string) bool` that returns true if `substr` is found anywhere in `s`, without using `strings.Contains`.

## Core Concept: Contains is Just Index != -1

### What Is It?

`strings.Contains(s, substr)` is implemented as: `strings.Index(s, substr) >= 0`. Once you have a working `Index` function, `Contains` is a one-liner.

For this challenge, you implement both — manually finding whether the substring appears anywhere in `s`.

### The Implementation

Using the sliding window from [38-findsubstring](../38-findsubstring/skills.md):

```go
func Contains(s, substr string) bool {
    if substr == "" {
        return true  // empty string is "contained" in everything
    }
    if len(substr) > len(s) {
        return false
    }
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return true
        }
    }
    return false
}
```

Or by delegating to your own `Index` function:

```go
func Index(s, substr string) int {
    if substr == "" { return 0 }
    if len(substr) > len(s) { return -1 }
    for i := 0; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return i
        }
    }
    return -1
}

func Contains(s, substr string) bool {
    return Index(s, substr) >= 0
}
```

### The Standard Library Functions

```go
import "strings"

strings.Contains(s, substr)      // true if substr appears anywhere
strings.ContainsAny(s, chars)    // true if any char in `chars` appears in s
strings.ContainsRune(s, r)       // true if specific rune r appears in s
```

```go
strings.ContainsAny("hello", "aeiou")  // true — 'e' and 'o' are vowels
strings.ContainsAny("xyz", "aeiou")    // false — no vowels
strings.ContainsRune("hello", 'e')     // true — 'e' is in "hello"
```

### ContainsAny vs Contains

- `Contains("hello", "ell")` — checks for the substring `"ell"` as a unit
- `ContainsAny("hello", "ell")` — checks if ANY of `'e'`, `'l'` appears (returns true if any one matches)

```go
strings.Contains("xyz", "ell")     // false — "ell" is not in "xyz"
strings.ContainsAny("xyz", "ell")  // false — none of 'e', 'l', 'l' in "xyz"

strings.Contains("hello", "ell")     // true — "ell" appears in "hello"
strings.ContainsAny("hello", "xyz")  // false — none of 'x', 'y', 'z' in "hello"
strings.ContainsAny("hello", "xyo")  // true — 'o' appears in "hello"
```

### Special Cases

```go
Contains("hello", "")  // true  — empty substring is always "contained"
Contains("", "abc")    // false — can't find non-empty substring in empty string
Contains("", "")       // true  — empty in empty
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Returning false for empty substr | `Contains("x", "")` should be true | Check `substr == ""` → return true |
| Not checking length bounds | `s[i:i+len(substr)]` panics if substr longer than remaining | Use `i <= len(s)-len(substr)` |

## Algorithm

1. If `substr == ""`, return `true`
2. If `len(substr) > len(s)`, return `false`
3. For `i` from `0` to `len(s)-len(substr)`:
   - If `s[i:i+len(substr)] == substr`, return `true`
4. Return `false`

## The Challenge

See [README.md](README.md).

**Next:** [57-stringindex](../57-stringindex/README.md)
