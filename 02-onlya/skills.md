# Skills for 02-onlya

## What You'll Learn

**Previous:** [01-only1](../01-only1/skills.md) | **Next:** [03-onlyb](../03-onlyb/skills.md)

**Challenge:** Write a program that prints the character `a` and nothing else.

## Core Concept: Printing a Single Character

### What Is It?

This challenge reinforces the basic Go program structure you learned in 01-only1. The only thing that changes is what you pass to `fmt.Println`. You are practicing that the argument to `fmt.Println` is just a value — it can be any string.

### How It Works

The program structure is identical to 01-only1:

```go
package main

import "fmt"

func main() {
    fmt.Println("a")
}
```

Output:
```
a
```

### String Literals vs Character Literals

In Go there is an important distinction between double quotes and single quotes:

```go
fmt.Println("a")   // "a" is a string — this works with Println
fmt.Println('a')   // 'a' is a rune (integer value 97) — Println prints 97
```

For this challenge you want to print the letter `a`, so use double quotes: `"a"`.

If you use single quotes, `fmt.Println('a')` will print `97` (the ASCII value of `a`), which is not what the challenge asks for.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `fmt.Println('a')` with single quotes | Prints `97` (the rune value), not `a` | Use double quotes: `fmt.Println("a")` |
| `fmt.Println("A")` | Wrong case — challenge asks for lowercase `a` | Use `"a"` |
| Copy-pasting 01-only1 and forgetting to change the value | Prints `1` instead of `a` | Update the argument |

## Solving This Challenge

### Algorithm

1. Declare `package main`
2. Import `fmt`
3. Inside `func main()`, call `fmt.Println("a")`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [03-onlyb](../03-onlyb/README.md)
