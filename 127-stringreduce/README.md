# 52-stringreduce

## Instructions

Write a program that reduces a string to a single value by applying a function cumulatively.

### Function Signature

```go
func Reduce(s string, f func(rune, rune) rune) rune
```

### Examples

| Input | Output |
|-------|--------|
| `"abc", Max` | `'c'` |
| `"hello", Min` | `'e'` |
| `"123", func(a, b rune) rune { return a + b }` | `'f'` (49+50+51) |
| `""` | `0` |

### Constraints

- Return 0 for empty string
- Apply function left-to-right
- Handle single character strings

### Hints

- Initialize with first rune, then apply to rest
- Use `for...range` to iterate runes
