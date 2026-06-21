# Prerequisites for fprime

## Before You Start

To solve this challenge you need to understand:

### 1. `strconv.Atoi` — Parsing a String to Integer
The input comes as a string from `os.Args`. You must convert it to an `int` to do arithmetic.

```go
import "strconv"

n, err := strconv.Atoi("42")
if err != nil {
    // not a valid integer
}
fmt.Println(n) // 42
```

### 2. `os.Args` — Reading Command-Line Arguments
`os.Args[0]` is the program name; `os.Args[1]` is the first argument. Always check `len(os.Args)` first.

```go
import "os"

if len(os.Args) != 2 {
    return // wrong number of arguments
}
arg := os.Args[1]
```

### 3. The Modulo Operator `%`
`n % d` gives the remainder of `n / d`. If the remainder is 0, `d` divides `n` evenly.

```go
fmt.Println(10 % 3) // 1  — not divisible
fmt.Println(9 % 3)  // 0  — divisible
```

### 4. Nested `for` Loops
You need an outer loop (over divisor candidates) and an inner loop (dividing out all copies of the same factor).

```go
d := 2
for d*d <= n {       // outer: try each divisor
    for n%d == 0 {   // inner: collect all copies of d
        n /= d
    }
    d++
}
```

### 5. Collecting Results in a Slice
Accumulate factors in a `[]int`, then print them joined by `*`.

```go
factors := []int{}
factors = append(factors, d)
```

### 6. Printing with a Separator
Print the first element without a `*`, then precede every subsequent element with `*`.

```go
for i, f := range factors {
    if i > 0 {
        fmt.Print("*")
    }
    fmt.Print(f)
}
fmt.Println()
```

## Review If Stuck
- [121-concatslice](../121-concatslice/skills.md) — covers building slices with `append`
- Prior `os.Args` challenges — covers argument parsing

## You're Ready When You Can...
- [ ] Use `strconv.Atoi` and check the error
- [ ] Check `len(os.Args)` and return early if wrong
- [ ] Use `%` to test divisibility
- [ ] Write a `for d*d <= n` loop
- [ ] Append to a `[]int` and print elements with a `*` separator

## Next Steps

- [123-hiddenp](../123-hiddenp/README.md)
