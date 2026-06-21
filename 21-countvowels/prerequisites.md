# Prerequisites for 21-countvowels

## Before You Start

To solve this challenge you need to understand:

### 1. The Accumulator Pattern

From [13-stringlength skills.md](../13-stringlength/skills.md) and [19-countalpha skills.md](../19-countalpha/skills.md): count by starting at zero and incrementing when a condition is met:

```go
count := 0
for _, c := range s {
    if condition {
        count++
    }
}
return count
```

### 2. Checking Equality with `==`

From [16-isempty skills.md](../16-isempty/skills.md): use `==` to test whether a rune equals a specific character:

```go
c == 'a'   // true only when c is the letter a
```

### 3. Combining Conditions with `||`

From [19-countalpha skills.md](../19-countalpha/skills.md): use `||` when any one of several conditions should trigger the action:

```go
if c == 'a' || c == 'e' || c == 'i' {
    // c is one of these vowels
}
```

### 4. Case Sensitivity

`'a'` and `'A'` are different rune values. The challenge requires counting both. You can either:
- List all ten vowel characters (5 lowercase + 5 uppercase)
- Convert to lowercase first using the ASCII trick from [17-toupper skills.md](../17-toupper/skills.md): `if c >= 'A' && c <= 'Z' { c = c + 32 }`

## Review If Stuck

- [19-countalpha skills.md](../19-countalpha/skills.md) — counting characters that match a condition
- [17-toupper skills.md](../17-toupper/skills.md) — converting case with ASCII math

## You're Ready When You Can...

- [ ] Count characters in a string that match a specific condition
- [ ] Check a rune against a list of characters using `==` and `||`
- [ ] Handle both uppercase and lowercase variants of the same letter

## Next Steps

- [22-reversestring](../22-reversestring/README.md)
- [22-reversestring README](../22-reversestring/README.md) — next challenge
