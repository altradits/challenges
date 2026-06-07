# Skills for 52-concatslice

## What You'll Learn

**Previous:** [51-concatalternate](../51-concatalternate/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Concatslice

## New Concepts Explained

### 1. Numeric operations and type conversion

Go supports various numeric types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`.

Common operations:
- `%` (modulo) for remainders
- `/` for division (integer division truncates)
- Type conversion: `int(x)`, `float64(x)`

### 2. Slice manipulation and operations

Slices are dynamic, flexible views into arrays. They're the most common data structure in Go:

```go
// Create a slice
numbers := []int{1, 2, 3, 4, 5}

// Slice an existing slice
subset := numbers[1:4]  // [2, 3, 4]

// Append to a slice
numbers = append(numbers, 6)
```

Slices have length (current elements) and capacity (max elements without reallocation).

### 3. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

### 4. Conditional logic and boolean returns

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

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [53-fprime](../53-fprime/skills.md) - Fprime
