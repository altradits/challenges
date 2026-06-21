# Skills for 79-isempty

## What You'll Learn

**Previous:** [78-lastchar](../78-lastchar/skills.md) | **Next:** [80-toupper](../80-toupper/skills.md)

**Challenge:** Write a function `IsEmpty(s string) bool` that returns `true` if the string has no characters, `false` otherwise.

## Core Concept: Boolean Return Values and Empty-String Detection

### What Is a `bool`?

A `bool` is the simplest type in Go: it is either `true` or `false`. Functions that answer yes/no questions return `bool`:

```go
func IsEmpty(s string) bool {
    // must return true or false
}
```

### Two Ways to Check for Empty

**Option A — compare with `""`:**

```go
func IsEmpty(s string) bool {
    return s == ""
}
```

This is the most direct and idiomatic way. `s == ""` evaluates to `true` when the string has zero characters, `false` otherwise.

**Option B — use `len(s)`:**

```go
func IsEmpty(s string) bool {
    return len(s) == 0
}
```

Both work identically. Option A reads more naturally.

**Option C — use `for...range` (what the README encourages):**

```go
func IsEmpty(s string) bool {
    for range s {
        return false   // found at least one character → not empty
    }
    return true        // loop ran zero times → empty
}
```

This reinforces your understanding of iteration: a `for...range` on an empty string runs zero times, so only the final `return true` executes.

### Empty String vs Whitespace

These are NOT the same:

| Value | `IsEmpty` result | Why |
|-------|-----------------|-----|
| `""` | `true` | Zero characters |
| `" "` | `false` | One space character |
| `"\t"` | `false` | One tab character |

The challenge tests only for truly empty strings (zero length), not for "only whitespace".

### Returning `bool` Directly vs `if/else`

You can return the result of a comparison directly — there is no need for an `if`:

```go
// Verbose (unnecessary)
if s == "" {
    return true
} else {
    return false
}

// Concise (preferred)
return s == ""
```

Both are correct; the second form is cleaner.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `return s == " "` | Checks for one space, not empty | `return s == ""` |
| Returning `0` or `1` instead of `bool` | Type mismatch, won't compile | Return `true` / `false` |
| `if s == "" { return true }` with no other return | Compile error: missing return | Add `return false` after |

## Solving This Challenge

### Algorithm

1. Return `s == ""` directly (one line!)

Or if using the loop style:
1. `for range s { return false }`
2. `return true`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [80-toupper](../80-toupper/skills.md)
