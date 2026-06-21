# Prerequisites for 11-forloops

## Before You Start

To solve this challenge you need to understand variables and short declaration, conditionals, and Go functions.

### 1. Variables and short declaration (`:=`)

The `:=` operator declares a new variable and assigns a value in one step:

```go
i := 5      // declares i as int with value 5
i--         // decrements i by 1 (i is now 4)
i++         // increments i by 1
```

Inside a for loop's initialization, you always use `:=` to declare the loop variable:

```go
for i := 5; i >= 0; i-- {
    // i starts at 5 and decreases each iteration
}
```

### 2. Comparison operators

The loop condition uses comparison operators to decide whether to keep running:

```go
i >= 0   // true while i is greater than or equal to 0
i < n    // true while i is less than n
i == 0   // true only when i equals 0
```

When the condition becomes false, the loop stops.

### 3. Conditionals (`if/else`) from 06-checknumber or 09

You have already used `if` to make decisions. Inside a loop you can use `if` to skip iterations or exit early:

```go
if i%2 == 0 {
    continue   // skip even numbers
}
```

### 4. Functions from 07-functions

This challenge asks you to write a function in `package piscine`. You need to know:

- `package piscine` at the top of the file
- `func CountDown(n int)` — a function that takes an `int` and returns nothing
- A function with no return type has no `return` statement at the end (or a bare `return`)

```go
package piscine

import "fmt"

func CountDown(n int) {
    // loop and print here
}
```

### 5. `fmt.Println` for printing integers

`fmt.Println` can print integers directly — it converts them to their string representation:

```go
fmt.Println(5)    // prints: 5
fmt.Println(0)    // prints: 0
```

## Review If Stuck

- [07-functions](../07-functions/skills.md) — covers function syntax and `package piscine`
- [06-checknumber](../06-checknumber/skills.md) — covers `for...range` on strings (related loop pattern)

## You're Ready When You Can...

- [ ] Declare a loop variable with `i := n` and count it down with `i--`
- [ ] Write a `for` loop condition that stops at (and includes) `0`
- [ ] Call `fmt.Println` with an integer argument inside a loop
- [ ] Write a function with an `int` parameter and no return value
- [ ] Explain the three parts of `for init; condition; post { }`

## Next Steps

After completing this exercise, move to:
- [12-printif](../12-printif/README.md)
