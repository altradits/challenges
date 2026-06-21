# Skills for fifthandskip

## What You'll Learn

**Previous:** [59-wdmatch](../59-wdmatch/skills.md) | **Next:** [61-notdecimal](../61-notdecimal/README.md)

**Challenge:** Collect non-space characters from the string into groups of 5. Separate each complete group with a space, and skip (discard) the 6th character after every group of 5.

## Core Concept: Grouping with Skip-Every-Nth

### What Is It?

This challenge combines two ideas:
1. **Ignore spaces** — strip spaces while counting real characters.
2. **Group into 5, skip 1** — after collecting 5 non-space characters, discard the 6th, then start a new group.

The output groups are separated by spaces (not the 6th character).

### How It Works

**Step 1 — Handle edge cases:**

```go
func FifthAndSkip(str string) string {
    if str == "" {
        return "\n"
    }
    // count non-space characters
    nonSpace := 0
    for _, c := range str {
        if c != ' ' {
            nonSpace++
        }
    }
    if nonSpace < 5 {
        return "Invalid Input\n"
    }
```

**Step 2 — Walk through non-space characters, collecting groups of 5:**

Use a counter that tracks position within the current group of 6 (5 kept + 1 skipped).

```go
    result := ""
    count := 0      // position within group of 6

    for _, c := range str {
        if c == ' ' {
            continue  // skip spaces entirely
        }
        pos := count % 6
        if pos < 5 {
            // positions 0-4: keep
            if pos == 0 && count > 0 {
                result += " "  // separate groups
            }
            result += string(c)
        }
        // position 5: discard (the 6th character)
        count++
    }
    return result + "\n"
}
```

**Trace `"abcdefghijklmnopqrstuwxyz"` (25 chars, no spaces):**

| count | pos | char | action |
|-------|-----|------|--------|
| 0 | 0 | a | keep (group 1 start) |
| 1 | 1 | b | keep |
| 2 | 2 | c | keep |
| 3 | 3 | d | keep |
| 4 | 4 | e | keep (group 1 end) |
| 5 | 5 | f | SKIP |
| 6 | 0 | g | keep (add space, group 2 start) |
| 7 | 1 | h | keep |
| 8 | 2 | i | keep |
| 9 | 3 | j | keep |
| 10 | 4 | k | keep |
| 11 | 5 | l | SKIP |
| ...

Result: `abcde ghijk mnopq stuwx z`

**Trace `"This is a short sentence"` (ignoring internal spaces):**

Non-space chars: `T h i s i s a s h o r t s e n t e n c e` = 20 chars

Groups: `Thisi` (skip `s`) `ashor` (skip `t`) `sente` (skip `n`) `ce`

Result: `Thisi ashor sente ce`

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Counting spaces when forming groups | Groups become wrong length | Use `continue` to skip space chars before counting |
| Adding a space before every group (including the first) | Output starts with a space | Only add a space when `count > 0` and `pos == 0` |
| Off-by-one: skipping the 5th instead of 6th | Groups of 4 instead of 5 | Positions 0–4 (5 chars) are kept; position 5 (the 6th) is skipped |

## Solving This Challenge

### Algorithm
1. If `str == ""`, return `"\n"`.
2. Count non-space chars; if < 5, return `"Invalid Input\n"`.
3. Walk non-space chars with counter `count`. `pos = count % 6`. If `pos < 5`: if `pos == 0 && count > 0` add space separator, then append char. If `pos == 5`: discard. Increment `count`.
4. Return result + `"\n"`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [61-notdecimal](../61-notdecimal/README.md)
