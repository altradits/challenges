# Skills for 41-join

**Previous:** [40-split](../40-split/skills.md) | **Next:** [42-cameltosnakecase](../42-cameltosnakecase/skills.md)

**Challenge:** Implement `Join(elems []string, sep string) string` that combines a slice of strings into one string with a separator between each element.

## Core Concept: Separator-Between (Not After) Pattern

### What Is It?

Joining requires adding the separator **between** elements, not after each one. The trick: start with the first element, then for every subsequent element, prepend the separator before adding it.

### The Wrong Way First (So You Understand the Fix)

A natural impulse is:

```go
// WRONG: adds separator after every element, including the last
result := ""
for _, s := range elems {
    result += s + sep
}
// "a,b,c" becomes "a,b,c,"  — trailing separator!
```

### The Correct Pattern

```go
func Join(elems []string, sep string) string {
    if len(elems) == 0 {
        return ""
    }
    result := elems[0]           // start with the first element (no separator yet)
    for i := 1; i < len(elems); i++ {
        result += sep + elems[i] // prepend separator before each subsequent element
    }
    return result
}
```

### Why Start from Index 1?

If you start the loop at index 0 and add the separator before each element, the first element gets a leading separator:

```
sep=","  elems=["a","b","c"]
i=0: result = "," + "a" = ",a"  ← WRONG
i=1: result = ",a" + "," + "b" = ",a,b"
i=2: result = ",a,b" + "," + "c" = ",a,b,c"  ← still wrong
```

By seeding `result = elems[0]` and looping from `i=1`, the first separator only appears between element 0 and element 1.

### Edge Cases

```go
Join([]string{}, ",")          // ""    — empty slice
Join([]string{"only"}, ",")    // "only" — single element, no sep added
Join([]string{"a","b"}, "")    // "ab"  — empty separator
Join([]string{"a","b"}, ", ")  // "a, b" — multi-char separator
```

### With `strings.Builder` (More Efficient)

```go
func Join(elems []string, sep string) string {
    if len(elems) == 0 {
        return ""
    }
    var b strings.Builder
    b.WriteString(elems[0])
    for i := 1; i < len(elems); i++ {
        b.WriteString(sep)
        b.WriteString(elems[i])
    }
    return b.String()
}
```

You'll learn why `strings.Builder` is preferable for large inputs in [50-stringbuilder](../50-stringbuilder/skills.md).

### Join is the Inverse of Split

```go
parts := Split("a,b,c", ",")  // ["a", "b", "c"]
back := Join(parts, ",")       // "a,b,c"
// back == original ✓
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Adding separator after every element | Trailing separator at end | Start with `elems[0]`, loop from index 1 |
| Not checking `len(elems) == 0` | Accessing `elems[0]` panics on empty slice | Always check first |
| Starting loop at index 0 | Leading separator at start | Start loop at index 1 |

## Algorithm

1. If `len(elems) == 0`, return `""`
2. Set `result = elems[0]`
3. For `i` from `1` to `len(elems)-1`:
   - `result += sep + elems[i]`
4. Return `result`

## The Challenge

See [README.md](README.md).

**Next:** [42-cameltosnakecase](../42-cameltosnakecase/README.md)
