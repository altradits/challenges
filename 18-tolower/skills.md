# Skills for 18-tolower

## What You'll Learn

**Previous:** [17-toupper](../17-toupper/skills.md) | **Next:** [19-countalpha](../19-countalpha/skills.md)

**Challenge:** Write a function `ToLower(s string) string` that converts all uppercase letters to lowercase, leaving all other characters unchanged.

## Core Concept: The Reverse of Case Conversion

### The ToLower Formula

In [17-toupper](../17-toupper/skills.md) you learned that uppercase and lowercase letters are 32 apart in ASCII. `ToUpper` subtracts 32. `ToLower` is the exact mirror — **add 32**:

```
'A' + 32 = 65 + 32 = 97 = 'a'
'Z' + 32 = 90 + 32 = 122 = 'z'
```

### Detecting Uppercase Letters

Only apply the formula to uppercase letters:

```go
if c >= 'A' && c <= 'Z' {
    // c is an uppercase letter — convert it
}
```

### Full Implementation

```go
func ToLower(s string) string {
    result := ""
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            result += string(c + 32)   // convert to lowercase
        } else {
            result += string(c)        // keep as-is
        }
    }
    return result
}
```

### Comparing ToUpper and ToLower Side by Side

| Operation | Check | Formula | Example |
|-----------|-------|---------|---------|
| `ToUpper` | `c >= 'a' && c <= 'z'` | `c - 32` | `'a'` → `'A'` |
| `ToLower` | `c >= 'A' && c <= 'Z'` | `c + 32` | `'A'` → `'a'` |

The structure of the two functions is identical — only the range check and the direction of the arithmetic change.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `c - 32` | Converts to uppercase, not lowercase | Use `c + 32` |
| Checking `'a'` to `'z'` range | Checks lowercase, should check uppercase | Use `'A'` to `'Z'` |
| Converting digits and symbols | Corrupts non-letter characters | Only convert when in `'A'`–`'Z'` range |

## Solving This Challenge

### Algorithm

1. `result := ""`
2. For each rune `c` in `s`:
   - If `'A' <= c <= 'Z'`: append `string(c + 32)` to result
   - Else: append `string(c)` to result
3. Return `result`

### Trace Through an Example

Input: `"Go is FUN!"`

| Char | In 'A'-'Z'? | Output |
|------|------------|--------|
| `G` | Yes | `g` |
| `o` | No | `o` |
| ` ` | No | ` ` |
| `i` | No | `i` |
| `s` | No | `s` |
| ` ` | No | ` ` |
| `F` | Yes | `f` |
| `U` | Yes | `u` |
| `N` | Yes | `n` |
| `!` | No | `!` |

Result: `"go is fun!"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [19-countalpha](../19-countalpha/README.md)
