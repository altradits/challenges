# Skills for 17-toupper

## What You'll Learn

**Previous:** [16-isempty](../16-isempty/skills.md) | **Next:** [18-tolower](../18-tolower/skills.md)

**Challenge:** Write a function `ToUpper(s string) string` that converts all lowercase letters to uppercase, leaving all other characters unchanged.

## Core Concept: ASCII Math for Case Conversion

### What Is ASCII?

Every character on a standard keyboard has a numeric code called its ASCII value. The letters form two neat ranges:

```
'A' = 65   'B' = 66   ...   'Z' = 90
'a' = 97   'b' = 98   ...   'z' = 122
```

The gap between any lowercase letter and its uppercase pair is exactly **32**:

```
'a' - 'A' = 97 - 65 = 32
'z' - 'Z' = 122 - 90 = 32
```

### The Conversion Formula

To convert a lowercase letter to uppercase, **subtract 32**:

```go
'a' - 32 = 65   which is 'A'
'z' - 32 = 90   which is 'Z'
```

### Detecting Lowercase Letters

You must only apply the formula to lowercase letters. Other characters (digits, spaces, uppercase letters) must pass through unchanged:

```go
if c >= 'a' && c <= 'z' {
    // c is a lowercase letter
}
```

### Building a New String

Because Go strings are immutable, you cannot change them in place. Build a new string by accumulating characters:

```go
func ToUpper(s string) string {
    result := ""
    for _, c := range s {
        if c >= 'a' && c <= 'z' {
            result += string(c - 32)   // convert to uppercase
        } else {
            result += string(c)        // keep as-is
        }
    }
    return result
}
```

Step by step:
1. Start with an empty `result`
2. For each rune `c` in `s`:
   - If `c` is a lowercase letter, subtract 32 and add to `result`
   - Otherwise add `c` unchanged
3. Return `result`

### Why This Works

In Go, runes are integers (`int32`). Arithmetic on runes is legal:

```go
var c rune = 'a'   // c = 97
c - 32             // = 65
string(c - 32)     // = "A"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Subtracting 32 from ALL characters | Corrupts digits, spaces, uppercase | Check `c >= 'a' && c <= 'z'` first |
| Adding 32 instead of subtracting | Converts to lowercase, not uppercase | Use `c - 32` for uppercase |
| Forgetting `string()` when appending | Type mismatch | `result += string(c - 32)` |
| Modifying `s` directly | Strings are immutable, won't compile | Build separate `result` string |

## Solving This Challenge

### Algorithm

1. `result := ""`
2. For each rune `c` in `s`:
   - If `'a' <= c <= 'z'`: append `string(c - 32)` to result
   - Else: append `string(c)` to result
3. Return `result`

### Trace Through an Example

Input: `"Hello!"`

| Char | In range 'a'-'z'? | Output char |
|------|------------------|-------------|
| `H` | No (uppercase already) | `H` |
| `e` | Yes | `E` |
| `l` | Yes | `L` |
| `l` | Yes | `L` |
| `o` | Yes | `O` |
| `!` | No (punctuation) | `!` |

Result: `"HELLO!"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [18-tolower](../18-tolower/README.md)
