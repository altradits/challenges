# Skills for 44-printrevcomb

## What You'll Learn

**Previous:** [43-printmemory](../43-printmemory/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Printrevcomb

## New Concepts Explained

### 1. Numeric operations and type conversion

Go supports various numeric types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`.

Common operations:
- `%` (modulo) for remainders
- `/` for division (integer division truncates)
- Type conversion: `int(x)`, `float64(x)`

### 2. Conditional logic and boolean returns

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

### 3. Formatted output with fmt package

The `fmt` package provides formatted I/O:

```go
fmt.Println("Hello")     // Print with newline
fmt.Printf("Value: %d", x)  // Formatted print
fmt.Scan(&x)             // Read input
```

Common verbs: `%d` (int), `%s` (string), `%v` (any value), `%T` (type)

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [45-thirdtimeisacharm](../45-thirdtimeisacharm/skills.md) - Thirdtimeisacharm
