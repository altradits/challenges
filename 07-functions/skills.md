# Skills for 07-functions

## What You'll Learn

**Previous:** [05-onlyz](../05-onlyz/skills.md) | **Next:** [06-checknumber](../06-checknumber/skills.md)

**Challenge:** Write a function `Greet(name string) string` that returns `"Hello, " + name + "!"`.

## Core Concept: Writing a Go Function

### What Is a Function?

A function is a named, reusable block of code that can accept inputs (parameters) and produce an output (return value). Every Go program you have written so far has used one function — `main`. Now you will write your own.

### Function Syntax

```go
func Name(param type) returnType {
    return value
}
```

Breaking that down:

| Part | Example | Meaning |
|------|---------|---------|
| `func` | `func` | keyword that declares a function |
| `Name` | `Greet` | the function's name (must start with a capital letter to be exported) |
| `param type` | `name string` | the input: parameter name and its type |
| `returnType` | `string` | the type of value the function sends back |
| `return value` | `return "Hello, " + name + "!"` | the value the function sends back to its caller |

### A Complete Example

```go
func Greet(name string) string {
    return "Hello, " + name + "!"
}
```

Calling it:

```go
result := Greet("World")   // result == "Hello, World!"
```

### Parameters and Return Types

**Parameters** are the inputs a function receives. You list each parameter as `name type`:

```go
func Add(a int, b int) int {
    return a + b
}

// Shorthand when two parameters share the same type:
func Add(a, b int) int {
    return a + b
}
```

**Return types** come after the parameter list. A function must return a value that matches its declared return type:

```go
func Double(n int) int {
    return n * 2
}

func IsEmpty(s string) bool {
    return s == ""
}
```

### Package Declaration: `package piscine`

All challenges from 06 onward ask you to write a function in `package piscine`, not `package main`. This is the shift from "a standalone program" to "a library function".

- `package main` is for executable programs — it has a `func main()` entry point.
- `package piscine` (or any other name) is for library code — it exposes functions for other packages to import and call.

When you solve these challenges, your solution file starts with:

```go
package piscine
```

### How Function Challenges Are Tested

The test harness lives in a separate file in `package main`. It imports `piscine` and calls your function:

```go
// Your file (piscine package):
package piscine

func Greet(name string) string {
    return "Hello, " + name + "!"
}

// Test file (main package):
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.Greet("World"))  // Hello, World!
}
```

Notice that `piscine.Greet` uses the package name as a prefix. This is how Go calls an exported function from another package.

### Exported vs Unexported Names

A name that starts with a **capital letter** is **exported** — it can be used from other packages:

```go
func Greet(name string) string { ... }   // exported — visible to package main
func greet(name string) string { ... }   // unexported — only visible inside piscine
```

For all challenges, use a capital first letter to match the expected function signature.

### Multiple Return Values

Go functions can return more than one value. This is common for returning a result alongside an error:

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

result, err := Divide(10, 2)
if err != nil {
    fmt.Println("error:", err)
}
```

The caller receives both values and must handle them. You will see this pattern throughout Go's standard library.

### Named Return Values

Go also lets you name the return values, which gives them zero values automatically and lets you use a bare `return`:

```go
func MinMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return   // returns min and max
}
```

Named returns are useful in longer functions but can reduce readability in short ones. Prefer explicit `return value` for clarity.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `package main` in solution file | Function is in the wrong package | Use `package piscine` |
| Lowercase function name `greet` | Function is unexported; test can't call it | Capitalize: `Greet` |
| Missing `return` statement | Compile error: missing return | Every code path must `return` a value |
| `return "Hello " + name` (no comma) | Wrong output format | Match the required format exactly |
| Calling `Greet("World")` without package prefix | Works inside `piscine` but fails from `main` | Use `piscine.Greet("World")` from outside |

## Solving This Challenge

### Algorithm

1. Declare `package piscine`
2. Write `func Greet(name string) string`
3. Return `"Hello, " + name + "!"`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [06-checknumber](../06-checknumber/README.md)
