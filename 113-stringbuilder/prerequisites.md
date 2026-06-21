# Prerequisites for 113-stringbuilder

## Before You Start

### 1. Why Strings are Immutable in Go

In Go, you cannot modify a string in place. Every `+=` creates a brand new string:

```go
s := "hello"
s += " world"  // "hello" is discarded; a new string "hello world" is allocated
```

For one concatenation this is fine. For 1000 concatenations in a loop, you're doing 1+2+3+...+1000 ≈ 500,000 byte-copies. That's O(n²).

### 2. What `strings.Builder` Is

`strings.Builder` is a type from the `strings` package that holds an internal buffer. Writes append to the buffer without copying the entire accumulated string. At the end, one call to `.String()` converts the buffer to a string.

```go
import "strings"

var b strings.Builder
b.WriteString("Hello")
b.WriteString(", ")
b.WriteString("World")
fmt.Println(b.String())  // "Hello, World"
```

### 3. The Three Write Methods

```go
var b strings.Builder
b.WriteRune('H')        // add a rune (any Unicode character)
b.WriteByte('i')        // add a byte (ASCII only)
b.WriteString(" there") // add a whole string
fmt.Println(b.String()) // "Hi there"
```

Use `WriteRune` when ranging over a string (you get runes). Use `WriteString` when adding whole strings.

### 4. Checking for Vowels

One clean approach: check if the character is in a vowel string using `strings.ContainsRune`:

```go
func isVowel(c rune) bool {
    return strings.ContainsRune("aeiouAEIOU", c)
}
```

Or use a `switch` statement:

```go
switch c {
case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
    // it's a vowel
}
```

### 5. The O(n²) vs O(n) Difference

```go
// 10,000 character string:
// += approach: ~50,000,000 byte-copies (slow)
// Builder approach: ~10,000 byte-copies total (fast)
```

For strings of reasonable size (< 1000 chars) you won't notice. For large strings, Builder is dramatically faster.

## Review If Stuck

- [111-union](../111-union/skills.md) and [112-inter](../112-inter/skills.md) — both use `result += string(c)` which is fine for small inputs. This challenge teaches the proper approach for large inputs.

## You're Ready When You Can...

- [ ] Explain why `+=` in a loop is O(n²) for strings
- [ ] Declare a `strings.Builder` with `var b strings.Builder`
- [ ] Write a rune to it with `b.WriteRune(c)`
- [ ] Write a string to it with `b.WriteString(s)`
- [ ] Get the result with `b.String()`

## Next Steps

- [114-stringsplit](../114-stringsplit/README.md)
