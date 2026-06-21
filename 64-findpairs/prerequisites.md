# Prerequisites for findpairs

## Before You Start

To solve this challenge you need to understand:

### 1. `strings.Split` and `strings.TrimSpace`
Split a comma-separated string into parts; trim whitespace from each part.

```go
import "strings"

parts := strings.Split("1, 2, 3", ",")
// parts = ["1", " 2", " 3"]

for _, p := range parts {
    p = strings.TrimSpace(p)  // "1", "2", "3"
}
```

### 2. `strconv.Atoi` with Error Handling
Parse each number string into an `int`. If it fails, report the invalid value.

```go
import "strconv"

n, err := strconv.Atoi("42")
if err != nil {
    fmt.Println("invalid:", "42")
    return
}
```

### 3. String Indexing for Bracket Validation
Check the first and last characters of the array string.

```go
s := "[1, 2, 3]"
if s[0] != '[' || s[len(s)-1] != ']' {
    fmt.Println("Invalid input.")
    return
}
inner := s[1 : len(s)-1]  // "1, 2, 3"
```

### 4. Map with Slice Values (`map[int][]int`)
Store multiple indices per value to handle repeated numbers.

```go
seen := make(map[int][]int)
seen[3] = append(seen[3], 0)  // value 3 was at index 0
seen[3] = append(seen[3], 2)  // value 3 also at index 2
```

### 5. Printing Formatted Output
Match the expected output format exactly using `fmt.Printf`.

```go
fmt.Printf("Pairs with sum %d: [[%d %d] [%d %d]]\n", target, i1, j1, i2, j2)
```

## Review If Stuck
- [55-inter](../55-inter/skills.md) — covers using maps for set operations
- [53-fprime](../53-fprime/skills.md) — covers `strconv.Atoi` and `os.Args` validation
- [63-slice](../63-slice/skills.md) — covers string slicing with `s[1:len-1]`

## You're Ready When You Can...
- [ ] Use `strings.Split` to break a string by a delimiter
- [ ] Use `strings.TrimSpace` to remove whitespace from a string
- [ ] Check `s[0]` and `s[len(s)-1]` to validate bracket format
- [ ] Use `strconv.Atoi` and print an error message on failure
- [ ] Use `map[int][]int` to store multiple indices per value
- [ ] Find pairs by looking up `target - v` in the map for each element

## Next Steps
- [65-revwstr](../65-revwstr/README.md)
