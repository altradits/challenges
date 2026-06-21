# Skills for 06-checknumber

## What You'll Learn

**Previous:** [07-functions](../07-functions/skills.md) | **Next:** [08-count-character](../08-count-character/skills.md)

**Challenge:** Write a function `CheckNumber(arg string) bool` that returns `true` if the string contains any digit character, `false` otherwise.

## Core Concept: `for...range` on a String + Rune Comparison

### What Is It?

This is your first challenge where you write a *function* (not just a `main`) and where you must inspect the individual characters of a string. To do that you need two new tools: the `for...range` loop over a string, and rune comparisons.

### How Strings Work in Go

A Go string is a sequence of bytes encoded in UTF-8. When you iterate over a string with `for...range`, Go automatically decodes each character into a **rune** (Go's name for a Unicode code point, type `int32`).

```go
s := "Hi9"
for i, r := range s {
    // i = byte index (0, 1, 2)
    // r = rune value ('H', 'i', '9')
}
```

### Checking Whether a Rune Is a Digit

A digit character in ASCII is any rune between `'0'` and `'9'` inclusive. You can compare runes directly:

```go
r >= '0' && r <= '9'
```

This works because rune literals like `'0'` and `'9'` are just integer constants (48 and 57). The comparison checks whether the rune's integer value falls in the digit range.

### Step-by-Step: Writing `CheckNumber`

```go
func CheckNumber(arg string) bool {
    for _, r := range arg {
        if r >= '0' && r <= '9' {
            return true   // found a digit, stop immediately
        }
    }
    return false   // scanned the whole string, no digit found
}
```

Walk-through:
1. `for _, r := range arg` — iterate over every rune in `arg`; we use `_` for the index because we don't need it
2. `if r >= '0' && r <= '9'` — check if the current rune is a digit
3. `return true` inside the loop — early return as soon as a digit is found
4. `return false` after the loop — only reached if no digit was found

### Pseudocode

```
function CheckNumber(arg):
    for each character r in arg:
        if r is between '0' and '9':
            return true
    return false
```

### Diagram: Iterating "Hello1"

```
"Hello1"
 H  e  l  l  o  1
 |  |  |  |  |  |
 72 101 108 108 111 49

Is 72 >= 48 && 72 <= 57?  No
Is 101 >= 48 && 101 <= 57? No
Is 108 >= 48 && 108 <= 57? No
Is 108 >= 48 && 108 <= 57? No
Is 111 >= 48 && 111 <= 57? No
Is 49 >= 48 && 49 <= 57?  YES -> return true
```

### Alternative: Using `unicode.IsDigit`

The `unicode` package has a cleaner function for this:

```go
import "unicode"

func CheckNumber(arg string) bool {
    for _, r := range arg {
        if unicode.IsDigit(r) {
            return true
        }
    }
    return false
}
```

`unicode.IsDigit` handles non-ASCII digits too (e.g., Arabic-Indic numerals). For ASCII-only input the range comparison is fine.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `arg[i]` instead of `for...range` | `arg[i]` returns a `byte`, not a `rune`; works for ASCII but breaks on multi-byte characters | Use `for _, r := range arg` |
| Comparing `r >= 0 && r <= 9` (no quotes) | Compares against integer 0 and 9, not digit characters | Use `'0'` and `'9'` with single quotes |
| Forgetting the final `return false` | Compile error: missing return | Every code path must return a value |
| `return false` inside the loop on non-match | Returns false on the first non-digit, never finishing the scan | Put `return false` only after the loop ends |

## Solving This Challenge

### Algorithm

1. Loop over every rune in `arg`
2. If the rune is between `'0'` and `'9'` (inclusive), return `true`
3. After the loop, return `false`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [08-count-character](../08-count-character/README.md)
