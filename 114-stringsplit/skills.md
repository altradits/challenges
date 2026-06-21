# Skills for 114-stringsplit

**Previous:** [113-stringbuilder](../113-stringbuilder/README.md) | **Next:** [115-stringjoin](../115-stringjoin/README.md)

**Challenge:** Implement `SplitWords(s string) []string` that splits a string into words, handling multiple consecutive spaces, leading/trailing spaces, and returning an empty slice for empty input — without using `strings.Split`.

## Core Concept: `strings.Fields` — Split on Any Whitespace

### What Is `strings.Fields`?

`strings.Fields(s)` splits a string into a slice of substrings separated by **any whitespace** (spaces, tabs, newlines). It:
- Handles multiple consecutive spaces automatically
- Ignores leading and trailing whitespace
- Returns an empty slice for an empty or all-whitespace string

```go
import "strings"

fmt.Println(strings.Fields("Hello World"))
// [Hello World]

fmt.Println(strings.Fields("  Go   is  fun  "))
// [Go is fun]

fmt.Println(strings.Fields(""))
// []  (empty slice)

fmt.Println(strings.Fields("   "))
// []  (all whitespace → empty slice)

fmt.Println(strings.Fields("single"))
// [single]
```

### How `strings.Fields` Differs from `strings.Split`

| Function | Behavior | Example |
|----------|----------|---------|
| `strings.Split(s, " ")` | Splits on exact string; consecutive spaces create empty elements | `"a  b"` → `["a", "", "b"]` |
| `strings.Fields(s)` | Splits on any whitespace; runs of spaces treated as one | `"a  b"` → `["a", "b"]` |

For word splitting (this challenge), `strings.Fields` is almost always what you want.

### The Implementation

Since the constraint is "do not use `strings.Split`" (not "do not use `strings.Fields`"), you can use `strings.Fields` directly:

```go
func SplitWords(s string) []string {
    return strings.Fields(s)
}
```

But if you want to understand the manual approach (and `strings.Fields` is also disallowed), here is how to implement it:

```go
func SplitWords(s string) []string {
    var result []string
    inWord := false
    start := 0

    for i, c := range s {
        isSpace := c == ' ' || c == '\t' || c == '\n'
        if !isSpace && !inWord {
            start = i   // beginning of a new word
            inWord = true
        } else if isSpace && inWord {
            result = append(result, s[start:i])  // end of word
            inWord = false
        }
    }
    if inWord {
        result = append(result, s[start:])  // last word
    }
    return result
}
```

### `strings.Fields` vs `strings.FieldsFunc`

`strings.FieldsFunc(s, f)` is a generalization: it splits wherever `f(c)` returns true.

```go
// Split on commas or semicolons:
parts := strings.FieldsFunc("a,b;c,d", func(c rune) bool {
    return c == ',' || c == ';'
})
// ["a", "b", "c", "d"]
```

You'll use `FieldsFunc` in [124-stringfield](../124-stringfield/skills.md).

### Edge Cases

```go
strings.Fields("")        // []        — empty slice, not [""]
strings.Fields("  ")      // []        — all whitespace
strings.Fields("a")       // ["a"]     — single word
strings.Fields("a b")     // ["a","b"] — two words
```

Note: `strings.Split("", ",")` returns `[""]` (a slice with one empty string), but `strings.Fields("")` returns `[]`. This is an important behavioral difference.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `strings.Split(s, " ")` | Extra empty strings for multiple spaces | Use `strings.Fields` |
| Not handling empty input | Returning `nil` vs `[]string{}` | `strings.Fields` returns `nil` for empty input — check if the tests accept `nil` |
| Forgetting trailing word in manual version | Last word missed | Add post-loop check `if inWord { append last word }` |

## Algorithm (using `strings.Fields`)

1. Return `strings.Fields(s)` directly

## The Challenge

See [README.md](README.md).

**Next:** [115-stringjoin](../115-stringjoin/README.md)
