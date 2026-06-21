# Prerequisites for concatalternate

## Before You Start

To solve this challenge you need to understand:

### 1. Slices and `append`
A slice is a dynamically sized list. Build the result by appending elements one at a time.

```go
result := []int{}
result = append(result, 42)  // [42]
result = append(result, 7)   // [42, 7]
```

### 2. `len()` on a Slice
`len(s)` returns the number of elements. You need this to decide which input slice is longer.

```go
a := []int{1, 2, 3}
b := []int{4, 5}
fmt.Println(len(a)) // 3
fmt.Println(len(b)) // 2
```

### 3. Slice Indexing and Offsets
Access any element by index (starting at 0). When slices differ in length, compute an offset to find where the "equal-length" tail begins.

```go
s := []int{10, 20, 30, 40}
fmt.Println(s[0]) // 10
fmt.Println(s[2]) // 30
```

### 4. Indexed `for` Loop
Use a classic `for` loop (not `range`) when you need to walk two slices in parallel with offsets.

```go
for i := 0; i < minLen; i++ {
    fmt.Println(slice1[offset1+i], slice2[offset2+i])
}
```

### 5. `if / else if / else` for Three Cases
You need to distinguish: slice2 is longer, slice1 is longer, or they are equal length.

```go
if len1 > len2 {
    // slice1 surplus goes first
} else if len2 > len1 {
    // slice2 surplus goes first
}
// equal length: no surplus, skip straight to interleave
```

## Review If Stuck
- Prior slice challenges — covers building slices dynamically with `append`
- Prior loop challenges — covers indexed `for i := 0; i < n; i++` loops

## You're Ready When You Can...
- [ ] Create an empty `[]int{}` and append values to it
- [ ] Use `len()` to compare two slice lengths
- [ ] Access a slice element by index: `s[i]`
- [ ] Write an indexed `for` loop with a computed offset
- [ ] Use `if / else if` to branch on which slice is longer

## Next Steps
- [52-concatslice](../52-concatslice/README.md)
