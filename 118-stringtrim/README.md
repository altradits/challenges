# 43. String Trim

## Instructions

Write a program that trims leading and trailing whitespace from a string.

### Function Signature

```go
func Trim(s string) string
```

### Examples

| Input | Output |
|-------|--------|
| `"  hello  "` | `"hello"` |
| `"\t\nGo\n\t"` | `"Go"` |
| `"no spaces"` | `"no spaces"` |
| `"   "` | `""` |

### Constraints

- Remove spaces, tabs, and newlines
- Do not use `strings.TrimSpace`
- Return empty string if only whitespace

### Hints

- Use `unicode.IsSpace` to check for whitespace characters
- Find the first and last non-space indices
