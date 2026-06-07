# 47. String Prefix

## Instructions

Write a program that checks if a string starts with a given prefix.

### Function Signature

```go
func HasPrefix(s, prefix string) bool
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "hello"` | `true` |
| `"hello world", "world"` | `false` |
| `"", ""` | `true` |
| `"abc", "abcd"` | `false` |

### Constraints

- Return true if prefix is empty
- Return false if s is shorter than prefix
- Do not use `strings.HasPrefix`

### Hints

- Compare the first len(prefix) characters of s
