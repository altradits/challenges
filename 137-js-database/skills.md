# Skills for 137-js-database

## What You'll Learn

**Previous:** [136-capstone](../136-capstone/skills.md)

**Challenge:** Js Database

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

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

**Next:** [138-js-frontend](../138-js-frontend/skills.md) - 138 Js Frontend
