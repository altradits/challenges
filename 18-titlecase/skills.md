# Skills for 18-titlecase

## What You'll Learn

**Previous:** [17-reversestring](../17-reversestring/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Titlecase

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String transformation and case conversion

Go's `unicode` package provides case conversion functions:
- `unicode.ToUpper(r)` - convert rune to uppercase
- `unicode.ToLower(r)` - convert rune to lowercase
- `unicode.IsUpper(r)` / `unicode.IsLower(r)` - check case

You can also use ASCII math: uppercase and lowercase letters differ by 32.

```go
// ASCII conversion
if c >= 'a' && c <= 'z' {
    c = c - 32  // to uppercase
}
```

### 3. String splitting and joining

Go's `strings` package provides split and join functions:
- `strings.Split(s, sep)` - split string into slice
- `strings.Join(slice, sep)` - join slice into string
- Manual implementation helps understand the logic

```go
// Manual split example
parts := []string{}
current := ""
for _, c := range s {
    if c == sep {
        parts = append(parts, current)
        current = ""
    } else {
        current += string(c)
    }
}
```

### 4. String filtering and cleaning

Filtering strings involves:
- Iterating through characters
- Checking conditions (is space? is digit? etc.)
- Building a new string with only wanted characters

```go
var result strings.Builder
for _, c := range s {
    if condition(c) {
        result.WriteRune(c)
    }
}
```

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [19-wordcount](../19-wordcount/skills.md) - Wordcount
