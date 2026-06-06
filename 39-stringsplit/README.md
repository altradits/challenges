# 39-stringsplit

## Instructions

Write a program that splits a string into words and returns them as a slice.

### Function Signature

```go
func SplitWords(s string) []string
```

### Examples

| Input | Output |
|-------|--------|
| `"Hello World"` | `["Hello", "World"]` |
| `"  Go   is  fun  "` | `["Go", "is", "fun"]` |
| `""` | `[]` |
| `"single"` | `["single"]` |

### Constraints

- Handle multiple consecutive spaces correctly
- Trim leading and trailing spaces
- Return an empty slice for empty input
- Do not use `strings.Split` - implement the logic manually

### Hints

- Use `strings.Fields` or implement manual splitting
- Consider edge cases with whitespace
