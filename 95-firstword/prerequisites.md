# Prerequisites for firstword

## Before You Start

To solve this challenge you need to understand:

### 1. strings.Fields
Splits a string by any whitespace and returns only the non-empty parts. Handles multiple spaces and leading/trailing spaces automatically.
```go
import "strings"

words := strings.Fields("hello   world")
// words == []string{"hello", "world"}
fmt.Println(len(words)) // 2

words2 := strings.Fields("   ")
// words2 == []string{}  (empty slice)
```

### 2. Slice Indexing
Access elements of a slice by position. Index 0 is the first element.
```go
words := []string{"hello", "world"}
fmt.Println(words[0]) // "hello"
fmt.Println(words[1]) // "world"
```

### 3. Checking Slice Length Before Indexing
Accessing an index on an empty slice causes a panic. Always check `len` first.
```go
if len(words) == 0 {
    return "\n"
}
fmt.Println(words[0]) // safe now
```

### 4. String Concatenation
Use `+` to join strings. Use `"\n"` to represent a newline character.
```go
word := "hello"
result := word + "\n" // "hello\n"
```

## Review If Stuck

- [../94-digitlen/skills.md](../94-digitlen/skills.md) — covers loops and counters; the slice-length check pattern is similar to guarding a counter

## You're Ready When You Can...

- [ ] Call `strings.Fields` on a string and know what it returns
- [ ] Access `words[0]` safely by checking `len(words) > 0` first
- [ ] Return a string that ends with `"\n"`
- [ ] Handle the empty-string input case (return just `"\n"`)

## Next Steps

- [96-gcd](../96-gcd/README.md)
