# Skills for 53-fprime

## What You'll Learn

**Previous:** [52-concatslice](../52-concatslice/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Fprime

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

### 3. Error handling and validation

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

**Next:** [54-hiddenp](../54-hiddenp/skills.md) - Hiddenp
