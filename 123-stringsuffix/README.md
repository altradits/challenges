# 48-stringsuffix

## Instructions

Write a program that checks if a string ends with a given suffix.

### Function Signature

```go
func HasSuffix(s, suffix string) bool
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "world"` | `true` |
| `"hello world", "hello"` | `false` |
| `"", ""` | `true` |
| `"abc", "abcd"` | `false` |

### Constraints

- Return true if suffix is empty
- Return false if s is shorter than suffix
- Do not use `strings.HasSuffix`

### Hints

- Compare the last len(suffix) characters of s
