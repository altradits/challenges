# Skills for 159-closures

## What You'll Learn

**Previous:** [27-context](../27-context/skills.md) | **Next:** [160-defer-panic-recover](../160-defer-panic-recover/skills.md)

**Challenge:** Use closures and first-class functions to build counters, adders, filters, and function composers.

## Core Concept: Functions Are Values

In Go, functions are **first-class values** — you can assign them to variables, pass them as arguments, and return them from other functions. This unlocks patterns impossible with plain functions.

```go
// Assign a function to a variable
double := func(x int) int { return x * 2 }
fmt.Println(double(5))  // 10

// Pass a function as an argument
func Apply(n int, f func(int) int) int {
    return f(n)
}
Apply(5, double)  // 10

// Return a function from a function
func Multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
triple := Multiplier(3)
triple(4)  // 12
```

## What Is a Closure?

A **closure** is a function that *captures* variables from the scope where it was created. The captured variable persists as long as the closure exists.

```go
func MakeCounter() func() int {
    count := 0              // this variable lives inside the closure
    return func() int {
        count++             // "closes over" count — modifies the outer variable
        return count
    }
}

c := MakeCounter()
c()  // 1
c()  // 2  — count persists between calls!
c()  // 3
```

Each call to `MakeCounter()` creates a **new, independent** count:

```go
c1 := MakeCounter()
c2 := MakeCounter()
c1()  // 1
c1()  // 2
c2()  // 1  — c2 has its own count
```

### Visual: Closure Memory Model

```
MakeCounter() call 1:
  [ stack frame ]    [ heap ]
  count → ─────────▶  int: 0
                         ▲
  returned func ────────┘  (captures the address of count)

After c1() called:
  [ heap ]
    int: 1   ← count was incremented through the closure
```

## Function Types

Every function has a **type** based on its parameter and return types:

```go
var f func(int) int         // a function taking int, returning int
var g func(string) bool     // a function taking string, returning bool
var h func()                // a function with no params, no return
```

When you write `func(x int) int { return x * 2 }` you are creating a value of type `func(int) int`.

## Higher-Order Functions

Functions that accept or return other functions are called **higher-order functions**. They are the foundation of functional-style Go:

```go
// Map applies fn to every element and returns new slice
func Map(nums []int, fn func(int) int) []int {
    result := make([]int, len(nums))
    for i, v := range nums {
        result[i] = fn(v)
    }
    return result
}

Map([]int{1, 2, 3}, func(n int) int { return n * n })
// [1, 4, 9]
```

```go
// Reduce folds a slice into a single value
func Reduce(nums []int, initial int, fn func(int, int) int) int {
    acc := initial
    for _, v := range nums {
        acc = fn(acc, v)
    }
    return acc
}

Reduce([]int{1, 2, 3, 4}, 0, func(a, b int) int { return a + b })
// 10
```

## Function Composition

Composing `f` after `g` means applying `g` first, then `f`:

```go
func Compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}
```

This is the mathematical definition of `(f ∘ g)(x) = f(g(x))`.

## Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Closure captures loop variable | All closures see the final value | Pass loop var as parameter to the closure |
| Forgetting to call returned function | `counter` is the function, `counter()` calls it | Always `()` to invoke |
| Mutation side effects | Multiple closures sharing a variable affect each other | Create independent closures with separate `MakeCounter()` calls |

### Loop Variable Capture Bug

```go
// BUG
funcs := make([]func(), 3)
for i := 0; i < 3; i++ {
    funcs[i] = func() { fmt.Println(i) }
}
funcs[0]()  // 3, not 0 — all closures captured the same i

// FIX: shadow the variable inside the loop
for i := 0; i < 3; i++ {
    i := i  // new i scoped to this iteration
    funcs[i] = func() { fmt.Println(i) }
}
// OR: pass as parameter
for i := 0; i < 3; i++ {
    funcs[i] = func(n int) func() { return func() { fmt.Println(n) } }(i)
}
```

## Solving the Challenge

```go
package piscine

func MakeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func MakeAdder(n int) func(int) int {
    return func(x int) int {
        return x + n
    }
}

func Filter(nums []int, keep func(int) bool) []int {
    result := []int{}
    for _, v := range nums {
        if keep(v) {
            result = append(result, v)
        }
    }
    return result
}

func Compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}
```

## Documentation

- [A Tour of Go — Function values](https://go.dev/tour/moretypes/24)
- [A Tour of Go — Closures](https://go.dev/tour/moretypes/25)
- [Go Spec — Function types](https://go.dev/ref/spec#Function_types)

**Next:** [160-defer-panic-recover](../160-defer-panic-recover/README.md)
