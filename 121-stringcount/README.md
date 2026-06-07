# 46. String Count

## Instructions

Write a program that counts occurrences of a substring in a string.

### Function Signature

```go
func Count(s, substr string) int
```

### Examples

| Input | Output |
|-------|--------|
| `"hello hello", "l"` | `4` |
| `"hello hello", "ll"` | `2` |
| `"hello hello", "xyz"` | `0` |
| `"", "a"` | `0` |

### Constraints

- Return 0 if substr is empty
- Count overlapping occurrences
- Do not use `strings.Count`

### Hints

- Use `strings.Index` in a loop
- Move forward by 1 after each match to count overlaps
