# Prerequisites for 92-countchar

## Before You Start

To solve this challenge you need to understand:

### 1. Functions with Two Parameters

From [91-findchar skills.md](../91-findchar/skills.md): a function can take more than one argument:

```go
func CountChar(s string, c rune) int {
    // s is the string, c is the character to count
}
```

### 2. The Accumulator Pattern

From [76-stringlength skills.md](../76-stringlength/skills.md) and [82-countalpha skills.md](../82-countalpha/skills.md): start at zero and increment when a condition is met:

```go
count := 0
for _, ch := range s {
    if condition {
        count++
    }
}
return count
```

### 3. Comparing Runes for Equality

From [84-countvowels skills.md](../84-countvowels/skills.md): use `==` to test if a rune matches a specific value:

```go
ch == c   // true when ch equals the target rune c
```

### 4. NOT Returning Early

From [83-checknumber skills.md](../83-checknumber/skills.md), you know early return exists. For counting, you must NOT return early — you need every match. Do NOT put `return` inside the loop for this challenge.

## Review If Stuck

- [91-findchar skills.md](../91-findchar/skills.md) — two-parameter function, rune comparison
- [82-countalpha skills.md](../82-countalpha/skills.md) — counting with a condition

## You're Ready When You Can...

- [ ] Write a function with two parameters (`string` and `rune`)
- [ ] Loop over a string and count characters that equal a specific target
- [ ] Explain why you should NOT return inside the loop for this challenge (unlike FindChar)

## Next Steps

- [92-countchar skills.md](skills.md) — teaches full-scan counting for a specific character
- [93-findlastchar README](../93-findlastchar/README.md) — next challenge
