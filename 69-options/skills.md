# Skills for 69-options

## What You'll Learn

**Previous:** [68-itoabase](../68-itoabase/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Options

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. Bit manipulation and binary operations

Go supports bitwise operations on integers:
- `&` (AND) - both bits set
- `|` (OR) - either bit set
- `^` (XOR) - exactly one bit set
- `&^` (AND NOT) - clear bits
- `<<` / `>>` - shift left/right

```go
// Check if a bit is set
if flags&(1<<i) != 0 {
    // bit i is set
}

// Set a bit
flags |= 1 << i

// Clear a bit
flags &^= 1 << i
```

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

### 4. Formatted output with fmt package

The `fmt` package provides formatted I/O:

```go
fmt.Println("Hello")     // Print with newline
fmt.Printf("Value: %d", x)  // Formatted print
fmt.Scan(&x)             // Read input
```

Common verbs: `%d` (int), `%s` (string), `%v` (any value), `%T` (type)

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [70-piglatin](../70-piglatin/skills.md) - Piglatin
