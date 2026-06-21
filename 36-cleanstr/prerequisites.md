# Prerequisites for cleanstr

## Before You Start

To solve this challenge you need to understand:

### 1. strings.Fields
Splits a string on any whitespace, collapsing multiple spaces and ignoring leading/trailing whitespace. Returns only non-empty strings.
```go
import "strings"

words := strings.Fields("  only    it's  harder   ")
// words == []string{"only", "it's", "harder"}
```

### 2. strings.Join
Joins a slice of strings with a separator placed *between* elements.
```go
import "strings"

words := []string{"only", "it's", "harder"}
result := strings.Join(words, " ")
// result == "only it's harder"
```

### 3. os.Args and Argument Count Validation
`os.Args[0]` is the program name. User arguments start at index 1. Total count includes the program name.
```go
import "os"

// Running: go run . "hello world"
// len(os.Args) == 2  (program + 1 arg)
if len(os.Args) != 2 {
    fmt.Println()
    return
}
```

### 4. Printing an Empty Newline
When there is no output content, the challenge still requires a newline. `fmt.Println()` with no arguments prints just a newline.
```go
fmt.Println() // prints "\n"
```

## Review If Stuck

- [../35-clean-the-list/skills.md](../35-clean-the-list/skills.md) — covers `strings.TrimSpace` and working with word-by-word processing
- [../23-firstword/skills.md](../23-firstword/skills.md) — introduces `strings.Fields`
- [../31-splitjoin/skills.md](../31-splitjoin/skills.md) — shows how Join works internally

## You're Ready When You Can...

- [ ] Use `strings.Fields` to split a string into clean words regardless of spacing
- [ ] Use `strings.Join` to reassemble words with exactly one space between them
- [ ] Check `len(os.Args) != 2` to validate single-argument programs
- [ ] Print a bare newline with `fmt.Println()` for invalid/empty input

## Next Steps

- [Next challenge](../37-expandstr/README.md)
