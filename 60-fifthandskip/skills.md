# Skills for 60-fifthandskip

## What You'll Learn

**Previous:** [59-wdmatch](../59-wdmatch/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Fifthandskip

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String splitting and joining

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

**Next:** [61-notdecimal](../61-notdecimal/skills.md) - Notdecimal
