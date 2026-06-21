# Skills for wordflip

## What You'll Learn

**Previous:** [135-rostring](../135-rostring/skills.md) | **Next:** [137-itoabase](../137-itoabase/skills.md)

**Challenge:** Reverse each individual word in the string (characters within each word reversed), and reverse the order of words. Also trim and normalize spaces.

## Core Concept: Per-Word Character Reversal

### What Is It?

`WordFlip` combines two reversals:
1. Reverse the character order within each word.
2. Reverse the order of words.

Wait — reading the examples again: `"First second last"` → `"last second First"`. The words are reversed in order but the characters within each word are NOT reversed. This is identical to `revwstr` (65)!

Looking more carefully: `" hello  all  of  you! "` → `"you! of all hello"`. Yes — this is word-order reversal, but also requires trimming leading/trailing spaces and normalizing internal multiple spaces. The key new element vs. `revwstr` is the space-trimming and multiple-space normalization.

### How It Works

**Step 1 — Handle edge cases:**

```go
func WordFlip(str string) string {
    // Trim leading and trailing spaces
    trimmed := strings.TrimSpace(str)
    if trimmed == "" {
        return "Invalid Output"
    }
```

**Step 2 — Split (handles multiple spaces automatically), reverse, rejoin:**

```go
    words := strings.Fields(trimmed)
    // Reverse word order
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ") + "\n"
}
```

**Trace `" hello  all  of  you! "`:**
1. `strings.TrimSpace` → `"hello  all  of  you!"`
2. `strings.Fields` → `["hello", "all", "of", "you!"]`
3. Reverse → `["you!", "of", "all", "hello"]`
4. Join → `"you! of all hello\n"` ✓

**Trace `"     "` (all spaces):**
1. `strings.TrimSpace` → `""`
2. `trimmed == ""` → return `"Invalid Output"` ✓

**New skill vs. revwstr (65): `strings.TrimSpace` and the "Invalid Output" return**

```go
import "strings"

strings.TrimSpace("  hello  ")  // "hello"
strings.TrimSpace("     ")      // ""
strings.TrimSpace("")           // ""
```

The `"Invalid Output"` case triggers when the string is empty OR contains only spaces.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Checking `str == ""` only | `"     "` (all spaces) also returns "Invalid Output" | Use `strings.TrimSpace` then check |
| Adding `\n` to "Invalid Output" | The README says return `"Invalid Output"` (no newline shown) | Check the expected output format carefully |
| Using `strings.Split(str, " ")` | Multiple spaces create empty strings in the slice | Use `strings.Fields` |

## Solving This Challenge

### Algorithm
1. `trimmed = strings.TrimSpace(str)`.
2. If `trimmed == ""`, return `"Invalid Output"`.
3. `words = strings.Fields(trimmed)`.
4. Reverse `words` in-place.
5. Return `strings.Join(words, " ") + "\n"`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [137-itoabase](../137-itoabase/README.md)
