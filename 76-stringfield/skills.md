# Skills for 76-stringfield

## What You'll Learn

**Previous:** [75-stringsuffix](../75-stringsuffix/skills.md) | **Next:** [77-stringmap](../77-stringmap/skills.md)

**Challenge:** Implement `Split(s, sep string) []string` — split on a custom delimiter.

## Core Concept: `strings.Fields` vs `strings.Split` vs `strings.FieldsFunc`

### Three Ways to Split in Go

```go
import "strings"

// 1. Split on ANY whitespace (spaces, tabs, newlines)
strings.Fields("  hello   world  ")
// → ["hello", "world"]  — ignores empty fields

// 2. Split on an EXACT separator
strings.Split("a,b,,c", ",")
// → ["a", "b", "", "c"]  — preserves empty fields

// 3. Split on a CUSTOM condition (function-based)
strings.FieldsFunc("a1b2c3", func(r rune) bool {
    return r >= '0' && r <= '9'
})
// → ["a", "b", "c"]  — splits wherever the function returns true
```

### Key Differences

| Function | Separator | Empty fields |
|----------|-----------|-------------|
| `strings.Fields` | Any whitespace | Dropped |
| `strings.Split` | Exact string | Kept |
| `strings.SplitN` | Exact string, max N parts | Kept |
| `strings.FieldsFunc` | Custom predicate | Dropped |

### How to Implement `Split` Manually

```go
func Split(s, sep string) []string {
    result := []string{}
    for {
        idx := strings.Index(s, sep)
        if idx == -1 {
            result = append(result, s)
            break
        }
        result = append(result, s[:idx])
        s = s[idx+len(sep):]
    }
    return result
}
```

### When `strings.Fields` vs `strings.Split`

- Parsing user input with irregular spacing → `strings.Fields`
- Parsing CSV or delimited data → `strings.Split`
- Parsing mixed content → `strings.FieldsFunc`

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `Fields` when empty fields matter | Drops `""` entries | Use `strings.Split` |
| Not appending the last segment | Loses everything after last sep | Append `s` after the loop |
| Using `Split` for whitespace splitting | Doesn't collapse multiple spaces | Use `Fields` |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [77-stringmap](../77-stringmap/README.md)
