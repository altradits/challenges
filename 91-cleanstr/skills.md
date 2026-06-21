# Skills for cleanstr

## What You'll Learn

**Previous:** [90-clean-the-list](../90-clean-the-list/skills.md) | **Next:** [92-expandstr](../92-expandstr/skills.md)

**Challenge:** A program that takes exactly one argument and prints its words separated by exactly one space, with no leading or trailing spaces.

## Core Concept: Normalizing Whitespace With strings.Fields and strings.Join

### What Is It?
`strings.Fields` splits on any amount of whitespace (including tabs) and returns only non-empty words. Joining the result with a single space gives the normalized output in one step.

```go
import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println()
        return
    }
    words := strings.Fields(os.Args[1])
    if len(words) == 0 {
        fmt.Println()
        return
    }
    fmt.Println(strings.Join(words, " "))
}
```

### Why strings.Fields + strings.Join?
- `strings.Fields("  only    it's  harder   ")` → `["only", "it's", "harder"]`
- `strings.Join(["only", "it's", "harder"], " ")` → `"only it's harder"`

The combination handles all edge cases: multiple spaces, tabs, leading/trailing whitespace.

### strings.Join
Joins a slice of strings with a separator between each element (not before the first or after the last):
```go
words := []string{"one", "two", "three"}
fmt.Println(strings.Join(words, " "))   // "one two three"
fmt.Println(strings.Join(words, ", "))  // "one, two, three"
fmt.Println(strings.Join(words, ""))    // "onetwothree"
```

### Argument Count Check
The challenge says: if the number of arguments is not 1, or if there are no words, print a newline.
```go
if len(os.Args) != 2 {  // program name + exactly 1 arg = 2
    fmt.Println()
    return
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Checking `len(os.Args) != 1` | Off by one — program name counts | Check `!= 2` |
| Using `strings.Split(s, " ")` | Doesn't collapse multiple spaces | Use `strings.Fields` |
| Not printing a newline when args are wrong | Challenge requires a newline | Use `fmt.Println()` (not bare `return`) |

## Solving This Challenge

### Algorithm
1. If `len(os.Args) != 2`, print newline and return.
2. Get `s := os.Args[1]`.
3. `words := strings.Fields(s)`.
4. If `len(words) == 0`, print newline and return.
5. Print `strings.Join(words, " ")`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [92-expandstr](../92-expandstr/README.md)
