# Prerequisites for 97-longestword

## Before You Start

To solve this challenge you need to understand:

### 1. `strings.Fields` — splitting on whitespace

`strings.Fields` splits a string into words, treating any whitespace as a separator.

```go
import "strings"

words := strings.Fields("  hello   world  ")
// words = ["hello", "world"]
```

Review: [90-wordcount](../90-wordcount/skills.md)

### 2. `len()` on strings

`len(s)` returns the byte count of a string (equals character count for ASCII).

```go
len("hello") // 5
len("")       // 0
```

Review: [76-stringlength](../76-stringlength/skills.md)

### 3. Tracking a running maximum

```go
best := ""
for _, item := range items {
    if len(item) > len(best) {
        best = item
    }
}
```

## You're Ready When You Can...

- [ ] Split a string into words
- [ ] Compare lengths of two strings with `len()`
- [ ] Write a loop that keeps track of the best value seen so far

## Next Steps

- [98-searchreplace](../98-searchreplace/README.md) — next challenge
