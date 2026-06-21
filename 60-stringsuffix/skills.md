# Skills for 60-stringsuffix

## What You'll Learn

**Previous:** [59-stringprefix](../59-stringprefix/skills.md) | **Next:** [61-stringfield](../61-stringfield/skills.md)

**Challenge:** Implement `HasSuffix(s, suffix string) bool` without using `strings.HasSuffix`.

## Core Concept: `strings.HasSuffix` and `strings.TrimSuffix`

### The Built-in Functions

```go
import "strings"

strings.HasSuffix("hello world", "world")  // true
strings.HasSuffix("hello world", "hello")  // false
strings.HasSuffix("hello", "")             // true
strings.HasSuffix("", "hello")             // false

strings.TrimSuffix("hello.go", ".go")      // "hello"
strings.TrimSuffix("hello.js", ".go")      // "hello.js" (unchanged)
```

### How to Implement `HasSuffix` Manually

```go
func HasSuffix(s, suffix string) bool {
    if len(suffix) == 0 {
        return true
    }
    if len(s) < len(suffix) {
        return false
    }
    return s[len(s)-len(suffix):] == suffix
}
```

The key expression is `s[len(s)-len(suffix):]` — this slices the **last** `len(suffix)` bytes.

### Prefix vs Suffix — Side-by-Side

```go
// Prefix: check the FRONT
s[:len(prefix)] == prefix

// Suffix: check the BACK
s[len(s)-len(suffix):] == suffix
```

### Common Uses

- File extension check: `strings.HasSuffix(filename, ".go")`
- Protocol check: `strings.HasSuffix(url, "/api")`
- Strip extension: `strings.TrimSuffix("main.go", ".go")` → `"main"`

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `s[len(s)-len(suffix)]` | Gets a byte, not a string | Add `:` → `s[len(s)-len(suffix):]` |
| Not guarding `len(s) < len(suffix)` | Negative index panics | Always check lengths first |
| Comparing byte instead of string | Type mismatch | Use slice `s[...]`, not `s[i]` |

## Algorithm

1. If `len(suffix) == 0`, return `true`
2. If `len(s) < len(suffix)`, return `false`
3. Return `s[len(s)-len(suffix):]` == `suffix`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [61-stringfield](../61-stringfield/README.md)
