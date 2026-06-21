# Prerequisites for chunk

## Before You Start

To solve this challenge you need to understand:

### 1. Sub-Slicing With slice[i:j]
Extract a portion of a slice from index `i` to `j-1` (j is exclusive).
```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7}
fmt.Println(s[0:3]) // [0 1 2]
fmt.Println(s[3:6]) // [3 4 5]
fmt.Println(s[6:8]) // [6 7]
```

### 2. Advancing a Loop by a Variable Step
Use `i += size` to step through a slice in chunks:
```go
size := 3
for i := 0; i < len(s); i += size {
    fmt.Println(i) // 0, 3, 6
}
```

### 3. Clamping the End Index
When computing the end of a chunk, ensure it doesn't exceed the slice length.
```go
end := i + size
if end > len(slice) {
    end = len(slice)  // clamp to valid range
}
chunk := slice[i:end]
```

### 4. Slice of Slices [][]int
A slice where each element is itself an `[]int` slice. Append sub-slices to build the result.
```go
result := [][]int{}
result = append(result, []int{0, 1, 2})
result = append(result, []int{3, 4, 5})
fmt.Println(result) // [[0 1 2] [3 4 5]]
```

### 5. fmt.Println on a Slice of Slices
`fmt.Println` knows how to print nested slices in the `[[...] [...]]` format automatically.
```go
chunks := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7}}
fmt.Println(chunks) // [[0 1 2] [3 4 5] [6 7]]
```

## Review If Stuck

- [../102-splitjoin/skills.md](../102-splitjoin/skills.md) — covers building a `[]string` slice with `append`; the same pattern applies to `[][]int`
- [../118-canjump/skills.md](../118-canjump/skills.md) — covers index arithmetic and bounds checking on slices

## You're Ready When You Can...

- [ ] Create a sub-slice with `s[i:j]` without exceeding the slice length
- [ ] Write a loop that advances by `size` using `i += size`
- [ ] Clamp the end index with `if end > len(slice) { end = len(slice) }`
- [ ] Append sub-slices to a `[][]int` result and print it

## Next Steps

- [120-concatalternate](../120-concatalternate/README.md)
