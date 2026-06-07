# Skills for 15-removespaces

## What You'll Learn

**Previous:** [13-printifnot](../13-printifnot/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Removespaces

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

### 3. String filtering and cleaning

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

**Next:** [16-retainfirsthalf](../16-retainfirsthalf/skills.md) - Retainfirsthalf
