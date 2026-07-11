# Skills for 162-variadic-functions

## What You'll Learn

**Previous:** [161-switch-statements](../161-switch-statements/skills.md) | **Next:** [163-type-conversions](../163-type-conversions/skills.md)

**Challenge:** Write functions that accept any number of arguments using the `...` syntax.

## Variadic Functions

A **variadic** function accepts a variable number of arguments. The last parameter uses `...type`:

```go
func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

Sum(1, 2, 3)        // 6
Sum(10, 20)         // 30
Sum()               // 0
```

Inside the function, `nums` is a **slice** of `int`. You can use `len(nums)`, `range`, indexing — anything you would do with a slice.

### Calling With a Slice

Use `...` at the call site to expand a slice into individual arguments:

```go
nums := []int{1, 2, 3, 4}
Sum(nums...)   // same as Sum(1, 2, 3, 4)
```

This is how `fmt.Println` and `append` work:

```go
parts := []interface{}{"hello", "world"}
fmt.Println(parts...)   // hello world

a := []int{1, 2}
b := []int{3, 4}
a = append(a, b...)    // a == [1, 2, 3, 4]
```

### Position Rule

The variadic parameter must be the **last** parameter:

```go
func Join(sep string, parts ...string) string { ... }   // OK
// func Bad(nums ...int, sep string) { }                  // compile error
```

### Nil and Empty

```go
Sum()         // nums is []int{} — len is 0, range is fine
Sum(nil...)   // only works if the type is a reference type
```

For integer/string variadics, calling with no arguments gives you an empty slice — safe to range over, not nil.

## Standard Library Examples

The variadic pattern appears everywhere:

```go
fmt.Printf(format string, a ...interface{})
fmt.Println(a ...interface{})
strings.Join(elems []string, sep string)  // not variadic, but similar intent
append(s []T, elems ...T) []T
```

## Visual: How `...` Expands

```
Call: Sum(1, 2, 3)
               ↓
Function: func Sum(nums ...int)
           nums = []int{1, 2, 3}

Call: Sum(slice...)
               ↓
Function: func Sum(nums ...int)
           nums = slice  (same underlying array)
```

When you expand a slice with `...`, the variadic receives the same backing array — no copy is made. Modifying elements inside the variadic function would mutate the original (but this is unusual to rely on).

## Solving the Challenge

```go
package piscine

import "strings"

func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func Max(nums ...int) int {
    if len(nums) == 0 {
        panic("Max requires at least one argument")
    }
    m := nums[0]
    for _, n := range nums[1:] {
        if n > m {
            m = n
        }
    }
    return m
}

func Join(sep string, parts ...string) string {
    return strings.Join(parts, sep)
}

func Wrap(fn func(int) int, nums ...int) []int {
    result := make([]int, len(nums))
    for i, n := range nums {
        result[i] = fn(n)
    }
    return result
}
```

## Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `Sum([]int{1,2,3})` | Passes a slice as a single argument | Use `Sum([]int{1,2,3}...)` |
| `func f(a ...int, b string)` | Variadic must be last | Reorder: `func f(b string, a ...int)` |
| Assuming variadic is never nil | Empty variadics give `[]T{}`, not nil | Use `len(nums) == 0` not `nums == nil` |

## Documentation

- [A Tour of Go — Variadic functions](https://go.dev/tour/moretypes/15)
- [Go Spec — Passing arguments to ... parameters](https://go.dev/ref/spec#Passing_arguments_to_..._parameters)
- [strings.Join](https://pkg.go.dev/strings#Join)
- [builtin.append](https://pkg.go.dev/builtin#append)

**Next:** [163-type-conversions](../163-type-conversions/README.md)
