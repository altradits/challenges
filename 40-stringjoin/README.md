# 40-stringjoin

## Instructions

Write a program that joins a slice of strings with a separator.

### Function Signature

```go
func Join(words []string, sep string) string
```

### Examples

| Input | Output |
|-------|--------|
| `["a","b","c"], "-"` | `"a-b-c"` |
| `["hello","world"], " "` | `"hello world"` |
| `[], ","` | `""` |
| `["single"], "|"` | `"single"` |

### Constraints

- Return empty string for empty slice
- Do not use `strings.Join` - implement manually
- Handle single-element slices correctly

### Hints

- Use `strings.Builder` for efficient concatenation
- Add separator between elements, not after the last one
