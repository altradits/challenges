# Prerequisites for 08-stringspackage

## Before You Start

To solve this challenge you need to understand strings, for loops, and the import statement.

### 1. String Literals and the `string` Type

A string is a sequence of characters enclosed in double quotes. You can pass strings to functions and return them.

```go
s := "Hello, world"
fmt.Println(len(s))  // 12 — number of bytes
```

### 2. The `import` Statement

To use a standard library package, name it in your import block. Multiple packages go inside parentheses.

```go
import (
    "fmt"
    "strings"
)
```

Once imported, call functions with the package name as a prefix: `strings.ToLower(...)`.

See [01-only1](../01-only1/skills.md) for a refresher on import.

### 3. For Loops Over Strings

You can iterate over a string with `range` to visit each character (rune):

```go
for _, r := range s {
    fmt.Println(r)
}
```

See [07-forloops](../07-forloops/skills.md) for a refresher on for loops.

### 4. Multiple Return Values

Go functions can return more than one value. Declare them in parentheses in the signature:

```go
func stats(s string) (int, int, bool) {
    return 0, 0, false
}

a, b, c := stats("hello")
```

### 5. Slices and `len`

`strings.Fields` returns a `[]string` (slice of strings). Use `len` to get its length:

```go
import "strings"

words := strings.Fields("hello world")
fmt.Println(len(words))  // 2
```

## Review If Stuck

- [01-only1](../01-only1/skills.md) — covers the import statement
- [07-forloops](../07-forloops/skills.md) — covers for loops and range

## You're Ready When You Can...

- [ ] Write an import block that imports `"strings"`
- [ ] Call `strings.ToLower` on a string and use the result
- [ ] Use `strings.Fields` and check the length of the returned slice
- [ ] Write a function that returns multiple values

## Next Steps

After completing this exercise, move to:
- [09-slicesintro](../09-slicesintro/README.md)
