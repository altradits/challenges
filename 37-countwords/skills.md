# Skills for 37-countwords

**Previous:** [36-cleanlist](../36-cleanlist/skills.md) | **Next:** [38-findsubstring](../38-findsubstring/skills.md)

**Challenge:** Count words in a string where words are sequences of letters and digits, separated by any non-word character (spaces, punctuation, newlines, tabs).

## Core Concept: Word Boundary Detection with a State Machine

### What Is It?

A **state machine** tracks whether you are currently "inside a word" or "outside a word". Every character you read either keeps you in the same state or transitions you to a new one. This is the standard technique for word counting because it handles all separator types — spaces, punctuation, tabs, newlines — with one simple rule: any non-word character ends the current word.

### What Makes a Word Character?

A word character is a letter or a digit. Punctuation (`.`, `,`, `!`, `?`, `-`, etc.) is NOT a word character — it acts as a separator just like a space.

```go
import "unicode"

func isWordChar(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}
```

You can also use direct comparisons for ASCII:

```go
func isWordChar(c rune) bool {
    return (c >= 'a' && c <= 'z') ||
           (c >= 'A' && c <= 'Z') ||
           (c >= '0' && c <= '9')
}
```

### How the State Machine Works

Two states: `inWord = false` (outside) and `inWord = true` (inside).

Transitions:
- Outside + word char → **enter word**, increment counter, set `inWord = true`
- Inside + word char → **stay inside**, do nothing
- Inside + non-word char → **exit word**, set `inWord = false`
- Outside + non-word char → **stay outside**, do nothing

```go
func CountWordsAdvanced(s string) int {
    count := 0
    inWord := false

    for _, c := range s {
        if isWordChar(c) {
            if !inWord {
                count++      // just entered a new word
                inWord = true
            }
            // else: still inside the same word, don't count again
        } else {
            inWord = false   // exited the word (any non-word char)
        }
    }
    return count
}
```

### Tracing Through an Example

Input: `"Hello, World!"`

| Char | isWordChar | inWord before | Action | count |
|------|-----------|---------------|--------|-------|
| `H`  | yes | false | enter word | 1 |
| `e`  | yes | true  | stay       | 1 |
| `l`  | yes | true  | stay       | 1 |
| `l`  | yes | true  | stay       | 1 |
| `o`  | yes | true  | stay       | 1 |
| `,`  | no  | true  | exit       | 1 |
| ` `  | no  | false | stay out   | 1 |
| `W`  | yes | false | enter word | 2 |
| `o`  | yes | true  | stay       | 2 |
| `r`  | yes | true  | stay       | 2 |
| `l`  | yes | true  | stay       | 2 |
| `d`  | yes | true  | stay       | 2 |
| `!`  | no  | true  | exit       | 2 |

Result: `2` — correct.

### Why Not Just Split on Spaces?

`strings.Fields("Hello, World!")` returns `["Hello,", "World!"]` — 2 results. Lucky here. But `strings.Fields("hi,there")` returns `["hi,there"]` — 1 result, even though `"hi,there"` contains 2 words separated by a comma. The state machine handles all separator types uniformly.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Only splitting on spaces | `"hello,world"` counts as 1 | Treat all non-word chars as separators |
| Forgetting digits | `"hello2world"` incorrectly counts as 2 | Include `unicode.IsDigit` in word char check |
| Counting on every word char | `"Hi"` counts as 2 instead of 1 | Only count when transitioning `false → true` |

## Algorithm

1. Set `count = 0`, `inWord = false`
2. For each rune `c` in `s`:
   - If `isWordChar(c)` and `!inWord`: increment `count`, set `inWord = true`
   - If `isWordChar(c)` and `inWord`: do nothing (still in same word)
   - If `!isWordChar(c)`: set `inWord = false`
3. Return `count`

## The Challenge

See [README.md](README.md).

**Next:** [38-findsubstring](../38-findsubstring/README.md)
