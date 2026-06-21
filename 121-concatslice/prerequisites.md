# Prerequisites for concatslice

## Before You Start

To solve this challenge you need to understand:

### 1. Slices — What They Are
A slice is a dynamically sized list of elements of the same type. You can create one with a literal or `make`.

```go
a := []int{1, 2, 3}    // slice literal
b := make([]int, 0)    // empty slice with make
```

### 2. `append` — Adding Elements to a Slice
`append` returns a new slice with one or more elements added. Always assign the result back.

```go
s := []int{1, 2}
s = append(s, 3)    // s is now [1 2 3]
s = append(s, 4, 5) // s is now [1 2 3 4 5]
```

### 3. The Spread Operator `...`
The `...` after a slice unpacks it into individual arguments. This lets you append every element of one slice onto another in one call.

```go
a := []int{1, 2, 3}
b := []int{4, 5, 6}
c := append(a, b...)  // c is [1 2 3 4 5 6]
```

### 4. `for range` Over a Slice
If you want to avoid `...`, you can loop over the second slice and append element by element.

```go
for _, v := range b {
    a = append(a, v)
}
```

### 5. Function Signatures with Slice Parameters
Functions can take and return slices just like any other type.

```go
func ConcatSlice(slice1, slice2 []int) []int {
    return append(slice1, slice2...)
}
```

## Review If Stuck
- [120-concatalternate](../120-concatalternate/skills.md) — covers `append` and slice indexing

## You're Ready When You Can...
- [ ] Create a `[]int` with a literal and with `make`
- [ ] Use `append(s, value)` and always assign the result back
- [ ] Use `append(a, b...)` to join two slices
- [ ] Write a function that accepts two `[]int` parameters and returns `[]int`

## Next Steps

- [122-fprime](../122-fprime/README.md)
