# Skills for 73-rpncalc

## What You'll Learn

**Previous:** [72-brackets](../72-brackets/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Rpncalc

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String searching and indexing

Go provides several ways to search within strings:
- `strings.Index()` - find first occurrence
- `strings.LastIndex()` - find last occurrence
- Manual iteration with `for...range` for custom search logic
- Compare runes or bytes directly

```go
// Manual search example
for i, c := range s {
    if c == target {
        return i
    }
}
return -1
```

### 3. String transformation and case conversion

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

**Next:** [74-brainfuck](../74-brainfuck/skills.md) - Brainfuck
