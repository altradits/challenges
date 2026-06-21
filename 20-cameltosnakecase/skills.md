# Skills for 20-cameltosnakecase

## What You'll Learn

**Previous:** [19-wordcount](../19-wordcount/skills.md) | **Next:** [21-countrepeats](../21-countrepeats/README.md)

**Challenge:** Write `CamelToSnakeCase(s string) string` that converts camelCase to snake_case — but only if the string is valid camelCase. Invalid input is returned unchanged.

## Core Concept: `unicode.IsUpper` and Building a New String with Inserted Characters

### What Is It?

This challenge combines everything you have learned so far: iterating with `for...range`, checking character properties with the `unicode` package, building a new string, and applying logic based on position in the string.

The new tools are `unicode.IsUpper(r)`, `unicode.IsLower(r)`, `unicode.IsLetter(r)`, and `unicode.ToLower(r)`.

### What Valid CamelCase Means (Per the README)

A string is valid camelCase only if:
1. It does NOT end on a capital letter
2. No two capital letters appear in a row
3. No digits or punctuation anywhere

If any of these rules are violated, return the string unchanged.

### Validation First

Before transforming, validate the input:

```go
import "unicode"

func isValidCamelCase(s string) bool {
    if len(s) == 0 {
        return false
    }
    runes := []rune(s)
    last := len(runes) - 1

    // Rule: must not end on a capital letter
    if unicode.IsUpper(runes[last]) {
        return false
    }

    for i, r := range runes {
        // Rule: only letters allowed
        if !unicode.IsLetter(r) {
            return false
        }
        // Rule: no two consecutive capitals
        if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(runes[i-1]) {
            return false
        }
    }
    return true
}
```

### Transformation: Insert `_` Before Each Capital

For valid camelCase, walk the runes. When you see an uppercase letter (and it is not the very first character), insert `_` then the lowercase version of that letter.

```go
func CamelToSnakeCase(s string) string {
    if !isValidCamelCase(s) {
        return s
    }

    var b strings.Builder
    for i, r := range s {
        if unicode.IsUpper(r) && i > 0 {
            b.WriteRune('_')
            b.WriteRune(unicode.ToLower(r))
        } else {
            b.WriteRune(r)
        }
    }
    return b.String()
}
```

### Diagram: Converting "camelToSnakeCase"

```
c  a  m  e  l  T  o  S  n  a  k  e  C  a  s  e
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
c  a  m  e  l  _t o  _s n  a  k  e  _c a  s  e
(T -> _t)     (S -> _s)              (C -> _c)

Result: "camel_To_Snake_Case"
```

Wait — the expected output for `"camelToSnakeCase"` is `"camel_To_Snake_Case"`. The capital letter becomes `_` + lowercase. Check: `T` -> `_t`? No — `T` -> `_T`... Let's re-read the README output:

```
CamelToSnakeCase("camelToSnakeCase")   -> "camel_To_Snake_Case"
```

The output keeps the letter as capital: `_To`, `_Snake`, `_Case`. So the transformation is: insert `_` before the capital, keep it as-is (do not lowercase it).

```go
if unicode.IsUpper(r) && i > 0 {
    b.WriteRune('_')
    b.WriteRune(r)   // keep capital as-is
} else {
    b.WriteRune(r)
}
```

### Checking the Example Outputs

```
"HelloWorld"  ->  "Hello_World"   (H stays, capital W gets _W)
"helloWorld"  ->  "hello_World"   (capital W gets _W)
"camelCase"   ->  "camel_Case"    (capital C gets _C)
"CAMELtoSnackCASE" -> "CAMELtoSnackCASE" (two consecutive caps: invalid, return as-is)
"camelToSnakeCase" -> "camel_To_Snake_Case"
"hey2"        -> "hey2"           (digit: invalid, return as-is)
```

For `"HelloWorld"`, `H` is at index 0 so no underscore; `W` is at index 5, insert `_` before it.

### `unicode.IsUpper`, `unicode.IsLower`, `unicode.IsLetter`

```go
unicode.IsUpper('A')   // true
unicode.IsUpper('a')   // false
unicode.IsUpper('1')   // false
unicode.IsLower('a')   // true
unicode.IsLetter('a')  // true
unicode.IsLetter('1')  // false
unicode.IsLetter('!')  // false
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not checking for consecutive capitals | `"CAMELCase"` gets transformed instead of returned unchanged | Check `IsUpper(runes[i]) && IsUpper(runes[i-1])` |
| Not checking for digits/punctuation | `"hey2"` gets transformed | Check `!unicode.IsLetter(r)` for each character |
| Lowercasing the letter after `_` | `"Hello_World"` becomes `"Hello_world"` | Keep the original rune: `b.WriteRune(r)` not `b.WriteRune(unicode.ToLower(r))` |
| Inserting `_` before the first character | `"HelloWorld"` becomes `"_Hello_World"` | Add `&& i > 0` to the condition |

## Solving This Challenge

### Algorithm

1. If `s` is empty, return `""`
2. Validate: check for consecutive capitals, non-letter characters, and capital at end — return `s` if invalid
3. Build result: for each rune, if it is uppercase and not the first character, write `_` then the rune; otherwise write the rune as-is
4. Return the built string

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [21-countrepeats](../21-countrepeats/README.md)
