# Skills for 44-printrevcomb

## What You'll Learn

**Previous:** [43-printmemory](../43-printmemory/skills.md)

**Challenge:** Printrevcomb

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

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

### 3. Numeric operations and type conversion

Go supports various numeric types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`.

Common operations:
- `%` (modulo) for remainders
- `/` for division (integer division truncates)
- Type conversion: `int(x)`, `float64(x)`

### 4. Pointer basics and memory addresses

Pointers hold memory addresses. Use `&` to get address, `*` to dereference:

```go
x := 42
ptr := &x    // ptr points to x
fmt.Println(*ptr)  // Prints 42 (dereferenced)
```

In Go, pointers are rarely needed for basic tasks due to pass-by-value semantics for most types.

**Next:** [45-thirdtimeisacharm](../45-thirdtimeisacharm/skills.md) - 45 Thirdtimeisacharm
