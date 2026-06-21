# Skills for 19-countalpha

## What You'll Learn

**Previous:** [18-tolower](../18-tolower/skills.md) | **Next:** [20-checknumber](../20-checknumber/skills.md)

**Challenge:** Write a function `CountAlpha(s string) int` that returns the number of alphabetic characters (a–z, A–Z) in the string.

## Core Concept: Counting Characters That Match a Condition

### What Is an Alphabetic Character?

An alphabetic character is any letter — uppercase or lowercase:

```
a b c ... z   (ASCII 97–122)
A B C ... Z   (ASCII 65–90)
```

Digits, spaces, and punctuation are NOT alphabetic.

### The Counting Pattern

This is the standard accumulator pattern applied to a filter condition:

```go
func CountAlpha(s string) int {
    count := 0
    for _, c := range s {
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
            count++
        }
    }
    return count
}
```

Step by step:
1. Start with `count := 0`
2. For each rune `c` in `s`:
   - If `c` is a lowercase letter OR an uppercase letter, increment `count`
3. Return `count`

### The `||` (OR) Operator

A character is alphabetic if it is in the lowercase range **OR** in the uppercase range. In Go the OR operator is `||`:

```go
(c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
```

Break it down:
- `c >= 'a' && c <= 'z'` — true if `c` is a lowercase letter
- `c >= 'A' && c <= 'Z'` — true if `c` is an uppercase letter
- `||` — the whole expression is true if either half is true

### Difference from Previous Challenges

| Challenge | Action per character | Returns |
|-----------|---------------------|---------|
| 17-toupper | Transform character | `string` |
| 16-isempty | Stop at first match | `bool` |
| 19-countalpha | Count matching characters | `int` |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Only checking `'a'`–`'z'` | Misses uppercase letters | Add `|| (c >= 'A' && c <= 'Z')` |
| Counting digits as alpha | Digits are not letters | Only include the letter ranges |
| Using `&` instead of `&&` | Bitwise AND, not logical AND | Use `&&` for conditions |

## Solving This Challenge

### Algorithm

1. `count := 0`
2. For each rune `c` in `s`:
   - If `(c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')`: `count++`
3. Return `count`

### Trace Through an Example

Input: `"H1e2l3"`

| Char | Alpha? | count |
|------|--------|-------|
| `H` | Yes | 1 |
| `1` | No | 1 |
| `e` | Yes | 2 |
| `2` | No | 2 |
| `l` | Yes | 3 |
| `3` | No | 3 |

Result: `3`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [20-checknumber](../20-checknumber/README.md)
