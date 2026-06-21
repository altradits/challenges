# Prerequisites for 05-onlyz

## Before You Start

This challenge requires the same skills as 01 through 04. You should be able to write the program skeleton from memory.

### 1. The complete Go program skeleton

```go
package main

import "fmt"

func main() {
    fmt.Println("z")
}
```

All four of these elements are required:
- `package main` — marks this as an executable
- `import "fmt"` — loads the print functions
- `func main()` — entry point, called automatically
- `fmt.Println(...)` — prints a value followed by a newline

### 2. String literals and case sensitivity

```go
fmt.Println("z")   // correct: lowercase z
fmt.Println("Z")   // wrong: uppercase Z
fmt.Println('z')   // wrong: prints 122 (the rune value of z)
```

## Review If Stuck

- [01-only1](../01-only1/skills.md) — covers the full skeleton for the first time
- [02-onlya](../02-onlya/skills.md) — string literals (double vs single quotes)
- [03-onlyb](../03-onlyb/skills.md) — case sensitivity
- [04-onlyf](../04-onlyf/skills.md) — `fmt.Println` vs `fmt.Print` vs `fmt.Printf`

## You're Ready When You Can...

- [ ] Write the full Go program structure from memory, without copying
- [ ] Correctly predict what `fmt.Println("z")` will output
- [ ] Explain the difference between `"z"` and `'z'` in Go

## Next Steps

After completing this exercise, move to:
- [06-checknumber](../06-checknumber/README.md) — your first challenge with real logic
