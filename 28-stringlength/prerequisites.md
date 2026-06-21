# Prerequisites for 28-stringlength

## Before You Start

This is the **first function-based challenge** in the series. You do not need prior Go experience beyond the basics below.

### 1. What a Go Function Is

A function is a reusable block of code with a name, inputs (parameters), and an output (return value). In Go, you define one like this:

```go
func Add(a int, b int) int {
    return a + b
}
```

- `func` — keyword that starts every function definition
- `Add` — the name (must start with a capital letter to be visible from tests)
- `(a int, b int)` — two parameters, both of type `int`
- `int` after the `)` — the return type
- `return a + b` — sends the result back to whoever called the function

For this challenge the signature is already given:

```go
func StringLength(s string) int {
    // your code here
}
```

### 2. Variable Declaration and Increment

You will need a counter variable:

```go
count := 0   // declare and initialise with :=
count++      // add 1 (shorthand for count = count + 1)
```

### 3. The `for` Loop — Basic Form

A counting loop in Go:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)   // prints 0 1 2 3 4
}
```

- Part 1 (`i := 0`) runs once at the start
- Part 2 (`i < 5`) is checked before every iteration
- Part 3 (`i++`) runs after every iteration

### 4. Package and Import Basics

Every Go file starts with a package declaration. For these challenges, that is `package main`. If you need `fmt.Println` you import it:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

The function you write (`StringLength`) and the test file live in the same package, so you do not need to import anything for the function itself.

## Review If Stuck

- If basic Go syntax is new to you, the [Go Tour](https://go.dev/tour/) covers variables and loops in its first few lessons.

## You're Ready When You Can...

- [ ] Write a Go function that takes a `string` parameter and returns an `int`
- [ ] Declare an integer variable and increment it inside a loop
- [ ] Explain the difference between a byte and a character in a string

## Next Steps

- [29-firstchar](../29-firstchar/README.md)
