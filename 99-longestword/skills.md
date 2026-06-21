# Skills for longestword

## What You'll Learn

**Previous:** [98-lastword](../98-lastword/skills.md) | **Next:** [100-replaceall](../100-replaceall/skills.md)

**Challenge:** Find the longest word in a string, returning the first one if there is a tie.

## Core Concept: Tracking a Running Maximum

### What Is It?
Iterate over all words and keep track of the longest one seen so far. This "running maximum" pattern appears in many problems: scan all items, update your best candidate whenever you find a better one.

```go
import "strings"

func LongestWord(s string) string {
    words := strings.Fields(s)
    if len(words) == 0 {
        return ""
    }
    longest := ""
    for _, word := range words {
        if len(word) > len(longest) {
            longest = word
        }
    }
    return longest
}
```

### Step-by-Step for `"the quick brown fox"`

Words: `["the", "quick", "brown", "fox"]`

| Iteration | word    | len(word) | len(longest) | Update? |
|-----------|---------|-----------|--------------|---------|
| 1         | "the"   | 3         | 0 ("")       | Yes → longest="the" |
| 2         | "quick" | 5         | 3            | Yes → longest="quick" |
| 3         | "brown" | 5         | 5            | No (tie — keep first) |
| 4         | "fox"   | 3         | 5            | No |

Result: `"quick"`

### Why Initialize `longest` as `""`?
Starting with an empty string means `len("") == 0`, so the very first word always becomes the initial longest candidate. This is cleaner than using a separate index variable.

### Strict Greater-Than for Ties
Using `>` (not `>=`) means you only update when you find something *strictly longer*, so ties keep the *first* word — matching the challenge requirement.

```go
if len(word) > len(longest) {  // > keeps first on tie
    longest = word
}
// vs. >= would keep the last on tie
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `>=` instead of `>` | On a tie, keeps the *last* word instead of the first | Use `>` strictly |
| Not checking for empty input | `strings.Fields("")` returns empty slice; returning `longest` = `""` is actually correct | Still guard or test carefully |
| Comparing `len(s)` directly instead of per-word | Measures the whole string | Use `strings.Fields` and compare each `word` |

## Solving This Challenge

### Algorithm
1. Split with `strings.Fields(s)`. If empty, return `""`.
2. Initialize `longest := ""`.
3. For each word: if `len(word) > len(longest)`, update `longest = word`.
4. Return `longest`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [100-replaceall](../100-replaceall/README.md)
