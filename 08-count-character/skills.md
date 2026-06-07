# Skills for 08-count-character

## What You'll Learn

**Previous:** [06-checknumber](../06-checknumber/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Count Character

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

### 3. Conditional logic and boolean returns

Go uses `if/else` for conditional branching. The condition doesn't need parentheses:

```go
if condition {
    // do something
} else if otherCondition {
    // do something else
} else {
    // default case
}
```

Boolean operators: `&&` (AND), `||` (OR), `!` (NOT).

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

**Next:** [10-findsubstring](../10-findsubstring/skills.md) - Findsubstring
