# Prerequisites for 04-onlyf

## Before You Start

To solve this challenge you need to understand the basic Go program structure and how to print a string.

### 1. Basic Go program structure
Three parts: `package main`, `import "fmt"`, `func main()`.

```go
package main

import "fmt"

func main() {
    fmt.Println("f")
}
```

### 2. `fmt.Println` adds a newline automatically
You do not need to add `\n` yourself when using `Println`.

```go
fmt.Println("f")   // prints f and then a newline
fmt.Print("f\n")   // same result, but manual newline
```

### 3. String values are case-sensitive
The challenge asks for lowercase `f`. `"f"` and `"F"` are different.

```go
fmt.Println("f")   // correct output
fmt.Println("F")   // wrong output
```

## Review If Stuck

- [01-only1](../01-only1/skills.md) — covers the full three-part program structure
- [02-onlya](../02-onlya/skills.md) — covers string literals
- [03-onlyb](../03-onlyb/skills.md) — covers case sensitivity

## You're Ready When You Can...

- [ ] Write the three-part Go program structure without looking it up
- [ ] Predict exactly what `fmt.Println("f")` will print
- [ ] Know that `fmt.Println` automatically appends a newline

## Next Steps

- [05-onlyz](../05-onlyz/README.md)
