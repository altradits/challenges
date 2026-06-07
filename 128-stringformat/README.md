# 53. String Format

## Instructions

Write a program that formats a string using positional arguments, similar to `fmt.Sprintf`.

### Function Signature

```go
func Format(template string, args ...interface{}) string
```

### Examples

| Input | Output |
|-------|--------|
| `"Hello %s", "World"` | `"Hello World"` |
| `"%d + %d = %d", 1, 2, 3` | `"1 + 2 = 3"` |
| `"%s is %d years old", "Alice", 30` | `"Alice is 30 years old"` |
| `"No placeholders"` | `"No placeholders"` |

### Constraints

- Support `%s` (string), `%d` (integer), `%f` (float)
- Return original template if no placeholders
- Handle mismatched argument counts gracefully

### Hints

- Use `fmt.Sprintf` internally or implement manual replacement
- Count placeholders and match with args
