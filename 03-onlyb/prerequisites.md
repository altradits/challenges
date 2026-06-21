# Prerequisites for 03-onlyb

## Before You Start

To solve this challenge you need to understand the basic Go program structure and how to pass a string to `fmt.Println`.

### 1. Basic Go program structure
You must have the three required parts: `package main`, `import "fmt"`, and `func main()`.

```go
package main

import "fmt"

func main() {
    fmt.Println("B")
}
```

### 2. Go is case-sensitive
`"B"` and `"b"` are different strings. `fmt.Println` and `fmt.println` are different identifiers (only the uppercase form is accessible).

```go
fmt.Println("B")  // correct: prints B
fmt.Println("b")  // wrong output: prints b
fmt.println("B")  // compile error: println is not exported
```

## Review If Stuck

- [01-only1](../01-only1/skills.md) — covers `package main`, `import "fmt"`, `func main()`, `fmt.Println()`
- [02-onlya](../02-onlya/skills.md) — covers string literals and double quotes

## You're Ready When You Can...

- [ ] Write a complete Go program with the three required parts
- [ ] Pass any string to `fmt.Println()` and predict the output
- [ ] Distinguish between uppercase and lowercase in string literals
- [ ] Explain why `fmt.Println` uses a capital `P`

## Next Steps

After completing this exercise, move to:
- [04-onlyf](../04-onlyf/README.md)
