# 38-stringbuilder

## Instructions

Write a program that takes a string and returns a new string with all vowels removed.

Vowels are: a, e, i, o, u (both lowercase and uppercase).

### Function Signature

```go
func RemoveVowels(s string) string
```

### Examples

| Input | Output |
|-------|--------|
| `"Hello World"` | `"Hll Wrld"` |
| `"AEIOUaeiou"` | `""` |
| `"Go Programming"` | `"G Prgrmmng"` |
| `""` | `""` |

### Constraints

- Do not use the `strings.ReplaceAll` function
- Build the result using a `strings.Builder` for efficiency
- Preserve the original case of consonants
- Handle empty strings gracefully

### Hints

- Use `strings.Builder` to efficiently build the result string
- Iterate through each rune in the input string
- Check if each rune is a vowel before appending
