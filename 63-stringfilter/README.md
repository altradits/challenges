# 51. String Filter

## Instructions

Write a program that filters characters from a string based on a predicate function.

### Function Signature

```go
func Filter(s string, f func(rune) bool) string
```

### Examples

| Input | Output |
|-------|--------|
| `"hello", IsVowel` | `"hll"` |
| `"abc123", IsDigit` | `"123"` |
| `"Go Lang", IsUpper` | `"GL"` |
| `""` | `""` |

### Constraints

- Keep characters where f(r) returns true
- Return empty string if no characters match
- Handle empty input

### Hints

- Use `strings.Builder` to build result
- Apply predicate to each rune
