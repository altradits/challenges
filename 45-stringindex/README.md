# 45-stringindex

## Instructions

Write a program that finds the index of the first occurrence of a substring.

### Function Signature

```go
func Index(s, substr string) int
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "world"` | `6` |
| `"hello world", "xyz"` | `-1` |
| `"abcabc", "b"` | `1` |
| `"", "a"` | `-1` |

### Constraints

- Return -1 if substring not found
- Return 0 if substring is empty
- Do not use `strings.Index`

### Hints

- Iterate through the string checking for matches
- Use rune-based iteration for Unicode safety
