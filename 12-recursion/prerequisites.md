# Prerequisites for 12-recursion

## Before You Start

To solve this challenge you need to understand functions, conditionals, and basic arithmetic.

### 1. Defining and Calling Functions

You must be comfortable writing a function that takes a parameter and returns a value. See [06-functions](../06-functions/skills.md) for a refresher.

```go
func Double(n uint) uint {
    return n * 2
}
```

### 2. Return Statements

A recursive function returns a value computed from a call to itself. The return type of the recursive call must match the function's own return type.

```go
func Factorial(n uint) uint {
    // ...
    return n * Factorial(n-1)  // returns uint — matches the signature
}
```

### 3. If/Else Conditionals

The base case of a recursive function is written as an `if` statement. Without it the function calls itself forever.

```go
if n == 0 {
    return 1  // stop here
}
```

### 4. Arithmetic Operators: `*` and `%`

Factorial uses multiplication (`*`). The next challenge (GCD) uses the modulo operator (`%`).

```go
5 * 4   // 20
10 % 3  // 1  (remainder after division)
```

### 5. The `uint` Type

`uint` is an **unsigned integer** — it can only hold zero or positive values. The function signature for this challenge uses `uint`, so all your variables and return values must also be `uint`.

```go
var n uint = 5
fmt.Println(n * 2)  // 10
```

## Review If Stuck

- [06-functions](../06-functions/skills.md) — covers defining functions with parameters and return values
- [85-printif](../85-printif/skills.md) — covers `if/else` conditionals

## You're Ready When You Can...

- [ ] Define a function with a `uint` parameter that returns a `uint`
- [ ] Write a function that calls itself (try writing `func f(n uint) uint { return f(n) }` and understand why it crashes)
- [ ] Write an `if/else` condition that checks `n == 0`
- [ ] Use `*` to multiply two `uint` values

## Next Steps

- [13-structs](../13-structs/README.md)
