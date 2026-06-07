# 44. String Contains

## Instructions

Write a program that checks if a string contains a substring.

### Function Signature

```go
func Contains(s, substr string) bool
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "world"` | `true` |
| `"hello world", "xyz"` | `false` |
| `"", "a"` | `false` |
| `"abc", ""` | `true` |

### Constraints

- Return true if substr is empty
- Return false if s is empty and substr is not
- Do not use `strings.Contains`

### Hints

- Use `strings.Index` to check for substring presence
