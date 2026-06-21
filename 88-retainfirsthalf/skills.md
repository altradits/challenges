# Skills for 88-retainfirsthalf

## What You'll Learn

**Previous:** [87-removespaces](../87-removespaces/skills.md) | **Next:** [89-reversestring](../89-reversestring/skills.md)

**Challenge:** Write `RetainFirstHalf(str string) string` that returns the first half of the string. Round down for odd lengths.

## Core Concept: String Slicing with `len()` and Integer Division

### What Is It?

This challenge combines two ideas: measuring a string's length with `len()`, and extracting a portion of a string with slice notation `s[:n]`. The new concept is using **integer division** to compute the midpoint — in Go, dividing two integers always truncates toward zero (rounds down).

### Integer Division in Go

When you divide two integers in Go, the result is always an integer — no decimal part:

```go
10 / 2   // 5
11 / 2   // 5 (not 5.5 — truncated)
7 / 2    // 3 (not 3.5)
1 / 2    // 0
```

This is exactly what the challenge wants: "round down" for odd lengths.

### String Slicing Recap

`s[:n]` returns the first `n` characters:

```go
s := "Hello World"
s[:5]   // "Hello"
s[:0]   // ""
s[:]    // "Hello World" (the whole string)
```

### Combining `len` and Integer Division

```go
func RetainFirstHalf(str string) string {
    return str[:len(str)/2]
}
```

Walk-through for different inputs:

```
"This is the 1st halfThis is the 2nd half"  len=40  40/2=20  -> str[:20] = "This is the 1st half"
"A"       len=1  1/2=0   -> str[:0] = ""  ... wait, but the challenge says return "A" for single char
""        len=0  0/2=0   -> str[:0] = ""  correct
"Hello World"  len=11  11/2=5  -> str[:5] = "Hello"  correct
```

### Handling the Single-Character Case

The README says: "If the string length is equal to one, return the string."

For `"A"`: `len("A") / 2 = 0`, which would return `""`. But the spec wants `"A"`.

So you need a special case:

```go
func RetainFirstHalf(str string) string {
    if len(str) <= 1 {
        return str
    }
    return str[:len(str)/2]
}
```

### Diagram: How Slicing Works

```
"Hello World"
 0 1 2 3 4 5 6 7 8 9 10
 H e l l o   W o r l d

len("Hello World") = 11
11 / 2 = 5

str[:5] = characters at indices 0,1,2,3,4 = "Hello"
```

```
"Hello"
 0 1 2 3 4
 H e l l o

len("Hello") = 5
5 / 2 = 2

str[:2] = "He"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `len(str) / 2.0` or converting to float | Unnecessary and wrong type | Just use `len(str) / 2` (integer division) |
| Forgetting the single-character special case | `"A"` returns `""` instead of `"A"` | Add `if len(str) <= 1 { return str }` |
| `str[0 : len/2]` — forgetting the colon | Syntax error | Use `str[:len(str)/2]` |
| Using `len` on result of `range` | `range` is for iteration, not length | Use `len(str)` for the count |

## Solving This Challenge

### Algorithm

1. If `str` has 0 or 1 characters, return `str` unchanged
2. Compute `half := len(str) / 2` (integer division rounds down automatically)
3. Return `str[:half]`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [89-reversestring](../89-reversestring/README.md)
