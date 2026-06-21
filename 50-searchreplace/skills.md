# Skills for 50-searchreplace

## What You'll Learn

**Previous:** [49-longestword](../49-longestword/skills.md) | **Next:** [51-cleanlist](../51-cleanlist/skills.md)

**Challenge:** Find the first occurrence of `old` in `text` and replace it with `new`.

## Core Concept: Substring Replacement via String Slicing

### What Is It?

Replacing a substring means splitting the original string into three parts — before, the match, and after — then joining them with the replacement in the middle.

### How It Works

```
text = "Hello World"
old  = "World"
new  = "Go"

Step 1: Find index of "World" → 6
Step 2: before = text[:6]         → "Hello "
Step 3: after  = text[6+5:]       → ""   (6 + len("World"))
Step 4: return before + new + after → "Hello Go"
```

### The Manual Implementation

```go
func SearchReplace(text, old, newStr string) string {
    // Find position of old in text
    idx := -1
    for i := 0; i <= len(text)-len(old); i++ {
        if text[i:i+len(old)] == old {
            idx = i
            break
        }
    }
    if idx == -1 {
        return text  // not found, return unchanged
    }
    return text[:idx] + newStr + text[idx+len(old):]
}
```

### Using `strings.Index`

```go
import "strings"

func SearchReplace(text, old, newStr string) string {
    idx := strings.Index(text, old)
    if idx == -1 {
        return text
    }
    return text[:idx] + newStr + text[idx+len(old):]
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `strings.ReplaceAll` | Replaces ALL occurrences | Use `strings.Index` + slicing for first-only |
| Forgetting the `-1` check | Panics if old not found | Always guard with `if idx == -1` |
| `text[idx+len(new):]` | Uses length of `new`, not `old` | The skip is `len(old)` — the piece you removed |

## Algorithm

1. Find index of `old` in `text` (return -1 if absent)
2. If not found, return `text` unchanged
3. Return `text[:idx]` + `newStr` + `text[idx+len(old):]`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [51-cleanlist](../51-cleanlist/README.md)
