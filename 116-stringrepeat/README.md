# 41. String Repeat

## Instructions

Write a program that repeats a string n times.

### Function Signature

```go
func Repeat(s string, n int) string
```

### Examples

| Input | Output |
|-------|--------|
| `"abc", 3` | `"abcabcabc"` |
| `"x", 5` | `"xxxxx"` |
| `"go", 0` | `""` |
| `"", 10` | `""` |

### Constraints

- Return empty string if n <= 0
- Do not use `strings.Repeat`
- Use `strings.Builder` for efficiency

### Hints

- Loop n times and append the string each iteration
- Handle negative n values
