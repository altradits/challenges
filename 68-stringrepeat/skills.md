# Skills for 68-stringrepeat

**Previous:** [67-stringjoin](../67-stringjoin/skills.md) | **Next:** [69-stringreplace](../69-stringreplace/skills.md)

**Challenge:** Implement `Repeat(s string, n int) string` that returns `s` repeated `n` times, using `strings.Builder` for efficiency, without using `strings.Repeat`.

## Core Concept: Loop + `strings.Builder` for Repetition

### What Is It?

Repeating a string `n` times means writing the same string into a builder exactly `n` times in a loop. The key insight: you know the exact output size in advance (`len(s) * n`), which lets you pre-allocate with `b.Grow`.

### The Implementation

```go
func Repeat(s string, n int) string {
    if n <= 0 || s == "" {
        return ""
    }
    var b strings.Builder
    b.Grow(len(s) * n)     // pre-allocate exact size
    for i := 0; i < n; i++ {
        b.WriteString(s)
    }
    return b.String()
}
```

### Why Pre-Allocate with `b.Grow`?

Without `b.Grow`, `strings.Builder` starts with a small buffer and doubles when it runs out. For `Repeat("abc", 1000)`, that means multiple reallocations. With `b.Grow(3 * 1000)`, a single allocation is made upfront — no reallocation at all.

```
Without Grow: allocations at 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096 bytes
With Grow(3000): single allocation of 3000 bytes
```

### Tracing `Repeat("ab", 4)`

`b.Grow(8)` — allocate 8 bytes

| Iteration | WriteString | Buffer content |
|-----------|------------|----------------|
| i=0 | "ab" | "ab" |
| i=1 | "ab" | "abab" |
| i=2 | "ab" | "ababab" |
| i=3 | "ab" | "abababab" |

Return: `"abababab"` — correct.

### Edge Cases

```go
Repeat("abc", 0)   // ""  — zero repetitions
Repeat("abc", -1)  // ""  — negative: treat as zero
Repeat("", 5)      // ""  — empty string repeated is still empty
Repeat("x", 1)     // "x" — one repetition = original
```

### The Standard Library Version

`strings.Repeat(s, n)` does exactly this. After completing this challenge, you know exactly how it's implemented internally.

### Why Not Just Use `+=`?

```go
// BAD for large n:
result := ""
for i := 0; i < n; i++ {
    result += s  // each iteration copies the entire accumulated string
}
// For n=1000, s="abc": copies 0+3+6+9+...+2997 ≈ 1.5M bytes total
```

With `strings.Builder`, only `len(s)*n` bytes are written total.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not checking `n <= 0` | Negative n causes no loop; result is `""` which is correct, but return early explicitly | Check `if n <= 0 { return "" }` |
| Forgetting `b.Grow` | Works correctly, just slower | Add `b.Grow(len(s) * n)` before the loop |
| Using `+=` | O(n²) for large n | Use `b.WriteString` |

## Algorithm

1. If `n <= 0` or `s == ""`, return `""`
2. `var b strings.Builder`
3. `b.Grow(len(s) * n)`
4. Loop `n` times: `b.WriteString(s)`
5. Return `b.String()`

## The Challenge

See [README.md](README.md).

**Next:** [69-stringreplace](../69-stringreplace/README.md)
