# Skills for 163-type-conversions

## What You'll Learn

**Previous:** [162-variadic-functions](../162-variadic-functions/skills.md) | **Next:** [164-constants-iota](../164-constants-iota/skills.md)

**Challenge:** Convert between numeric types, strings, bytes, and runes. Understand UTF-8.

## Type Conversions in Go

Go has **no implicit type conversion**. Every conversion is explicit:

```go
var n int = 42
var f float64 = float64(n)   // explicit conversion
var m int = int(f)           // explicit truncation (not rounding)

// This would NOT compile:
// var f float64 = n   // cannot use int as float64
```

The syntax is `targetType(value)` — looks like a function call.

### Numeric Conversions

```go
int(3.9)     // 3   — truncates toward zero (floor for positives, ceil for negatives)
int(-3.9)    // -3  — NOT -4
float64(7)   // 7.0
int32(1000)  // 1000
int8(300)    // 44  — wraps around! int8 max is 127
uint(−1)     // huge number — sign bit becomes part of the value
```

**Narrowing conversions silently wrap** — Go does not panic. Always check the target range.

### String ↔ []byte ↔ []rune

These three types are interchangeable via conversions:

```go
s := "hello"
b := []byte(s)    // []byte{104, 101, 108, 108, 111}
s2 := string(b)   // "hello"

r := []rune(s)    // []rune{104, 101, 108, 108, 111}
s3 := string(r)   // "hello"
```

For ASCII, `[]byte` and `[]rune` are identical. They differ for multi-byte characters.

## Strings Are UTF-8 Byte Sequences

Go source code is always UTF-8. A Go string is a sequence of **bytes** — not characters. A single Unicode character can be 1, 2, 3, or 4 bytes.

```go
s := "日本語"
len(s)             // 9 — bytes!
len([]rune(s))     // 3 — Unicode characters
```

## Runes — Unicode Code Points

A **rune** is Go's name for a Unicode code point. Its type is `int32`:

```go
var r rune = '€'
fmt.Println(r)         // 8364 — the Unicode code point for €
fmt.Println(string(r)) // €

// Single-quoted literals are rune literals
'A'   // rune(65)
'€'   // rune(8364)
'日'  // rune(26085)
```

### Iterating Over Strings

```go
// Iterating by byte — can split multi-byte characters!
for i := 0; i < len(s); i++ {
    b := s[i]   // b is a byte (uint8)
}

// Iterating by rune — correct for Unicode
for i, r := range s {
    _ = i  // byte index of this rune's start
    _ = r  // the rune itself
}
```

Always range over strings when you need characters, not bytes.

### len vs utf8.RuneCountInString

```go
import "unicode/utf8"

s := "日本語"
len(s)                           // 9 (bytes)
utf8.RuneCountInString(s)        // 3 (runes) — same as len([]rune(s))
```

## String ↔ int (strconv)

```go
import "strconv"

n, err := strconv.Atoi("42")        // string → int
s := strconv.Itoa(42)               // int → string

f, err := strconv.ParseFloat("3.14", 64)  // string → float64
s := strconv.FormatFloat(3.14, 'f', 2, 64) // float64 → string
```

## Solving the Challenge

```go
package piscine

import "unicode/utf8"

func IntToFloat(n int) float64 {
    return float64(n)
}

func FloatToInt(f float64) int {
    return int(f)
}

func RuneToString(r rune) string {
    return string(r)
}

func StringToRunes(s string) []rune {
    return []rune(s)
}

func BytesToString(b []byte) string {
    return string(b)
}

func CountRunes(s string) int {
    return utf8.RuneCountInString(s)
}
```

## Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `len("日本語")` for character count | Returns bytes, not characters | Use `utf8.RuneCountInString` or `len([]rune(s))` |
| `string(65)` — thinking it gives "65" | Gives "A" (the character with code point 65) | Use `strconv.Itoa(65)` for "65" |
| Narrowing int without range check | Silent wrap-around | Check value fits in target type first |
| `s[i]` to get characters | Gives a byte, breaks on multi-byte chars | Use `range s` to get runes |

## Documentation

- [Go Blog — Strings, bytes, runes, and characters in Go](https://go.dev/blog/strings)
- [unicode/utf8 package](https://pkg.go.dev/unicode/utf8)
- [strconv package](https://pkg.go.dev/strconv)
- [A Tour of Go — Strings and rune literals](https://go.dev/tour/basics/11)

**Next:** [164-constants-iota](../164-constants-iota/README.md)
