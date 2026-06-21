# Prerequisites for 87-removespaces

## Before You Start

To solve this challenge you need to understand how to iterate over a string, filter characters, and build a new string from selected characters.

### 1. `for...range` on a string gives you runes
Iterating a string produces one rune per character:

```go
for _, r := range "hello" {
    // r is each character as a rune
}
```

### 2. Rune comparison with `!=`
To check that a rune is NOT a space:

```go
if r != ' ' {
    // r is not a space
}
```

`' '` (single quotes, a space character) is a rune literal.

### 3. `strings.Builder` for efficient string construction
Use `strings.Builder` to accumulate characters without creating many intermediate strings.

```go
import "strings"

var b strings.Builder
b.WriteRune('H')
b.WriteRune('i')
result := b.String()   // "Hi"
```

### 4. `strings.ReplaceAll` as a shortcut
The standard library can do this in one call:

```go
import "strings"

strings.ReplaceAll("hello world", " ", "")  // "helloworld"
```

### 5. Go strings are immutable
You cannot modify a string in place. You must build a new one.

```go
s := "hello"
s[0] = 'H'   // compile error: cannot assign to s[0]
```

## Review If Stuck

- [08-stringspackage](../08-stringspackage/skills.md) — covers `strings.ReplaceAll`, `strings.Builder`, and related functions
- [81-checknumber](../81-checknumber/skills.md) — covers `for...range` and rune comparisons
- [82-count-character](../82-count-character/skills.md) — covers iterating and testing individual runes

## You're Ready When You Can...

- [ ] Write a `for...range` loop over a string
- [ ] Test a rune against `' '` (space) using `!=`
- [ ] Use `strings.Builder` to collect characters, then call `.String()`
- [ ] Alternatively, use `strings.ReplaceAll` to remove all occurrences of a character

## Next Steps

- [88-retainfirsthalf](../88-retainfirsthalf/README.md)
