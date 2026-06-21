# Prerequisites for 82-countalpha

## Before You Start

To solve this challenge you need to understand:

### 1. The Accumulator Pattern

From [76-stringlength skills.md](../76-stringlength/skills.md): counting with a variable that starts at zero:

```go
count := 0
for _, c := range s {
    count++   // add 1 per character
}
return count
```

Here you will only increment when a condition is met, not for every character.

### 2. ASCII Letter Ranges

From [80-toupper skills.md](../80-toupper/skills.md) and [81-tolower skills.md](../81-tolower/skills.md): letters occupy two contiguous ranges:

```
Uppercase: 'A' (65) through 'Z' (90)
Lowercase: 'a' (97) through 'z' (122)
```

Checking a range:

```go
c >= 'A' && c <= 'Z'   // uppercase letter
c >= 'a' && c <= 'z'   // lowercase letter
```

### 3. The `||` Operator (Logical OR)

You need to count letters from EITHER range. Use `||` to combine the two conditions:

```go
if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
    // c is a letter
}
```

`||` is true when at least one side is true.

### 4. `&&` vs `||`

| Operator | Meaning | True when |
|----------|---------|-----------|
| `&&` | AND | Both sides are true |
| `\|\|` | OR | At least one side is true |

## Review If Stuck

- [76-stringlength skills.md](../76-stringlength/skills.md) — `for...range` and counting
- [80-toupper skills.md](../80-toupper/skills.md) — ASCII letter ranges and range checks

## You're Ready When You Can...

- [ ] Loop over a string and count characters that satisfy a condition
- [ ] Write a range check for lowercase letters using `&&`
- [ ] Combine two range checks with `||` to cover both letter cases
- [ ] Return the final count after the loop

## Next Steps

- [82-countalpha skills.md](skills.md) — teaches the combined letter range check
- [83-checknumber README](../83-checknumber/README.md) — next challenge
