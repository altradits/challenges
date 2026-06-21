# Skills for 15-lastchar

## What You'll Learn

**Previous:** [14-firstchar](../14-firstchar/skills.md) | **Next:** [16-isempty](../16-isempty/skills.md)

**Challenge:** Write a function `LastChar(s string) string` that returns the last character of the string, or `""` if the string is empty.

## Core Concept: `len(s)` and the Last-Index Formula

### What Is `len()`?

`len(s)` is a built-in Go function that returns the number of **bytes** in a string. For ASCII strings, that equals the number of characters:

```go
len("Hello")    // 5
len("Go")       // 2
len("")         // 0
```

### The Last-Index Formula

Strings in Go are 0-indexed. A string of length 5 has valid indices 0, 1, 2, 3, 4:

```
String:  "Hello"
Index:    0   1   2   3   4
Char:     H   e   l   l   o
```

The last valid index is always `len(s) - 1`. Using `len(s)` directly as an index causes a runtime panic:

```go
s := "Hello"
s[5]          // PANIC: index out of range [5] with length 5
s[len(s)-1]   // OK: index 4, returns 'o'
```

### Putting It Together

```go
func LastChar(s string) string {
    if s == "" {
        return ""
    }
    return string(s[len(s)-1])
}
```

Step by step:
1. Guard against empty: `s == ""` — if true, return `""`
2. `len(s)` — get the byte count
3. `len(s) - 1` — compute the last index
4. `s[len(s)-1]` — get that byte (type `byte`)
5. `string(...)` — convert the byte to a one-character string

### Alternative: Iterate and Keep Overwriting

```go
func LastChar(s string) string {
    var last rune
    for _, c := range s {
        last = c   // overwrite every time — after loop this is the last rune
    }
    if last == 0 {
        return ""
    }
    return string(last)
}
```

This approach is UTF-8 safe (works for multi-byte characters) but is slower for long strings. The index approach is fine for this challenge.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `s[len(s)]` | Panic: one past the end | Use `s[len(s)-1]` |
| No empty check | Panic on `""` because `len("")-1 = -1` | Check `s == ""` first |
| `return s[len(s)-1]` without `string()` | Returns a `byte`, not a `string` | Wrap with `string(...)` |

## Solving This Challenge

### Algorithm

1. If `s == ""`, return `""`
2. Compute `lastIndex := len(s) - 1`
3. Return `string(s[lastIndex])`

### Trace Through an Example

Input: `"Go is fun"`

- `s == ""` → false
- `len(s)` → 9
- Last index: `9 - 1 = 8`
- `s[8]` → byte value for `'n'`
- `string('n')` → `"n"`
- Return `"n"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [16-isempty](../16-isempty/README.md)
