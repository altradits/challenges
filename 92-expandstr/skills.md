# Skills for expandstr

## What You'll Learn

**Previous:** [91-cleanstr](../91-cleanstr/skills.md) | **Next:** [93-findprevprime](../93-findprevprime/skills.md)

**Challenge:** A program that takes one argument and prints words separated by exactly three spaces, with no leading or trailing spaces.

## Core Concept: strings.Join With a Custom Separator

### What Is It?
The difference from cleanstr is the separator: three spaces instead of one. The same `strings.Fields` + `strings.Join` pattern works, but you pass `"   "` (three spaces) to `strings.Join`.

```go
import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        return  // display nothing
    }
    words := strings.Fields(os.Args[1])
    if len(words) == 0 {
        return  // display nothing
    }
    fmt.Println(strings.Join(words, "   "))
}
```

### Key Difference From cleanstr
- cleanstr: `strings.Join(words, " ")` — one space
- expandstr: `strings.Join(words, "   ")` — three spaces
- cleanstr prints a newline on bad input; expandstr prints nothing

```console
$ go run . "  only  it's harder   " | cat -e
only   it's   harder$
```

### When to Print Nothing
cleanstr requires a newline on invalid input. expandstr requires *nothing*:
```go
if len(os.Args) != 2 {
    return  // no fmt.Println() — truly no output
}
```

### Checking for No Words
If the argument is all spaces (`go run . "   "`), `strings.Fields` returns an empty slice. The challenge says display nothing in this case:
```go
words := strings.Fields(os.Args[1])
if len(words) == 0 {
    return
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using one space in `strings.Join` | Output has single spaces instead of triple | Count your spaces: `"   "` (three) |
| Printing a newline on bad input | Challenge says print nothing | Use bare `return`, not `fmt.Println()` |
| Not trimming words before joining | `strings.Fields` already does this | No extra work needed — `strings.Fields` handles it |

## Solving This Challenge

### Algorithm
1. If `len(os.Args) != 2`, return (no output).
2. `words := strings.Fields(os.Args[1])`.
3. If `len(words) == 0`, return (no output).
4. Print `strings.Join(words, "   ")`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [93-findprevprime](../93-findprevprime/README.md)
