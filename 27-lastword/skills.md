# Skills for lastword

## What You'll Learn

**Previous:** [../26-hashcode/skills.md](../26-hashcode/skills.md) | **Next:** [../28-longestword/README.md](../28-longestword/README.md)

**Challenge:** Return the last word in a string, followed by a newline.

## Core Concept: Accessing the Last Element of a Slice

### What Is It?
Once you split a string into words with `strings.Fields`, the last word is always at index `len(words)-1`. This off-by-one pattern is how you access any last element in Go — there is no built-in `last()` function.

```go
words := []string{"one", "two", "three"}
fmt.Println(words[len(words)-1]) // "three"
```

### How It Works

```go
import "strings"

func LastWord(s string) string {
    words := strings.Fields(s)
    if len(words) == 0 {
        return "\n"
    }
    return words[len(words)-1] + "\n"
}
```

`strings.Fields` handles leading/trailing spaces and multiple consecutive spaces, so `LastWord(" lorem,ipsum ")` correctly returns `"lorem,ipsum\n"`.

### Manual Backward Scan (Without strings.Fields)
Scan from the end of the string to skip trailing spaces, then find the start of the last word:

```go
func LastWord(s string) string {
    // Skip trailing spaces
    end := len(s) - 1
    for end >= 0 && s[end] == ' ' {
        end--
    }
    if end < 0 {
        return "\n"
    }
    // Find start of last word
    start := end
    for start > 0 && s[start-1] != ' ' {
        start--
    }
    return s[start:end+1] + "\n"
}
```

### The `len(slice)-1` Last-Element Pattern
This is the standard Go idiom for the last element:
```go
words := strings.Fields(s)
last := words[len(words)-1]  // correct
// NOT: words[len(words)]    // panic — one past the end
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `words[len(words)]` | Index out of range panic (off by one) | Use `words[len(words)-1]` |
| Not checking for empty slice | Panic when string has no words | Guard with `if len(words) == 0` |
| Forgetting `\n` | Missing required newline | Append `+ "\n"` |

## Solving This Challenge

### Algorithm
1. Split with `strings.Fields(s)`.
2. If slice is empty, return `"\n"`.
3. Return `words[len(words)-1] + "\n"`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../28-longestword/README.md](../28-longestword/README.md)
