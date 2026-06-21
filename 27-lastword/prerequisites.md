# Prerequisites for lastword

## Before You Start

To solve this challenge you need to understand:

### 1. strings.Fields
Splits a string on any whitespace, ignoring leading/trailing spaces and runs of multiple spaces. Returns a slice of words.
```go
import "strings"

words := strings.Fields("  hello   world  ")
// words == []string{"hello", "world"}

words2 := strings.Fields("   ")
// words2 == []string{}  (empty slice)
```

### 2. Slice Length and the Last-Element Pattern
`len(slice)` gives the count of elements. Last element is always at `len(slice)-1`.
```go
words := []string{"one", "two", "three"}
fmt.Println(len(words))             // 3
fmt.Println(words[len(words)-1])    // "three"
```

### 3. Empty Slice Guard
Never index an empty slice — it panics. Check length before accessing.
```go
words := strings.Fields("   ")
if len(words) == 0 {
    return "\n"
}
// safe to use words[len(words)-1] now
```

### 4. Trailing Newline in Return Value
The challenge requires `\n` at the end of the string you return.
```go
return words[len(words)-1] + "\n"
```

## Review If Stuck

- [../23-firstword/skills.md](../23-firstword/skills.md) — covers `strings.Fields` and safe slice indexing for `words[0]`; the same pattern applies to the last element

## You're Ready When You Can...

- [ ] Call `strings.Fields` and understand the slice it returns
- [ ] Access the last element with `words[len(words)-1]`
- [ ] Check for an empty slice before any indexing
- [ ] Return a string with a trailing newline

## Next Steps

- [Next challenge](../28-longestword/README.md)
