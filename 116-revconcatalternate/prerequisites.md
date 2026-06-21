# Prerequisites for revconcatalternate

## Before You Start

To solve this challenge you need to understand:

### 1. Iterating a Slice in Reverse
Use a `for` loop counting down from `len-1` to `0`.

```go
s := []int{1, 2, 3, 4, 5}
for i := len(s) - 1; i >= 0; i-- {
    fmt.Println(s[i])  // 5, 4, 3, 2, 1
}
```

### 2. Minimum of Two Lengths
Compute `minLen` to know the overlap portion.

```go
len1, len2 := len(slice1), len(slice2)
minLen := len1
if len2 < minLen {
    minLen = len2
}
```

### 3. Computing the Surplus Range
The surplus elements of the longer slice are the ones beyond index `minLen`, i.e., from index `minLen` to `len-1`.

```go
// Surplus of slice1 when len1 > len2, in reverse:
for i := len1 - 1; i >= minLen; i-- {
    result = append(result, slice1[i])
}
```

### 4. `append` to Build the Result
Always append to a `[]int{}` result slice.

```go
result := []int{}
result = append(result, 42)
```

### 5. `ConcatAlternate` Pattern (Challenge 51)
This challenge is the reverse-traversal version of `ConcatAlternate`. Review that solution first.

## Review If Stuck
- [105-concatalternate](../105-concatalternate/skills.md) — teaches the alternating merge pattern this challenge reverses
- [106-concatslice](../106-concatslice/skills.md) — covers `append` and slice building

## You're Ready When You Can...
- [ ] Iterate a slice from last index down to 0
- [ ] Compute `minLen = min(len1, len2)`
- [ ] Identify the surplus elements of the longer slice (indices `minLen` to `len-1`)
- [ ] Use `append` to build a result slice element by element
- [ ] Interleave two slices in reverse by decrementing a shared index

## Next Steps
- [117-slice](../117-slice/README.md)
