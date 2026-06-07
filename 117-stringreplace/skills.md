# Skills for 117-stringreplace

## What You'll Learn

**Previous:** [116-stringrepeat](../116-stringrepeat/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Stringreplace

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String building and concatenation

In Go, strings are immutable, so building strings character by character requires care. You can:
- Use `+` for simple concatenation
- Use `strings.Builder` for efficient string building in loops
- Convert runes to strings with `string(rune)`

```go
// Simple concatenation
result := "Hello" + " " + "World"

// Using strings.Builder for efficiency
var b strings.Builder
for _, c := range input {
    b.WriteRune(c)
}
result := b.String()
```

### 3. String searching and indexing

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

### 4. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [118-stringtrim](../118-stringtrim/skills.md) - Stringtrim
