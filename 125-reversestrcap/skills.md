# Skills for reversestrcap

## What You'll Learn

**Previous:** [124-inter](../124-inter/skills.md) | **Next:** [126-saveandmiss](../126-saveandmiss/skills.md)

**Challenge:** For each argument, put the last letter of each word in uppercase and all other letters in lowercase. Print the result followed by a newline.

## Core Concept: Per-Word Capitalization Rules

### What Is It?

This challenge requires you to identify word boundaries and apply different case rules depending on whether a letter is the last in its word. A "word" here is defined by spacing — the last letter before a space (or end of string) gets uppercased; all other letters get lowercased.

### How It Works

**Key insight:** You can only tell if a letter is the "last" of a word by looking at what comes AFTER it. The trick: lowercase every letter as you go, and uppercase the letter just before you find a space.

**Two-pass approach — build a byte slice, then fix the last letters:**

```go
func transformArg(arg string) string {
    // Step 1: convert everything to lowercase into a byte slice
    result := []byte(strings.ToLower(arg))

    // Step 2: uppercase the last letter of each word
    for i := 0; i < len(result); i++ {
        ch := result[i]
        // Is this a letter followed by a non-letter (or end)?
        isLetter := (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
        nextIsNonLetter := i+1 >= len(result) || !(result[i+1] >= 'a' && result[i+1] <= 'z')
        if isLetter && nextIsNonLetter {
            // Uppercase it: lowercase letter + 'A' - 'a' = uppercase
            if ch >= 'a' && ch <= 'z' {
                result[i] = ch - 32
            }
        }
    }
    return string(result)
}
```

**Alternative single-pass — look ahead one position:**

```go
func transformArg(arg string) string {
    runes := []rune(strings.ToLower(arg))
    for i, r := range runes {
        isLetter := unicode.IsLetter(r)
        isLastInWord := i+1 >= len(runes) || !unicode.IsLetter(runes[i+1])
        if isLetter && isLastInWord {
            runes[i] = unicode.ToUpper(r)
        }
    }
    return string(runes)
}
```

**Main function — handle multiple arguments:**

```go
func main() {
    if len(os.Args) < 2 {
        return
    }
    for _, arg := range os.Args[1:] {
        fmt.Println(transformArg(arg))
    }
}
```

**Trace `"First SMALL TesT"`:**
1. Lowercase: `"first small test"`
2. Walk: `t` (pos 4) — next char is ` ` (non-letter) → uppercase → `T`
3. Walk: `l` (pos 10) — next char is ` ` (non-letter) → uppercase → `L`
4. Walk: `t` (pos 15) — end of string → uppercase → `T`
5. Output: `firsT smalL tesT`

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Uppercasing chars that aren't letters (e.g. digits, apostrophes) | `"it'S"` instead of `"it'S"` — apostrophe already ignored, but digits should stay as-is | Only uppercase if `unicode.IsLetter(r)` |
| Not handling digits at end of "word" | Digits like `0123456789` are not letters, so the letter before them should be uppercased | Check `!unicode.IsLetter(next)` — a digit counts as non-letter |
| Using `strings.ToUpper` on the whole string first | All letters become upper, then you can't tell original caps | Always lowercase first, then apply the rule |

## Solving This Challenge

### Algorithm
1. If no args, return.
2. For each argument:
   a. Convert to lowercase (all letters).
   b. Walk character by character; if a letter is followed by a non-letter (or end of string), uppercase it.
   c. Print the result followed by a newline.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [126-saveandmiss](../126-saveandmiss/README.md)
