# Skills for 49-canjump

## What You'll Learn

**Previous:** [48-addprimesum](../48-addprimesum/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Canjump

## New Concepts Explained

### 1. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

### 2. Formatted output with fmt package

The `fmt` package provides formatted I/O:

```go
fmt.Println("Hello")     // Print with newline
fmt.Printf("Value: %d", x)  // Formatted print
fmt.Scan(&x)             // Read input
```

Common verbs: `%d` (int), `%s` (string), `%v` (any value), `%T` (type)

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [50-chunk](../50-chunk/skills.md) - Chunk
