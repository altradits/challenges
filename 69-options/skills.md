# Skills for 69-options

## What You'll Learn

**Previous:** [68-itoabase](../68-itoabase/skills.md)

**Challenge:** Options

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

### 2. Numeric operations and type conversion

Go supports various numeric types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `float32`, `float64`.

Common operations:
- `%` (modulo) for remainders
- `/` for division (integer division truncates)
- Type conversion: `int(x)`, `float64(x)`

### 3. Pointer basics and memory addresses

Pointers hold memory addresses. Use `&` to get address, `*` to dereference:

```go
x := 42
ptr := &x    // ptr points to x
fmt.Println(*ptr)  // Prints 42 (dereferenced)
```

In Go, pointers are rarely needed for basic tasks due to pass-by-value semantics for most types.

### 4. Command-line argument handling

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

**Next:** [70-piglatin](../70-piglatin/skills.md) - 70 Piglatin
