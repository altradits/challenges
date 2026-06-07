# Skills for 48-addprimesum

## What You'll Learn

**Previous:** [47-zipstring](../47-zipstring/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Addprimesum

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

### 3. Command-line argument handling

Access command-line arguments via `os.Args`:

```go
import "os"

func main() {
    args := os.Args[1:]  // Skip program name
    for _, arg := range args {
        fmt.Println(arg)
    }
}
```

Or use the `flag` package for more complex argument parsing.

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [49-canjump](../49-canjump/skills.md) - Canjump
