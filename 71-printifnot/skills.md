# Skills for 71-printifnot

## What You'll Learn

**Previous:** [70-printif](../70-printif/skills.md) | **Next:** [72-removespaces](../72-removespaces/skills.md)

**Challenge:** Write `PrintIfNot(str string) string` that returns `"G\n"` if `len(str) < 3` (including empty), and `"Invalid Input\n"` otherwise.

## Core Concept: Logical Negation and Inverted Conditions

### What Is It?

`PrintIfNot` is the exact mirror image of `PrintIf`. The conditions are flipped. This challenge teaches you that the same logic can be expressed in multiple ways: using `<` instead of `>=`, using `!` (not) to invert a condition, or simply swapping the branches of an `if/else`.

### The `!` (NOT) Operator

In Go, `!` inverts a boolean value:

```go
!true   // false
!false  // true

isEmpty := len(str) == 0
!isEmpty  // true when str is NOT empty
```

### Three Equivalent Ways to Write `PrintIfNot`

**Option 1 — Direct condition:**

```go
func PrintIfNot(str string) string {
    if len(str) < 3 {
        return "G\n"
    }
    return "Invalid Input\n"
}
```

**Option 2 — Using `!` to negate the PrintIf condition:**

```go
func PrintIfNot(str string) string {
    if !(str == "" || len(str) >= 3) {
        return "Invalid Input\n"
    }
    return "G\n"
}
```

**Option 3 — Swap the branches of PrintIf:**

`PrintIf` returns `"G\n"` when `len >= 3` or empty. `PrintIfNot` returns `"G\n"` when `len < 3`. They return opposite results for `len >= 3`, and the same result for empty (`"G\n"`). The cleanest solution is option 1.

### Comparing PrintIf vs PrintIfNot

| Input | `len` | `PrintIf` returns | `PrintIfNot` returns |
|-------|-------|-------------------|----------------------|
| `""` | 0 | `"G\n"` | `"G\n"` |
| `"a"` | 1 | `"Invalid Input\n"` | `"G\n"` |
| `"ab"` | 2 | `"Invalid Input\n"` | `"G\n"` |
| `"abc"` | 3 | `"G\n"` | `"Invalid Input\n"` |
| `"abcdefz"` | 7 | `"G\n"` | `"Invalid Input\n"` |

### De Morgan's Law (Useful for Complex Conditions)

When you negate a compound condition, `&&` and `||` flip:

```
!(A || B)  ==  !A && !B
!(A && B)  ==  !A || !B
```

Example:
```go
// Original condition from PrintIf
str == "" || len(str) >= 3

// Negated (De Morgan's):
str != "" && len(str) < 3
```

So `PrintIfNot` could also be written as:
```go
if str != "" && len(str) < 3 {
    return "G\n"
}
return "Invalid Input\n"
```

But the simplest form remains: `if len(str) < 3`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Copying PrintIf without flipping the logic | Returns same results as PrintIf | The condition should be `< 3`, not `>= 3` |
| `if len(str) <= 3` (less than or equal) | Also includes `len == 3`, which should return `"Invalid Input\n"` | Use strictly `< 3` |
| Forgetting that empty string (len=0) returns `"G\n"` | `0 < 3` is `true`, so empty is handled automatically — no special case needed | No extra check required |

## Solving This Challenge

### Algorithm

1. If `len(str) < 3` (which includes empty), return `"G\n"`
2. Otherwise return `"Invalid Input\n"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [72-removespaces](../72-removespaces/README.md)
