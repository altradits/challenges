# Skills for 69-stringreplace

**Previous:** [68-stringrepeat](../68-stringrepeat/skills.md) | **Next:** [70-stringtrim](../70-stringtrim/skills.md)

**Challenge:** Implement `Replace(s, old, new string) string` that replaces all occurrences of `old` with `new`, using `strings.Index` to find each occurrence, without using `strings.ReplaceAll`.

## Core Concept: Index-Based Replace with `strings.Index`

### What Is It?

This challenge revisits the replace-all pattern from [54-replaceall](../54-replaceall/skills.md), but now you can use `strings.Index` to jump directly to each occurrence rather than scanning character by character.

`strings.Index(s, substr)` returns the byte index of the first occurrence of `substr` in `s`, or -1 if not found.

### The Algorithm

```go
func Replace(s, old, new string) string {
    if old == "" {
        return s
    }
    var b strings.Builder
    remaining := s
    for {
        idx := strings.Index(remaining, old)
        if idx == -1 {
            b.WriteString(remaining)  // write rest, no more matches
            break
        }
        b.WriteString(remaining[:idx])  // write text before the match
        b.WriteString(new)              // write replacement
        remaining = remaining[idx+len(old):]  // advance past the match
    }
    return b.String()
}
```

### How It Works Step by Step

For `Replace("foo bar foo", "foo", "baz")`:

| Iteration | remaining | idx | Write before | Write new | New remaining |
|-----------|-----------|-----|--------------|-----------|---------------|
| 1 | "foo bar foo" | 0 | "" | "baz" | " bar foo" |
| 2 | " bar foo" | 5 | " bar " | "baz" | "" |
| 3 | "" | -1 | "" (write rest) | — | — |

Result: `"baz bar baz"` — correct.

### `strings.Index` vs Manual Scan

| Approach | When to use |
|----------|-------------|
| Manual scan (102) | No `strings` package available; learning the fundamentals |
| `strings.Index` (117) | In production Go code; cleaner and often faster |

Both produce the same result. `strings.Index` is the standard approach in real Go programs.

### Using `strings.Builder` Here

After the loop, `b.String()` returns the result without O(n²) copying. Each `b.WriteString` call appends directly to the buffer.

### The "Remaining Slice" Pattern

Instead of tracking an integer index into the original string, this pattern slices `remaining` down after each match. `remaining = remaining[idx+len(old):]` creates a new string view starting just after the match. This is clean and easy to reason about.

### Standard Library: `strings.Replace` and `strings.ReplaceAll`

```go
strings.Replace(s, old, new, n)    // replace first n occurrences (-1 = all)
strings.ReplaceAll(s, old, new)    // same as strings.Replace(s, old, new, -1)
```

After this challenge, you understand exactly what `strings.ReplaceAll` does internally.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not handling `old == ""` | `strings.Index(s, "")` returns 0, loop infinite | Guard with `if old == "" { return s }` |
| Using `s[idx+len(old):]` but forgetting `remaining` update | Never advances, infinite loop | Use `remaining = remaining[idx+len(old):]` |
| Forgetting to write the tail | Last segment (after final match) dropped | Write `remaining` when `idx == -1` |

## Algorithm

1. If `old == ""`, return `s`
2. Create `var b strings.Builder`, set `remaining = s`
3. Loop:
   - `idx = strings.Index(remaining, old)`
   - If `idx == -1`: write `remaining`, break
   - Write `remaining[:idx]`, write `new`, set `remaining = remaining[idx+len(old):]`
4. Return `b.String()`

## The Challenge

See [README.md](README.md).

**Next:** [70-stringtrim](../70-stringtrim/README.md)
