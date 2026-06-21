# Prerequisites for longestword

## Before You Start

To solve this challenge you need to understand:

### 1. strings.Fields
Splits a string on any whitespace and returns a slice of non-empty words. Handles multiple spaces and leading/trailing spaces.
```go
import "strings"

words := strings.Fields("the quick brown fox")
// words == []string{"the", "quick", "brown", "fox"}
```

### 2. for/range Over a Slice
Iterate over every element of a slice using `range`. The blank identifier `_` discards the index.
```go
words := []string{"the", "quick", "fox"}
for _, word := range words {
    fmt.Println(word, len(word))
}
```

### 3. len() on a String
`len(s)` returns the number of bytes in a string — for ASCII words, this equals the number of characters.
```go
fmt.Println(len("quick")) // 5
fmt.Println(len("fox"))   // 3
```

### 4. Tracking a Running Maximum
Declare a variable to hold the best candidate so far. Update it whenever you find something better.
```go
best := ""
for _, word := range words {
    if len(word) > len(best) {
        best = word
    }
}
```

### 5. Empty Slice Guard
`strings.Fields` returns an empty slice for all-whitespace or empty input. Check before using the result.
```go
if len(words) == 0 {
    return ""
}
```

## Review If Stuck

- [../98-lastword/skills.md](../98-lastword/skills.md) — covers `strings.Fields` and safe slice access patterns
- [../95-firstword/skills.md](../95-firstword/skills.md) — covers `strings.Fields` and the empty slice guard

## You're Ready When You Can...

- [ ] Split a string into words with `strings.Fields`
- [ ] Iterate over a slice with `for _, word := range words`
- [ ] Compare `len(word)` values to find the longest
- [ ] Use strict `>` to keep the first word on a tie
- [ ] Return `""` for empty or all-whitespace input

## Next Steps

- [100-replaceall](../100-replaceall/README.md)
