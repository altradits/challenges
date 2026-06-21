# 49. String Field

## Instructions

Write a program that splits a string into fields using a delimiter.

### Function Signature

```go
func Split(s, sep string) []string
```

### Examples

| Input | Output |
|-------|--------|
| `"a,b,c", ","` | `["a", "b", "c"]` |
| `"a:b:c", ":"` | `["a", "b", "c"]` |
| `"abc", "x"` | `["abc"]` |
| `"", ","` | `[""]` |

### Constraints

- Return slice with one empty string if input is empty
- Do not use `strings.Split`
- Handle consecutive delimiters (empty fields)

### Hints

- Use `strings.Index` to find each delimiter
- Build result slice incrementally
