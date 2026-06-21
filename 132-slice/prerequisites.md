# Prerequisites for slice

## Before You Start

To solve this challenge you need to understand:

### 1. Variadic Function Parameters (`...int`)
A variadic parameter receives zero or more arguments as a slice inside the function.

```go
func example(nums ...int) {
    fmt.Println(len(nums))    // how many were passed
    fmt.Println(nums[0])      // first one (if any)
}
example(1, 2, 3)  // len=3
example()         // len=0
```

### 2. Go Slice Expressions — `s[start:end]`
Extract a sub-slice with start (inclusive) and end (exclusive) indices.

```go
a := []string{"a", "b", "c", "d", "e"}
fmt.Println(a[1:3])  // ["b" "c"]
fmt.Println(a[2:])   // ["c" "d" "e"]  (to end)
fmt.Println(a[:2])   // ["a" "b"]      (from start)
```

### 3. Negative Index Normalization
Go does not natively support negative indices. Convert them: `-1` means `len-1`, `-2` means `len-2`.

```go
func normalize(idx, length int) int {
    if idx < 0 {
        return length + idx
    }
    return idx
}
// normalize(-2, 5) → 3
// normalize(2, 5)  → 2
```

### 4. Returning `nil` for a Slice
`nil` is a valid zero value for a slice. When the indices are invalid, return `nil` (not an empty slice).

```go
return nil  // represents no result / invalid slice
```

### 5. Checking `len(nbrs)` to Handle 1 or 2 Arguments
The variadic parameter can have 1 or 2 elements. Use `len(nbrs)` to decide.

```go
if len(nbrs) == 1 {
    // one-argument form
} else {
    // two-argument form
}
```

## Review If Stuck
- [121-concatslice](../121-concatslice/skills.md) — covers slice append and basic slice operations
- [131-revconcatalternate](../131-revconcatalternate/skills.md) — covers indexed slice access

## You're Ready When You Can...
- [ ] Define a function with a variadic `...int` parameter
- [ ] Check `len(nbrs)` to determine how many arguments were passed
- [ ] Use `s[start:end]` to take a sub-slice
- [ ] Convert a negative index to a positive one using `length + idx`
- [ ] Return `nil` from a `[]string` function

## Next Steps

- [133-findpairs](../133-findpairs/README.md)
