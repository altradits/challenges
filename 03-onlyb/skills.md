# Skills for 03-onlyb

## What You'll Learn

**Previous:** [02-onlya](../02-onlya/skills.md) | **Next:** [04-onlyf](../04-onlyf/skills.md)

**Challenge:** Write a program that prints the character `B` and nothing else.

## Core Concept: Case Sensitivity in Strings and Go Itself

### What Is It?

This challenge looks identical to the previous ones, but there is a subtle point: the character to print is an uppercase `B`, not lowercase `b`. This is a good opportunity to understand that Go is case-sensitive everywhere — in the language keywords, in function names, and in string values.

### How It Works

```go
package main

import "fmt"

func main() {
    fmt.Println("B")
}
```

Output:
```
B
```

### Case Sensitivity in Go: Two Layers

**Layer 1 — The language itself is case-sensitive:**

```go
func main() { }   // correct: lowercase func, main
Func Main() { }   // syntax error: Go keywords are always lowercase
```

**Layer 2 — String values are case-sensitive:**

```go
fmt.Println("B")  // prints uppercase B
fmt.Println("b")  // prints lowercase b — wrong for this challenge
```

The challenge says `B` (uppercase). If you print `b` (lowercase), you fail the test.

### Why Does Uppercase vs Lowercase Matter in Go?

In Go, the case of the first letter of a name also controls visibility (exported vs unexported):

```go
fmt.Println   // uppercase P — exported, usable from outside the fmt package
fmt.println   // lowercase p — unexported, not accessible outside fmt
```

This is why you write `fmt.Println` not `fmt.println`. You will encounter this pattern throughout your Go journey.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `fmt.Println("b")` lowercase | Prints `b`, challenge expects `B` | Use `"B"` |
| `fmt.println("B")` | `println` is unexported — compile error | Use `fmt.Println` (capital P) |

## Solving This Challenge

### Algorithm

1. Declare `package main`
2. Import `fmt`
3. Inside `func main()`, call `fmt.Println("B")`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [04-onlyf](../04-onlyf/README.md)
