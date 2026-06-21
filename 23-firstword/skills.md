# Skills for firstword

## What You'll Learn

**Previous:** [../22-digitlen/skills.md](../22-digitlen/skills.md) | **Next:** [../24-recursion/README.md](../24-recursion/README.md)

**Challenge:** Return the first word in a string, followed by a newline.

## Core Concept: strings.Fields for Word Extraction

### What Is It?
`strings.Fields` splits a string on any whitespace (spaces, tabs, multiple consecutive spaces) and returns a slice of non-empty words. It is the cleanest way to extract words without manually managing space-counting logic.

```go
import "strings"

words := strings.Fields("  hello   world  ")
// words == []string{"hello", "world"}
fmt.Println(len(words)) // 2
```

### How It Works

```go
import "strings"

func FirstWord(s string) string {
    words := strings.Fields(s)
    if len(words) == 0 {
        return "\n"
    }
    return words[0] + "\n"
}
```

`words[0]` is the first element of the returned slice — always the first non-space word.

### Why Not strings.Split?
`strings.Split(s, " ")` splits on exactly one space. Multiple spaces create empty string elements in the slice, making it harder to work with:

```go
strings.Split("a  b", " ")  // ["a", "", "b"] — has an empty string
strings.Fields("a  b")       // ["a", "b"] — clean
```

### Manual Approach (Without strings.Fields)
Scan for the first non-space byte, then scan forward until the next space or end of string:

```go
func FirstWord(s string) string {
    start := -1
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' && start == -1 {
            start = i
        } else if s[i] == ' ' && start != -1 {
            return s[start:i] + "\n"
        }
    }
    if start != -1 {
        return s[start:] + "\n"
    }
    return "\n"
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Accessing `words[0]` without checking length | Panic: index out of range on empty slice | Always check `len(words) == 0` first |
| Forgetting `\n` | Output missing the required newline | Append `+ "\n"` to the return value |
| Using `strings.Split(s, " ")` | Fails on multiple spaces between words | Use `strings.Fields` instead |

## Solving This Challenge

### Algorithm
1. Call `strings.Fields(s)` to get a slice of words.
2. If the slice is empty, return `"\n"`.
3. Return `words[0] + "\n"`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../24-recursion/README.md](../24-recursion/README.md)
