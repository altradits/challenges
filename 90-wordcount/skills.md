# Skills for 90-wordcount

## What You'll Learn

**Previous:** [89-retainfirsthalf](../89-retainfirsthalf/skills.md) | **Next:** [91-findchar](../91-findchar/skills.md)

**Challenge:** Write a function `WordCount(s string) int` that returns the number of words in a string (words are sequences of non-space characters).

## Core Concept: State Tracking for Word Boundaries

### What Is a Word Boundary?

A word starts when you transition from a space (or the beginning of the string) to a non-space character. It ends when you hit a space. The challenge is to count these **transitions**, not individual characters.

### Why a Simple Counter Does Not Work

If you just count spaces, `"a  b"` (two spaces) gives 2, suggesting 3 words — but there are only 2. The double space should not create an extra word.

### The `inWord` Flag Pattern

Track whether the current position is inside a word:

```go
func WordCount(s string) int {
    count := 0
    inWord := false
    for _, c := range s {
        if c != ' ' {
            if !inWord {
                count++       // found the start of a new word
                inWord = true
            }
            // already in a word: nothing to do
        } else {
            inWord = false    // space: we left a word (or are between words)
        }
    }
    return count
}
```

Step by step:
1. `inWord = false` — not inside a word yet
2. For each character:
   - Non-space AND not yet in a word → new word starts, `count++`, `inWord = true`
   - Non-space AND already in a word → do nothing, still in the same word
   - Space → `inWord = false`, we are between words (or before the first word)
3. Return `count`

### Tracing `"One  two   three"`

| c | non-space? | inWord before | Action | count |
|---|-----------|--------------|--------|-------|
| O | yes | false | count++, inWord=true | 1 |
| n | yes | true | — | 1 |
| e | yes | true | — | 1 |
| ` ` | no | true | inWord=false | 1 |
| ` ` | no | false | inWord=false | 1 |
| t | yes | false | count++, inWord=true | 2 |
| w | yes | true | — | 2 |
| o | yes | true | — | 2 |
| ` ` | no | true | inWord=false | 2 |
| ` ` | no | false | — | 2 |
| ` ` | no | false | — | 2 |
| t | yes | false | count++, inWord=true | 3 |
| ... | | | | 3 |

Result: `3` — correct despite multiple spaces.

### Handling Edge Cases Automatically

- Empty string: loop runs zero times, returns `0`
- Only spaces: `inWord` is never set to `true`, returns `0`
- Leading/trailing spaces: `inWord` handles them naturally

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Counting spaces | `"a  b"` returns 2 | Count word-start transitions |
| No `inWord` flag | Multiple spaces create phantom words | Track `inWord` state |
| Checking `\t` or `\n` as separators too | Might over-split or miss words | This challenge uses only `' '` as separator |

## Solving This Challenge

### Algorithm

1. `count := 0`, `inWord := false`
2. For each rune `c` in `s`:
   - If `c != ' '` and `!inWord`: `count++`, `inWord = true`
   - If `c != ' '` and `inWord`: nothing
   - If `c == ' '`: `inWord = false`
3. Return `count`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [91-findchar](../91-findchar/skills.md)
