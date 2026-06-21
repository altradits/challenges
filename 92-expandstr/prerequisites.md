# Prerequisites for expandstr

## Before You Start

To solve this challenge you need to understand:

### 1. strings.Fields
Splits on any whitespace, returning only non-empty words.
```go
import "strings"

words := strings.Fields("   only  it's harder   ")
// words == []string{"only", "it's", "harder"}
```

### 2. strings.Join With a Separator
Joins a slice of strings, placing the separator *between* elements.
```go
import "strings"

words := []string{"only", "it's", "harder"}
result := strings.Join(words, "   ") // three spaces
// result == "only   it's   harder"
```

### 3. os.Args and Argument Count
Program name counts in `os.Args`. Exactly one user argument means `len(os.Args) == 2`.
```go
import "os"

if len(os.Args) != 2 {
    return  // print nothing
}
```

### 4. Printing Nothing (vs. Printing a Newline)
This challenge says display nothing on bad input. `return` exits without any output.
```go
if len(os.Args) != 2 {
    return  // truly no output, not even a newline
}
```

Compare with cleanstr, which requires `fmt.Println()` (a newline) on bad input.

## Review If Stuck

- [../91-cleanstr/skills.md](../91-cleanstr/skills.md) — covers the `strings.Fields` + `strings.Join` + `os.Args` pattern; this challenge is nearly identical but uses 3 spaces and no output on bad args

## You're Ready When You Can...

- [ ] Use `strings.Fields` to split a string into clean words
- [ ] Use `strings.Join(words, "   ")` with exactly three spaces
- [ ] Check `len(os.Args) != 2` and `return` with no output
- [ ] Distinguish "print nothing" (bare `return`) from "print newline" (`fmt.Println()`)

## Next Steps

- [93-findprevprime](../93-findprevprime/README.md)
