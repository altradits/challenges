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

## The Challenge

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