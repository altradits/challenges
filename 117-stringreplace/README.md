# 42. String Replace

## Instructions

Write a program that replaces all occurrences of a substring with another substring.

### Function Signature

```go
func Replace(s, old, new string) string
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "l", "L"` | `"heLLo worLd"` |
| `"foo bar foo", "foo", "baz"` | `"baz bar baz"` |
| `"abc", "x", "y"` | `"abc"` |
| `"", "a", "b"` | `""` |

### Constraints

- Replace all occurrences, not just the first
- Handle empty strings gracefully
- Do not use `strings.ReplaceAll`

### Hints

- Use `strings.Index` to find each occurrence
- Build the result incrementally
