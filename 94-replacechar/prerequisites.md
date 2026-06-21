# Prerequisites for 94-replacechar

## Before You Start

To solve this challenge you need to understand:

### 1. Building a Result String with an `if/else`

From [80-toupper skills.md](../80-toupper/skills.md): the pattern where every character is appended — either in a transformed form or unchanged:

```go
result := ""
for _, c := range s {
    if condition(c) {
        result += string(transformed(c))
    } else {
        result += string(c)
    }
}
```

This challenge is identical in structure. The condition is `c == old`, and the transformation is appending `new` instead.

### 2. Functions with Three Parameters

You have written one-parameter and two-parameter functions. This challenge has three:

```go
func ReplaceChar(s string, old, new rune) string
```

`old, new rune` is shorthand for `old rune, new rune` — two rune parameters in one declaration.

### 3. Rune Comparison

From [91-findchar skills.md](../91-findchar/skills.md) and [92-countchar skills.md](../92-countchar/skills.md): `c == old` checks if the current character equals the target.

### 4. `string(rune)` Conversion

From [77-firstchar skills.md](../77-firstchar/skills.md): convert a rune to a one-character string for appending:

```go
string('A')   // "A"
```

## Review If Stuck

- [80-toupper skills.md](../80-toupper/skills.md) — conditional string building with if/else
- [87-removespaces skills.md](../87-removespaces/skills.md) — filtering pattern (for comparison)

## You're Ready When You Can...

- [ ] Write a loop that appends either a replacement or the original character based on a condition
- [ ] Declare a function with three parameters including two runes
- [ ] Explain why there must be an `else` branch (unlike RemoveSpaces)

## Next Steps

- [94-replacechar skills.md](skills.md) — teaches the substitute-or-pass-through pattern
- [95-printif README](../95-printif/README.md) — next challenge
