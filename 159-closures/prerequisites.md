# Prerequisites for 159-closures

## Before You Start

You should be comfortable with:

- [06-functions](../06-functions/README.md) — declaring and calling functions
- [09-slicesintro](../09-slicesintro/README.md) — working with slices
- [14-pointers](../14-pointers/README.md) — how Go manages memory

## Key Ideas to Have Solid

### Functions as Values

A function in Go is a value just like `int` or `string`. You can store it in a variable:

```go
f := func(x int) int { return x * 2 }
f(5)  // 10
```

### Function Types

The type of a function is determined by its parameters and return values:

```go
func(int) int       // takes int, returns int
func(string) bool   // takes string, returns bool
func()              // takes nothing, returns nothing
```

### What a Closure Is

A closure is a function that remembers variables from the scope where it was created — even after that scope has exited.

## You're Ready When You Can...

- [ ] Write an anonymous function and assign it to a variable
- [ ] Pass a function as a parameter to another function
- [ ] Explain what "first-class function" means
- [ ] Return a function from a function
- [ ] Use `append` to build a slice inside a loop

## Next Steps

- [160-defer-panic-recover](../160-defer-panic-recover/README.md)
