# Prerequisites for 42-cameltosnakecase

## Before You Start

### 1. `unicode.IsUpper` and `unicode.IsLower`

These functions from the `unicode` package check the case of a letter rune:

```go
import "unicode"

fmt.Println(unicode.IsUpper('A'))  // true
fmt.Println(unicode.IsUpper('a'))  // false
fmt.Println(unicode.IsUpper('!'))  // false

fmt.Println(unicode.IsLower('z'))  // true
fmt.Println(unicode.IsLower('Z'))  // false
```

### 2. `unicode.IsLetter`

Checks whether a rune is any kind of letter (upper or lower case, any alphabet):

```go
fmt.Println(unicode.IsLetter('a'))  // true
fmt.Println(unicode.IsLetter('Z'))  // true
fmt.Println(unicode.IsLetter('5'))  // false
fmt.Println(unicode.IsLetter('!'))  // false
```

### 3. Iterating with Index

When you need to know the position of the character (to skip the first one), use `for i, c := range s`:

```go
s := "HelloWorld"
for i, c := range s {
    if unicode.IsUpper(c) && i > 0 {
        fmt.Printf("Uppercase at position %d: %c\n", i, c)
    }
}
// Uppercase at position 5: W
```

### 4. Converting a Rune to String

`string(c)` converts a single rune to a one-character string:

```go
var c rune = 'H'
fmt.Println(string(c))  // "H"
fmt.Println("_" + string(c))  // "_H"
```

### 5. Validate-Then-Transform Pattern

A common Go pattern: check whether input is valid, return unchanged if not, otherwise proceed:

```go
func process(s string) string {
    if !isValid(s) {
        return s  // invalid input: return as-is
    }
    // transform and return
}
```

## Review If Stuck

- [37-countwords](../37-countwords/skills.md) — covers `unicode.IsLetter`, `unicode.IsDigit`, and iteration with index
- [41-join](../41-join/skills.md) — covers building result strings character by character

## You're Ready When You Can...

- [ ] Use `unicode.IsUpper(c)` to check if a rune is uppercase
- [ ] Use `unicode.IsLetter(c)` to validate a character is a letter
- [ ] Iterate a string with `for i, c := range s` and use `i` to skip the first character
- [ ] Explain the two-phase validate-then-transform approach

## Next Steps

- [43-itoa](../43-itoa/README.md)
