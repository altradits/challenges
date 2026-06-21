# Skills for 29-firstchar

## What You'll Learn

**Previous:** [28-stringlength](../28-stringlength/skills.md) | **Next:** [30-lastchar](../30-lastchar/skills.md)

**Challenge:** Write a function `FirstChar(s string) string` that returns the first character of the string, or `""` if the string is empty.

## Core Concept: String Indexing and Type Conversion

### What Is It?

Go lets you access a single character from a string with bracket notation: `s[0]`. The catch is that `s[0]` returns a **byte** (type `uint8`), not a string. You need to convert it to get printable text back.

### Two Ways to Get the First Character

**Option A — byte indexing (simple, works for ASCII):**

```go
func FirstChar(s string) string {
    if s == "" {
        return ""
    }
    return string(s[0])   // s[0] is a byte; string() converts it to text
}
```

**Option B — for...range (safer for Unicode):**

```go
func FirstChar(s string) string {
    for _, c := range s {
        return string(c)   // return immediately on the first rune
    }
    return ""   // loop never ran — string was empty
}
```

Option B handles multi-byte characters (like `é`, `中`) correctly because `for...range` produces runes, not bytes. For this challenge either approach passes the tests.

### The Type Hierarchy

Understanding these three types prevents confusion:

| Type | What it is | Example |
|------|-----------|---------|
| `byte` (`uint8`) | Raw 8-bit value | `s[0]` → `72` for `'H'` |
| `rune` (`int32`) | Unicode code point | First value from `for...range` |
| `string` | Sequence of bytes | `"Hello"` |

Converting any of these with `string()` gives you the character as text:

```go
b := s[0]              // byte 72
fmt.Println(b)         // 72  (the number)
fmt.Println(string(b)) // "H" (the character)
```

### The Empty String Guard

If `s` is empty and you try `s[0]`, Go panics at runtime:

```
runtime error: index out of range [0] with length 0
```

Always check for empty before indexing:

```go
if s == "" {
    return ""
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `return s[0]` | Returns `byte` (uint8), not `string` | `return string(s[0])` |
| No empty check before `s[0]` | Panic on empty input | Check `s == ""` first |
| `return s[:0]` | Returns `""` (the empty prefix, not the first char) | Use `s[:1]` for first character |

## Solving This Challenge

### Algorithm

1. If `s == ""`, return `""`
2. Otherwise return `string(s[0])`

### Trace Through Examples

Input: `"Hello"`
- `s == ""` → false
- `s[0]` → byte `72` (letter H)
- `string(72)` → `"H"`
- Return `"H"`

Input: `""`
- `s == ""` → true
- Return `""`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [30-lastchar](../30-lastchar/README.md)
