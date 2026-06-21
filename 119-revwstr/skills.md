# Skills for revwstr

## What You'll Learn

**Previous:** [118-findpairs](../118-findpairs/skills.md) | **Next:** [120-rostring](../120-rostring/skills.md)

**Challenge:** Reverse the order of words in a string. Words are alphanumeric sequences separated by exactly one space.

## Core Concept: Splitting, Reversing, and Rejoining Words

### What Is It?

Reversing the words in a string means taking `"the time of contempt"` and outputting `"contempt of time the"`. The characters within each word stay in order — only the word positions are reversed.

### How It Works

**Step 1 — Split the string into words:**

`strings.Fields` splits on any whitespace (spaces, tabs) and ignores leading/trailing spaces — perfect for this challenge.

```go
import "strings"

words := strings.Fields("the time of contempt")
// words = ["the", "time", "of", "contempt"]
```

**Step 2 — Reverse the slice of words:**

```go
for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
    words[i], words[j] = words[j], words[i]
}
// words = ["contempt", "of", "time", "the"]
```

**Step 3 — Join with a single space and print:**

```go
fmt.Println(strings.Join(words, " "))
// "contempt of time the"
```

**Full program:**

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    s := os.Args[1]
    if s == "" {
        fmt.Println()
        return
    }
    words := strings.Fields(s)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    fmt.Println(strings.Join(words, " "))
}
```

**`strings.Fields` vs `strings.Split`:**

| Function | `"  a  b  c  "` |
|----------|-----------------|
| `strings.Split(s, " ")` | `["", "", "a", "", "b", "", "c", "", ""]` — includes empty strings |
| `strings.Fields(s)` | `["a", "b", "c"]` — handles any whitespace, ignores leading/trailing |

Use `strings.Fields` when you want "clean" word splitting.

**Two-pointer swap:**

The `i, j = i+1, j-1` idiom swaps from both ends simultaneously until they meet in the middle:

```go
i=0,j=3 → swap words[0] and words[3]
i=1,j=2 → swap words[1] and words[2]
i=2,j=1 → i >= j, stop
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `strings.Split(s, " ")` | Empty strings in the slice when multiple spaces present | Use `strings.Fields` |
| Building the reverse manually with `append` | More complex, error-prone | Use in-place two-pointer swap |
| Printing a newline for empty input | `fmt.Println()` on empty string | Check `s == ""` and print bare newline |

## Solving This Challenge

### Algorithm
1. Validate exactly 1 argument; return otherwise.
2. If empty string, print newline and return.
3. `words = strings.Fields(s)`.
4. Reverse `words` in-place with two-pointer swap.
5. `fmt.Println(strings.Join(words, " "))`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [120-rostring](../120-rostring/README.md)
