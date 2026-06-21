# Skills for 04-onlyf

## What You'll Learn

**Previous:** [03-onlyb](../03-onlyb/skills.md) | **Next:** [05-onlyz](../05-onlyz/skills.md)

**Challenge:** Write a program that prints the character `f` and nothing else.

## Core Concept: `fmt.Print` vs `fmt.Println` vs `fmt.Printf`

### What Is It?

By now the basic structure (`package main`, `import "fmt"`, `func main()`) should feel familiar. This challenge is a good moment to understand the three main print functions in the `fmt` package, so you know which one to reach for in future challenges.

### The Three Print Functions

#### `fmt.Println` — print with newline

```go
fmt.Println("f")
```

Prints `f` followed by a newline character. Most beginner programs use this.

#### `fmt.Print` — print without adding a newline

```go
fmt.Print("f")
```

Prints `f` with no automatic newline. If you call it twice:

```go
fmt.Print("f")
fmt.Print("!")
// Output: f!   (on one line)
```

#### `fmt.Printf` — print with formatting

```go
fmt.Printf("%s\n", "f")
fmt.Printf("The letter is: %s\n", "f")
```

Uses a format string. `%s` is the verb for strings, `%d` for integers. The `\n` adds a newline manually.

### For This Challenge

`fmt.Println("f")` is the simplest correct solution:

```go
package main

import "fmt"

func main() {
    fmt.Println("f")
}
```

Output:
```
f
```

### When to Use Each

| Function | Use When |
|----------|----------|
| `fmt.Println` | You want a newline added automatically |
| `fmt.Print` | You want to control newlines yourself |
| `fmt.Printf` | You need to embed variables into a formatted string |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `fmt.Println("F")` uppercase | Prints `F`, challenge expects `f` | Use lowercase `"f"` |
| `fmt.Printf("f")` without `\n` | Correct output, but no newline — may cause issues in tests | Use `fmt.Println("f")` or add `\n` |

## Solving This Challenge

### Algorithm

1. Declare `package main`
2. Import `fmt`
3. Inside `func main()`, call `fmt.Println("f")`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [05-onlyz](../05-onlyz/README.md)
