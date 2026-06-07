# Skills for 64-findpairs

## What You'll Learn

**Previous:** [63-slice](../63-slice/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Findpairs

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

**Next:** [65-revwstr](../65-revwstr/skills.md) - Revwstr
