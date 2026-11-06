# 25. Count Words

## What You'll Learn

This exercise teaches you **advanced word boundary detection with punctuation**. By the end, you'll understand:
- Handling punctuation as word separators
- Multiple delimiter types
- State tracking with complex conditions
- Real-world text processing

## Theory: Word Boundaries

### What Separates Words?

Words can be separated by:
- Spaces: `' '`
- Tabs: `'\t'`
- Newlines: `'\n'`
- Punctuation: `'.'`, `','`, `'!'`, `'?'`, etc.

### Checking for Word Characters

A word character is typically:
- A letter: `'a'-'z'` or `'A'-'Z'`
- A digit: `'0'-'9'`
- Sometimes: `'_'` (underscore)

```go
func isWordChar(c rune) bool {
    return (c >= 'a' && c <= 'z') ||
           (c >= 'A' && c <= 'Z') ||
           (c >= '0' && c <= '9')
}
```

## Step-by-Step Approach

1. **Initialize** counter and in-word flag
2. **Iterate** through characters
3. **If word char AND not in word**: new word, increment
4. **If word char**: set in-word flag
5. **If not word char**: reset flag
6. **Return** count

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Only checking spaces | Misses punctuation | Check all non-word chars |
| Forgetting digits | "hello2world" is 1 word | Include digits as word chars |
| State confusion | Multiple flags needed | Single in-word flag |

## Task

Write a function that counts words, where words are sequences of letters and digits.

### Expected Function

```go
func CountWordsAdvanced(s string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"Hello, World!"` | `2` | Comma and exclamation separate |
| `"Go is fun"` | `3` | Simple spaces |
| `"hello2world"` | `1` | Digits are part of words |
| `""` | `0` | Empty |

## Knowledge Check

Before coding, make sure you can answer:
1. What characters make up a word?
2. How does punctuation affect word boundaries?
3. Why is state tracking important here?

## Next Steps

After completing this, you'll be ready for:
- [53-findsubstring](../53-findsubstring/README.md) - Findsubstring
- [54-replaceall](../54-replaceall/README.md) - Replaceall

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 52
```

or run the tests directly:

```bash
go test ./52-countwords/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 52` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
