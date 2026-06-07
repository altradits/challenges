# Skills for 71-romannumbers

## What You'll Learn

**Previous:** [70-piglatin](../70-piglatin/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Romannumbers

## New Concepts Explained

### 1. Numeric operations and type conversion

Go supports various numeric types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`.

Common operations:
- `%` (modulo) for remainders
- `/` for division (integer division truncates)
- Type conversion: `int(x)`, `float64(x)`

### 2. Looping constructs (for, range)

Go has only one looping construct: the `for` loop. It can be used in several ways:

```go
// Traditional for loop
for i := 0; i < 10; i++ { }

// While-style loop
for condition { }

// Range loop (for collections)
for index, value := range collection { }
```

For strings, `for...range` iterates over runes, making it safe for UTF-8.

### 3. Formatted output with fmt package

The `fmt` package provides formatted I/O:

```go
fmt.Println("Hello")     // Print with newline
fmt.Printf("Value: %d", x)  // Formatted print
fmt.Scan(&x)             // Read input
```

Common verbs: `%d` (int), `%s` (string), `%v` (any value), `%T` (type)

### 4. Error handling and validation

Go handles errors explicitly. Functions often return `(value, error)`:

```go
result, err := someFunction()
if err != nil {
    // handle error
    return
}
// use result
```

Always check errors - Go doesn't have exceptions!

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [72-brackets](../72-brackets/skills.md) - Brackets
