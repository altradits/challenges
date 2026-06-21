# Prerequisites for 77-cameltosnakecase

## Before You Start

To solve this challenge you need to understand rune iteration, character classification with the `unicode` package, string building, and validation logic.

### 1. `unicode.IsUpper`, `unicode.IsLower`, `unicode.IsLetter`
These functions check the category of a rune:

```go
import "unicode"

unicode.IsUpper('A')    // true
unicode.IsUpper('a')    // false
unicode.IsLower('a')    // true
unicode.IsLetter('a')   // true
unicode.IsLetter('1')   // false  (digits are not letters)
unicode.IsLetter('!')   // false
```

### 2. Accessing the index in `for...range`
When you need both the position and the character:

```go
for i, r := range s {
    if i == 0 {
        // first character
    }
    // r is the current rune
}
```

### 3. Writing `_` to a `strings.Builder`

```go
import "strings"

var b strings.Builder
b.WriteRune('_')       // write underscore
b.WriteRune(r)         // write the current rune
```

### 4. Validation before transformation
Check all the rules up front. If the input is invalid, return it unchanged:

```go
func CamelToSnakeCase(s string) string {
    if !isValid(s) {
        return s
    }
    // transform...
}
```

### 5. Checking previous character (lookback)
When processing character at index `i`, you can check the character at `i-1` by indexing a pre-built rune slice:

```go
runes := []rune(s)
for i, r := range runes {
    if i > 0 && unicode.IsUpper(r) && unicode.IsUpper(runes[i-1]) {
        // two consecutive capitals
    }
}
```

## Review If Stuck

- [72-removespaces](../72-removespaces/skills.md) — covers `strings.Builder` and filtering characters
- [75-titlecase](../75-titlecase/skills.md) — covers `unicode.ToUpper`/`ToLower` and character-by-character processing
- [66-checknumber](../66-checknumber/skills.md) — covers `for...range` and early return for validation

## You're Ready When You Can...

- [ ] Use `unicode.IsUpper(r)` and `unicode.IsLetter(r)` to classify a rune
- [ ] Access both the index `i` and character `r` in a `for...range` loop
- [ ] Write a validation function that returns `false` on the first rule violation
- [ ] Use `strings.Builder` to append `_` and a rune conditionally
- [ ] Convert a string to `[]rune` so you can look at `runes[i-1]`

## Next Steps

After completing this exercise, move to:
- [78-countrepeats](../78-countrepeats/README.md)
