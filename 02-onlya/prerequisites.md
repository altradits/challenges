# Prerequisites for 02-onlya

## Before You Start

To solve this challenge you need to understand the basic Go program structure introduced in 01-only1.

### 1. Basic Go program structure
Every Go executable needs `package main`, an import, and `func main()`.

```go
package main

import "fmt"

func main() {
    fmt.Println("something")
}
```

### 2. String literals use double quotes
In Go, text values are wrapped in double quotes. Single quotes mean something different (they create a rune, which is an integer).

```go
fmt.Println("a")   // prints: a
fmt.Println('a')   // prints: 97  (NOT what you want here)
```

## Review If Stuck

- [01-only1](../01-only1/skills.md) — covers `package main`, `import "fmt"`, `func main()`, and `fmt.Println()`

## You're Ready When You Can...

- [ ] Write the three-part Go program skeleton from memory
- [ ] Pass a string argument to `fmt.Println()`
- [ ] Know the difference between `"a"` (string) and `'a'` (rune)
- [ ] Run `go run .` and read the output

## Next Steps

After completing this exercise, move to:
- [03-onlyb](../03-onlyb/README.md)
