# 159-closures — Closures and First-Class Functions

## Challenge

Implement in `package piscine`:

```go
// MakeCounter returns a function that increments and returns an internal count
// each time it is called, starting from 0.
func MakeCounter() func() int

// MakeAdder returns a function that adds n to any argument passed to it.
func MakeAdder(n int) func(int) int

// Filter returns a new slice containing only elements for which keep returns true.
func Filter(nums []int, keep func(int) bool) []int

// Compose returns a new function that applies g first, then f (f(g(x))).
func Compose(f, g func(int) int) func(int) int
```

## Examples

```
counter := MakeCounter()
counter()  →  1
counter()  →  2
counter()  →  3

add5 := MakeAdder(5)
add5(3)   →  8
add5(10)  →  15

Filter([]int{1,2,3,4,5}, func(n int) bool { return n%2==0 })  →  [2 4]

double := func(x int) int { return x * 2 }
inc    := func(x int) int { return x + 1 }
Compose(double, inc)(3)  →  8   // double(inc(3)) = double(4) = 8
```

Read `skills.md` before you start.
