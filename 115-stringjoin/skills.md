# Skills for 115-stringjoin

**Previous:** [114-stringsplit](../114-stringsplit/README.md) | **Next:** [116-stringrepeat](../116-stringrepeat/README.md)

**Challenge:** Implement `Join(words []string, sep string) string` that joins a slice of strings with a separator, using `strings.Builder` for efficiency, without using `strings.Join`.

## Core Concept: `strings.Builder` for Efficient Joining

### The Pattern

This challenge combines two skills from the 100-series:
1. The "separator between, not after" pattern from [104-join](../104-join/skills.md)
2. The `strings.Builder` efficiency from [113-stringbuilder](../113-stringbuilder/skills.md)

```go
func Join(words []string, sep string) string {
    if len(words) == 0 {
        return ""
    }
    var b strings.Builder
    b.WriteString(words[0])           // first element, no separator
    for i := 1; i < len(words); i++ {
        b.WriteString(sep)            // separator before every element after the first
        b.WriteString(words[i])
    }
    return b.String()
}
```

### Why `strings.Builder` Over `+`?

For joining `n` words each of length `L`:
- String `+` approach: copies words[0] (L bytes), then words[0]+sep+words[1] (2L+1 bytes), then 3L+2... Total: O(n²·L)
- `strings.Builder` approach: each write appends to the buffer. Total: O(n·L)

For large word lists, `strings.Builder` is dramatically faster.

### With `b.Grow` Optimization

If you know the approximate total size upfront, hint to `Builder` to pre-allocate:

```go
func Join(words []string, sep string) string {
    if len(words) == 0 {
        return ""
    }
    // estimate total length
    total := len(sep) * (len(words) - 1)
    for _, w := range words {
        total += len(w)
    }
    var b strings.Builder
    b.Grow(total)  // pre-allocate to avoid buffer doublings
    
    b.WriteString(words[0])
    for i := 1; i < len(words); i++ {
        b.WriteString(sep)
        b.WriteString(words[i])
    }
    return b.String()
}
```

`b.Grow(n)` is optional but good practice when you know the size.

### Testing All Edge Cases

```go
Join([]string{"a","b","c"}, "-")   // "a-b-c"
Join([]string{"hello","world"}, " ") // "hello world"
Join([]string{}, ",")               // ""
Join([]string{"single"}, "|")       // "single"  (no sep added)
Join([]string{"a","b"}, "")         // "ab"      (empty sep = direct concat)
Join([]string{"a","b","c"}, ", ")   // "a, b, c" (multi-char sep)
```

### Comparison: Manual Join (104) vs Standard Join (115)

The logic is identical; the difference is using `strings.Builder` instead of `+=`:

```go
// 104 approach (+=, fine for small inputs):
result := words[0]
for i := 1; i < len(words); i++ {
    result += sep + words[i]
}

// 115 approach (Builder, efficient for large inputs):
var b strings.Builder
b.WriteString(words[0])
for i := 1; i < len(words); i++ {
    b.WriteString(sep)
    b.WriteString(words[i])
}
result := b.String()
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting `len(words) == 0` check | `words[0]` panics on empty slice | Check length first |
| Starting loop at `i=0` | Leading separator before first word | Start at `i=1`, seed with `words[0]` |
| Using `+=` in the loop | O(n²) for large word lists | Use `b.WriteString` |

## Algorithm

1. If `len(words) == 0`, return `""`
2. Declare `var b strings.Builder`
3. `b.WriteString(words[0])`
4. For `i` from `1` to `len(words)-1`: `b.WriteString(sep)`, `b.WriteString(words[i])`
5. Return `b.String()`

## The Challenge

See [README.md](README.md).

**Next:** [116-stringrepeat](../116-stringrepeat/README.md)
