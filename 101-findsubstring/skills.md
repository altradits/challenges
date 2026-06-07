# Skills for 101-findsubstring

## What You'll Learn

**Previous:** [100-countwords](../100-countwords/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Findsubstring

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

### 3. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

### 4. Formatted output with fmt package

The `fmt` package provides formatted I/O:

```go
fmt.Println("Hello")     // Print with newline
fmt.Printf("Value: %d", x)  // Formatted print
fmt.Scan(&x)             // Read input
```

Common verbs: `%d` (int), `%s` (string), `%v` (any value), `%T` (type)

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [102-replaceall](../102-replaceall/skills.md) - Replaceall
