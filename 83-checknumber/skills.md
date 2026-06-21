# Skills for 83-checknumber

## What You'll Learn

**Previous:** [82-countalpha](../82-countalpha/skills.md) | **Next:** [84-countvowels](../84-countvowels/skills.md)

**Challenge:** Write a function `CheckNumber(s string) bool` that returns `true` if the string contains at least one digit, `false` otherwise.

## Core Concept: Early Return — Stop as Soon as You Find a Match

### What Is "Early Return"?

In the previous challenges you always scanned the entire string. For a question like "does this string contain ANY digit?", you can stop the moment you find the first digit:

```go
func CheckNumber(s string) bool {
    for _, c := range s {
        if c >= '0' && c <= '9' {
            return true   // found one — stop immediately
        }
    }
    return false   // checked everything, no digit found
}
```

This is the **early return** pattern. It is more efficient than counting: once the answer is definitely "yes", there is no reason to keep looking.

### The Digit Range

Digits `'0'` through `'9'` have contiguous ASCII values:

```
'0' = 48
'1' = 49
...
'9' = 57
```

Check for any digit with a range comparison:

```go
c >= '0' && c <= '9'
```

### Counting vs Detecting

| Goal | Strategy | Returns |
|------|----------|---------|
| How many digits? | Count every match, scan all | `int` |
| Does any digit exist? | Return on first match | `bool` |

`CheckNumber` is a **detection** question, so early return is the right tool.

### Contrast with CountAlpha

In [82-countalpha](../82-countalpha/skills.md), you needed to scan the whole string to count. Here, as soon as you find one digit the answer is `true` and you are done.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Returning a count instead of `bool` | Wrong return type | Return `true`/`false` |
| Forgetting `return false` after the loop | Compile error: missing return | Always add the fallback `return false` |
| Checking individual digits: `c == '0' \|\| c == '1' \|\| ...` | Tedious and error-prone | Use `c >= '0' && c <= '9'` |

## Solving This Challenge

### Algorithm

1. For each rune `c` in `s`:
   - If `c >= '0' && c <= '9'`: return `true`
2. Return `false`

### Trace Through an Example

Input: `"Hello1"`

| Char | Digit? | Action |
|------|--------|--------|
| `H` | No | continue |
| `e` | No | continue |
| `l` | No | continue |
| `l` | No | continue |
| `o` | No | continue |
| `1` | Yes | return `true` immediately |

Input: `"Hello"`

All characters checked, none are digits — return `false`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [84-countvowels](../84-countvowels/skills.md)
