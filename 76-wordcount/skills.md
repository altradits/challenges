# Skills for 76-wordcount

## What You'll Learn

**Previous:** [75-titlecase](../75-titlecase/skills.md) | **Next:** [77-cameltosnakecase](../77-cameltosnakecase/skills.md)

**Challenge:** Write `WordCount(s string) int` that returns the number of words in a string, where words are separated by one or more spaces.

## Core Concept: Counting Words with `strings.Fields`

### What Is It?

Counting words sounds like it requires a complex loop, but `strings.Fields` makes it almost trivial: split the string on whitespace and count the resulting slice. The key insight is that `strings.Fields` already handles all the edge cases — multiple spaces, leading/trailing spaces, and empty strings.

### The One-Liner Solution

```go
import "strings"

func WordCount(s string) int {
    return len(strings.Fields(s))
}
```

`strings.Fields(s)` returns a `[]string` containing the words, and `len()` on a slice counts the elements.

```
""                -> Fields -> []              -> len = 0
"   "             -> Fields -> []              -> len = 0
"Hello World"     -> Fields -> ["Hello","World"] -> len = 2
"One  two   three"-> Fields -> ["One","two","three"] -> len = 3
```

All the edge cases from the README are handled automatically.

### How `strings.Fields` Works Internally

`strings.Fields` splits on any Unicode whitespace (spaces, tabs, newlines) and ignores empty segments. It is equivalent to splitting on a regular expression `\s+` and removing empty results.

```go
strings.Fields("  hello   world  ")
// Step 1: Split on whitespace -> ["", "", "hello", "", "", "world", "", ""]
// Step 2: Remove empty strings -> ["hello", "world"]
// Returns: ["hello", "world"]
```

### `len()` on a Slice

`len()` works on strings AND slices:

```go
len("hello")                   // 5  — byte count of string
len([]string{"a", "b", "c"})  // 3  — element count of slice
```

### Manual Implementation (No Standard Library)

If you want to understand the logic from scratch:

```go
func WordCount(s string) int {
    count := 0
    inWord := false
    for _, r := range s {
        if r != ' ' {
            if !inWord {
                count++
                inWord = true
            }
        } else {
            inWord = false
        }
    }
    return count
}
```

This uses a boolean `inWord` flag — `true` when you are currently inside a word, `false` when in whitespace. You increment the counter only when you *enter* a new word (transition from whitespace to non-whitespace).

### Diagram: Manual Count of "One  two   three"

```
O  n  e     t  w  o        t  h  r  e  e
|  |  |  |  |  |  |  |  |  |  |  |  |  |

inWord: F
'O': not space, inWord=false -> count=1, inWord=true
'n': not space, inWord=true  -> no count change
'e': not space, inWord=true  -> no count change
' ': space -> inWord=false
' ': space -> inWord=false
't': not space, inWord=false -> count=2, inWord=true
'w': not space, inWord=true  -> ...
...
count = 3
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Splitting with `strings.Split(s, " ")` | Does not handle multiple spaces; `"one  two"` gives `["one", "", "two"]` | Use `strings.Fields` |
| `len(strings.Split(s, " "))` | Counts empty strings from multiple spaces | Use `strings.Fields` instead |
| Counting spaces instead of words | Off by one or wrong for leading/trailing spaces | Count word *transitions*, not spaces |

## Solving This Challenge

### Algorithm (with standard library)

1. Call `strings.Fields(s)` to split into words
2. Return `len()` of the resulting slice

### Algorithm (manual)

1. Initialize `count := 0` and `inWord := false`
2. For each rune: if non-space and not in a word, increment count and set `inWord = true`; if space, set `inWord = false`
3. Return `count`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [77-cameltosnakecase](../77-cameltosnakecase/README.md)
