# Prerequisites for 90-wordcount

## Before You Start

To solve this challenge you need to understand:

### 1. The `for...range` Loop

From [76-stringlength skills.md](../76-stringlength/skills.md): iterating over every character in a string using `for _, c := range s`.

### 2. Boolean State Variables

From [88-countrepeats skills.md](../88-countrepeats/skills.md): a boolean variable declared before the loop that persists its value between iterations:

```go
inRepeat := false
for _, c := range s {
    if someCondition {
        inRepeat = true
    } else {
        inRepeat = false
    }
}
```

For this challenge the flag is called `inWord`.

### 3. Comparing a Rune to a Space

You already know rune comparisons from [87-removespaces skills.md](../87-removespaces/skills.md):

```go
c == ' '    // c is a space
c != ' '    // c is not a space (a word character)
```

### 4. Compound `if` Conditions

From [82-countalpha skills.md](../82-countalpha/skills.md): combining conditions with `&&` (and) and `||` (or):

```go
if c != ' ' && !inWord {
    // c is a non-space character AND we are not already inside a word
}
```

`!inWord` means "the flag is currently false" (not in a word).

## Review If Stuck

- [88-countrepeats skills.md](../88-countrepeats/skills.md) — state tracking with a boolean flag across iterations
- [87-removespaces skills.md](../87-removespaces/skills.md) — comparing runes to `' '`

## You're Ready When You Can...

- [ ] Declare a `bool` variable before a loop and update it inside the loop
- [ ] Use `c != ' '` to detect a non-space character
- [ ] Use `&&` to combine two conditions
- [ ] Explain why two consecutive spaces should NOT count as two word boundaries

## Next Steps

- [90-wordcount skills.md](skills.md) — teaches the `inWord` flag pattern for counting word transitions
- [91-findchar README](../91-findchar/README.md) — next challenge
