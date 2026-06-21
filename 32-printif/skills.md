# Skills for 32-printif

## What You'll Learn

**Previous:** [31-replacechar](../31-replacechar/skills.md) | **Next:** [33-printifnot](../33-printifnot/skills.md)

**Challenge:** Write a function `PrintIf(str string) string` that returns `"G\n"` if the string length is 0 or >= 3, and `"Invalid Input\n"` otherwise.

## Core Concept: Conditional Return Based on `len()` — The `||` Operator

### The Decision Rule

```
length == 0  → "G\n"
length == 1  → "Invalid Input\n"
length == 2  → "Invalid Input\n"
length >= 3  → "G\n"
```

In other words, return `"G\n"` when `len(str) == 0 OR len(str) >= 3`.

### The Implementation

```go
func PrintIf(str string) string {
    if len(str) == 0 || len(str) >= 3 {
        return "G\n"
    }
    return "Invalid Input\n"
}
```

Or equivalently, check the invalid case first:

```go
func PrintIf(str string) string {
    if len(str) >= 1 && len(str) < 3 {
        return "Invalid Input\n"
    }
    return "G\n"
}
```

### `len()` Returns a Count You Can Compare

From [15-lastchar](../15-lastchar/skills.md) you know `len(s)` returns an integer. You can use that integer in any comparison:

```go
len(str) == 0    // empty
len(str) >= 3    // three or more characters
len(str) < 3     // fewer than three characters
```

### `||` (Logical OR)

`condition1 || condition2` is `true` when at least one condition is true:

```go
len(str) == 0 || len(str) >= 3
```

- Empty string (len=0): first condition true → result is true → return `"G\n"`
- len=1 or len=2: both false → result is false → return `"Invalid Input\n"`
- len=3 or more: second condition true → result is true → return `"G\n"`

### The `\n` Newline

The return values include `\n` (a newline character). This is part of the literal string — include it exactly as shown in the README. The caller uses `fmt.Print` (not `fmt.Println`) so the newline must be embedded in the returned string.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `len(str) > 3` instead of `>= 3` | Excludes length-3 strings | Use `>= 3` |
| Forgetting `\n` in the return strings | Test output mismatches | Return `"G\n"` and `"Invalid Input\n"` |
| `if len(str) == 0 && len(str) >= 3` | Impossible condition — always false | Use `||`, not `&&` |

## Solving This Challenge

### Algorithm

1. If `len(str) == 0 || len(str) >= 3`: return `"G\n"`
2. Otherwise: return `"Invalid Input\n"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [33-printifnot](../33-printifnot/README.md)
