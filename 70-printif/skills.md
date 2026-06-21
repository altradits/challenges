# Skills for 70-printif

## What You'll Learn

**Previous:** [69-findsubstring](../69-findsubstring/skills.md) | **Next:** [71-printifnot](../71-printifnot/skills.md)

**Challenge:** Write `PrintIf(str string) string` that returns `"G\n"` if `len(str) >= 3` or if `str` is empty, and `"Invalid Input\n"` otherwise.

## Core Concept: `if/else` Branching and `len()` as a Condition

### What Is It?

This is your first challenge that explicitly requires an `if/else` decision. The program must choose between two possible return values based on the length of the input string. The key tool is `len(s)` combined with a comparison operator.

### How `if/else` Works in Go

```go
if condition {
    // runs when condition is true
} else {
    // runs when condition is false
}
```

There are no parentheses around the condition in Go. The braces `{}` are always required.

```go
// Correct Go style
if len(str) >= 3 {
    return "G\n"
}

// Syntax error: condition needs no parens, but braces are mandatory
if (len(str) >= 3)   // missing braces
```

### Using `len()` in a Condition

`len(str)` returns an `int`. You can use it directly in a comparison:

```go
len("ab")    // 2
len("abc")   // 3
len("")       // 0

len("abc") >= 3   // true
len("ab") >= 3    // false
len("") >= 3      // false — but the challenge has a special rule for empty
```

### Reading the Requirements Carefully

The challenge says:
- If `len(str) >= 3` -> return `"G\n"`
- If the string is **empty** -> return `"G\n"` (special case)
- Otherwise -> return `"Invalid Input\n"`

This means you need to check for empty string OR length >= 3:

```go
func PrintIf(str string) string {
    if str == "" || len(str) >= 3 {
        return "G\n"
    }
    return "Invalid Input\n"
}
```

The `||` operator means "or" — the condition is true if either side is true.

### Alternative: Checking the Complement

You can also think of it as: return `"Invalid Input\n"` only when the string is non-empty AND shorter than 3.

```go
func PrintIf(str string) string {
    if len(str) > 0 && len(str) < 3 {
        return "Invalid Input\n"
    }
    return "G\n"
}
```

Both forms are correct.

### Diagram: Decision Tree

```
Input str
    |
    v
Is str == "" ?
    |         \
   YES         NO
    |           |
 return "G\n"   Is len(str) >= 3 ?
                    |           \
                   YES           NO
                    |             |
                return "G\n"  return "Invalid Input\n"
```

### Boolean Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `&&` | AND — both must be true | `len(s) > 0 && len(s) < 3` |
| `\|\|` | OR — at least one must be true | `s == "" \|\| len(s) >= 3` |
| `!` | NOT — invert a boolean | `!isEmpty` |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting the empty string special case | Empty string returns `"Invalid Input\n"` instead of `"G\n"` | Add `str == ""` to the condition |
| `return "G"` without `\n` | Missing newline — tests will fail | Use `"G\n"` |
| `if len(str) = 3` (single `=`) | Assignment, not comparison — syntax error | Use `>=` or `==` |

## Solving This Challenge

### Algorithm

1. If `str` is empty OR `len(str) >= 3`, return `"G\n"`
2. Otherwise return `"Invalid Input\n"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [71-printifnot](../71-printifnot/README.md)
