# 15. Word Count

## What You'll Learn

This exercise teaches you **word boundary detection and state tracking**. By the end, you'll understand:
- What defines a "word" in text processing
- Tracking state across iterations (in-word vs. between-words)
- Handling multiple consecutive delimiters
- The difference between characters and words

## Theory: Words and Delimiters

### What is a Word?

A word is a **sequence of non-space characters** separated by spaces:

```
"Hello World" → ["Hello", "World"] → 2 words
"Go is fun"   → ["Go", "is", "fun"] → 3 words
```

### Delimiters

A delimiter is a character that **separates** elements:

| Delimiter | Description | Code |
|-----------|-------------|------|
| Space | Most common word separator | `' '` |
| Tab | Horizontal tab | `'\t'` |
| Newline | Line break | `'\n'` |

**This exercise uses space (`' '`) as the delimiter.**

### State Tracking Pattern

To count words, track whether you're currently **inside a word**:

```go
inWord := false
wordCount := 0

for _, c := range s {
    if c != ' ' {
        if !inWord {
            wordCount++      // Found start of new word
            inWord = true
        }
    } else {
        inWord = false       // Space means we're between words
    }
}
```

### Why State Tracking?

Without tracking state, `"Hello  World"` (two spaces) would be miscounted:

```
Without state: H(e)l(l)o( ) ( )W(o)r(l)d → counts 2 transitions → 2 words ✓
With bad logic: Might count the gap as a word → wrong!
```

## Step-by-Step Approach

1. **Initialize**: `inWord = false`, `count = 0`
2. **Iterate** through each character
3. **If not space AND not in word**: new word found, increment count
4. **If not space**: set `inWord = true`
5. **If space**: set `inWord = false`
6. **Return** count

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Counting spaces | Spaces aren't words | Count transitions from space to non-space |
| Forgetting state | Multiple spaces cause issues | Track `inWord` flag |
| Not handling empty | Returns 1 instead of 0 | Empty string = 0 words |

## Practice Tips

- Trace `"Hello  World"`: H(new word, count=1) e l l o (space, inWord=false) (space, inWord=false) W(new word, count=2) o r l d
- Leading/trailing spaces should be ignored
- Multiple spaces between words should be ignored

## The Challenge

Write a function that counts words in a string (words = sequences of non-space characters).

### Expected Function

```go
func WordCount(s string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello World"` | `2` | Two words |
| `"Go is fun!"` | `3` | Three words |
| `""` | `0` | Empty string |
| `"   "` | `0` | Only spaces |
| `"One  two   three"` | `3` | Multiple spaces ignored |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(WordCount("Hello World"))     // 2
    fmt.Println(WordCount("Go is fun!"))      // 3
    fmt.Println(WordCount(""))                // 0
    fmt.Println(WordCount("   "))             // 0
    fmt.Println(WordCount("One  two   three")) // 3
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What defines a word boundary?
2. Why do we need the `inWord` flag?
3. How should multiple spaces be handled?

## Next Steps

After completing this, you'll be ready for:
- [91-findchar](../91-findchar/README.md) - Findchar
- [92-countchar](../92-countchar/README.md) - Countchar