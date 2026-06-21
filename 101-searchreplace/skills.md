# Skills for searchreplace

## What You'll Learn

**Previous:** [100-replaceall](../100-replaceall/skills.md) | **Next:** [102-splitjoin](../102-splitjoin/skills.md)

**Challenge:** A command-line program that takes exactly 3 arguments and replaces every occurrence of a single character with another.

## Core Concept: Reading and Validating os.Args

### What Is It?
`os.Args` is a slice of strings containing the command-line arguments. `os.Args[0]` is the program name; the actual arguments start at index 1. Validating the count before using them is the first thing any CLI program should do.

```go
import "os"

func main() {
    // os.Args[0] is the program name
    // os.Args[1], os.Args[2], ... are the user's arguments
    if len(os.Args) != 4 {  // program + 3 args = 4 total
        return
    }
    text  := os.Args[1]  // the string
    old   := os.Args[2]  // character to find
    newCh := os.Args[3]  // character to replace with
}
```

### Replacing a Single Character
Since the second and third arguments are single characters (strings of length 1), you can use `strings.Contains` and then loop to replace:

```go
import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 4 {
        return
    }
    text  := os.Args[1]
    old   := os.Args[2]
    newCh := os.Args[3]

    result := ""
    for _, c := range text {
        if string(c) == old {
            result += newCh
        } else {
            result += string(c)
        }
    }
    fmt.Println(result)
}
```

Or use `strings.ReplaceAll` since this challenge does NOT forbid it:
```go
fmt.Println(strings.ReplaceAll(text, old, newCh))
```

### When to Print Nothing
The challenge says: if the argument count is wrong, print nothing. A bare `return` in `main()` exits without printing:
```go
if len(os.Args) != 4 {
    return  // exits main, no output
}
```

### When the Search Character Isn't Found
If `old` is not in `text`, just print the text unchanged. Your replacement loop handles this naturally — no matches means the original text is copied unchanged.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Checking `len(os.Args) != 3` | Off by one — program name counts | Check `!= 4` (program + 3 args) |
| Using `os.Args[1]` without checking length first | Panic if no arguments given | Always check `len(os.Args)` before indexing |
| Printing a newline when args are wrong | Challenge says display nothing | Use bare `return`, not `fmt.Println()` |

## Solving This Challenge

### Algorithm
1. If `len(os.Args) != 4`, return (print nothing).
2. Get `text = os.Args[1]`, `old = os.Args[2]`, `newCh = os.Args[3]`.
3. Replace every occurrence of `old` in `text` with `newCh`.
4. Print the result with `fmt.Println`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [102-splitjoin](../102-splitjoin/README.md)
