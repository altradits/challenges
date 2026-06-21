# Skills for 14-stringspackage

## What You'll Learn

**Previous:** [13-printifnot](../13-printifnot/skills.md) | **Next:** [15-removespaces](../15-removespaces/skills.md)

**Challenge:** Write `SentenceStats(s string) (words int, chars int, hasHello bool)` that returns the word count, character count (excluding spaces), and whether the string contains the word `"hello"` (case-insensitive).

## Core Concept: The `strings` Package

### What Is It?

The `strings` package is part of Go's standard library. It provides ready-made functions for the most common string operations: searching, splitting, joining, replacing, trimming, and case conversion. You never have to write these yourself.

To use it, add it to your import block:

```go
import "strings"
```

### Searching

```go
strings.Contains("hello world", "world")  // true
strings.Contains("hello world", "xyz")    // false

strings.HasPrefix("hello.go", "hello")    // true  — starts with?
strings.HasSuffix("hello.go", ".go")      // true  — ends with?

strings.Index("hello", "ll")              // 2  — position of first match (-1 if not found)
```

### Splitting and Joining

```go
// Fields splits on any whitespace; multiple spaces are treated as one
strings.Fields("  hello   world  ")  // []string{"hello", "world"}

// Split splits on an exact separator
strings.Split("a,b,c", ",")           // []string{"a", "b", "c"}

// Join combines a slice into a single string
strings.Join([]string{"a", "b", "c"}, "-")  // "a-b-c"
```

### Replacing and Cleaning

```go
// Replace replaces the first n occurrences (use -1 for all)
strings.Replace("aabbcc", "b", "x", 1)   // "aaxbcc"  (first only)
strings.Replace("aabbcc", "b", "x", -1)  // "aaxxcc"  (all)

// ReplaceAll is shorthand for Replace(..., -1)
strings.ReplaceAll("aabbcc", "b", "x")   // "aaxxcc"

// TrimSpace removes leading and trailing whitespace
strings.TrimSpace("  hello  ")           // "hello"

// Trim removes a specific set of characters from both ends
strings.Trim("###hello###", "#")         // "hello"
```

### Case Conversion

```go
strings.ToUpper("hello")  // "HELLO"
strings.ToLower("WORLD")  // "world"
```

### Counting

```go
strings.Count("cheese", "e")  // 3
strings.Count("hello", "l")   // 2
```

## Quick Reference Table

| Function | What it does | Example |
|----------|-------------|---------|
| `strings.Contains(s, sub)` | Does `s` contain `sub`? | `Contains("hello", "ell")` → `true` |
| `strings.HasPrefix(s, pre)` | Does `s` start with `pre`? | `HasPrefix("hello", "he")` → `true` |
| `strings.HasSuffix(s, suf)` | Does `s` end with `suf`? | `HasSuffix("hello.go", ".go")` → `true` |
| `strings.Index(s, sub)` | Position of `sub` in `s` | `Index("hello", "ll")` → `2` |
| `strings.Fields(s)` | Split on any whitespace | `Fields("a  b")` → `["a","b"]` |
| `strings.Split(s, sep)` | Split on exact separator | `Split("a,b", ",")` → `["a","b"]` |
| `strings.Join(sl, sep)` | Join slice into string | `Join(["a","b"], "-")` → `"a-b"` |
| `strings.ReplaceAll(s, old, new)` | Replace all occurrences | `ReplaceAll("aabb","b","x")` → `"aaxx"` |
| `strings.TrimSpace(s)` | Remove leading/trailing spaces | `TrimSpace("  hi  ")` → `"hi"` |
| `strings.ToUpper(s)` | Convert to uppercase | `ToUpper("hi")` → `"HI"` |
| `strings.ToLower(s)` | Convert to lowercase | `ToLower("HI")` → `"hi"` |
| `strings.Count(s, sub)` | Count non-overlapping occurrences | `Count("cheese","e")` → `3` |

## Solving the Challenge

### Word Count

`strings.Fields` splits by any whitespace and ignores leading, trailing, and consecutive spaces. Its length gives the word count:

```go
import "strings"

words := len(strings.Fields(s))
```

### Character Count (Excluding Spaces)

Count all characters that are not a space. One way is to count spaces and subtract:

```go
spaces := strings.Count(s, " ")
chars := len(s) - spaces
```

Another way is `strings.ReplaceAll`:

```go
chars := len(strings.ReplaceAll(s, " ", ""))
```

### Case-Insensitive Contains

Convert the string to lowercase before checking:

```go
hasHello := strings.Contains(strings.ToLower(s), "hello")
```

### Putting It Together

```go
func SentenceStats(s string) (words int, chars int, hasHello bool) {
    words = len(strings.Fields(s))
    chars = len(strings.ReplaceAll(s, " ", ""))
    hasHello = strings.Contains(strings.ToLower(s), "hello")
    return
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting `import "strings"` | Compile error: `undefined: strings` | Always add `"strings"` to the import block |
| Using `strings.Split(s, " ")` for word count | Multiple consecutive spaces produce empty string elements, inflating the count | Use `strings.Fields` which handles all whitespace correctly |
| `strings.Contains(s, "Hello")` for case-insensitive check | Misses `"HELLO"`, `"hello"`, etc. | Use `strings.Contains(strings.ToLower(s), "hello")` |
| `len(s)` for character count | Includes spaces | Remove spaces first or subtract their count |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [15-removespaces](../15-removespaces/README.md)
