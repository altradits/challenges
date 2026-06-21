# Prerequisites for 41-join

## Before You Start

### 1. Slices and `len`

A `[]string` slice holds an ordered list of strings. `len(slice)` returns how many elements it has.

```go
words := []string{"hello", "world", "go"}
fmt.Println(len(words))   // 3
fmt.Println(words[0])     // "hello"
fmt.Println(words[2])     // "go"
```

### 2. Iterating a Slice by Index

When you need the index to start from 1, use a traditional `for` loop (not `range`):

```go
words := []string{"a", "b", "c"}
for i := 1; i < len(words); i++ {
    fmt.Println(words[i])  // prints "b", then "c"
}
```

### 3. String Concatenation with `+=`

Build a result string by appending:

```go
result := "start"
result += "-middle"
result += "-end"
fmt.Println(result)  // "start-middle-end"
```

### 4. Accessing the First Element Safely

Before accessing `slice[0]`, always verify the slice is not empty:

```go
if len(elems) == 0 {
    return ""
}
first := elems[0]  // safe: we know len >= 1
```

### 5. Relationship to Split (from 40-split)

Join is the reverse operation of Split. If you split `"a,b,c"` by `","`, you get `["a","b","c"]`. If you join `["a","b","c"]` with `","`, you get `"a,b,c"` again.

```go
parts := []string{"a", "b", "c"}
sep := ","
// Join should produce: "a,b,c"
```

## Review If Stuck

- [40-split](../40-split/skills.md) — covers the inverse operation and slice fundamentals

## You're Ready When You Can...

- [ ] Access elements of a `[]string` by index
- [ ] Check if a slice is empty with `len(slice) == 0`
- [ ] Write a loop that starts at index 1 instead of 0
- [ ] Explain why the separator goes between elements, not after them

## Next Steps

- [42-cameltosnakecase](../42-cameltosnakecase/README.md)
