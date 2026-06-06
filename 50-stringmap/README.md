# 50-stringmap

## Instructions

Write a program that applies a function to each character in a string.

### Function Signature

```go
func Map(s string, f func(rune) rune) string
```

### Examples

| Input | Output |
|-------|--------|
| `"abc", ToUpper` | `"ABC"` |
| `"Hello", ToLower` | `"hello"` |
| `"123", func(r rune) rune { return r + 1 }` | `"234"` |
| `""` | `""` |

### Constraints

- Apply the function to each rune
- Return the transformed string
- Handle empty strings

### Hints

- Use `strings.Builder` to build the result
- Iterate with `for...range` to handle runes correctly
