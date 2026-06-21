# 19. Replace Character

## What You'll Learn

This exercise teaches you **character substitution and string building**. By the end, you'll understand:
- How to replace specific characters in a string
- Building a new string with substitutions
- The difference between replace and replaceAll
- Handling cases where the character doesn't exist

## Theory: Character Replacement

### What is Replacement?

Replacement means **substituting one character for another**:

```
"Hello" with 'l' → 'L'
Result: "HeLLo"
```

### Replace vs ReplaceAll

| Operation | Description | Example |
|-----------|-------------|---------|
| Replace first | Only first occurrence | `"Hello"` → `"HeLlo"` |
| Replace all | Every occurrence | `"Hello"` → `"HeLLo"` |

**This exercise replaces ALL occurrences.**

### The Build-New-String Pattern

```go
func replaceAll(s string, old, new rune) string {
    var result string
    for _, c := range s {
        if c == old {
            result += string(new)  // Substitute
        } else {
            result += string(c)    // Keep original
        }
    }
    return result
}
```

## Step-by-Step Approach

1. **Initialize** empty result string
2. **Iterate** through each character
3. **If match**: append replacement character
4. **Else**: append original character
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Modifying original | Strings are immutable | Build new string |
| Only replacing first | Should replace all | Check every character |
| Wrong type conversion | rune vs string | Use `string(c)` |

## Practice Tips

- Test with: `"Hello"`, replace `'l'` with `'L'` → `"HeLLo"`
- Test with: `"Hello"`, replace `'z'` with `'X'` → `"Hello"` (no change)
- Test with: `""`, any replace → `""`

## The Challenge

Write a function that replaces all occurrences of a character with another character.

### Expected Function

```go
func ReplaceChar(s string, old, new rune) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello", 'l', 'L')` | `"HeLLo"` | Both 'l's replaced |
| `("Hello", 'z', 'X')` | `"Hello"` | No 'z' to replace |
| `("", 'a', 'b')` | `""` | Empty string |
| `("aaaa", 'a', 'b')` | `"bbbb"` | All replaced |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(ReplaceChar("Hello", 'l', 'L'))  // HeLLo
    fmt.Println(ReplaceChar("Hello", 'z', 'X'))  // Hello
    fmt.Println(ReplaceChar("", 'a', 'b'))        // (empty)
    fmt.Println(ReplaceChar("aaaa", 'a', 'b'))    // bbbb
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What's the difference between replace and replaceAll?
2. Why do we build a new string?
3. How do you convert a rune to a string?

## Next Steps

After completing this, you'll be ready for:
- [32-printif](../32-printif/README.md) - Printif
- [33-printifnot](../33-printifnot/README.md) - Printifnot