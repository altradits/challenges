# Prerequisites for 31-replacechar

## Before You Start

To solve this challenge you need to understand:

### 1. Building a Result String with an `if/else`

From [17-toupper skills.md](../17-toupper/skills.md): the pattern where every character is appended — either in a transformed form or unchanged:

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

From [28-findchar skills.md](../28-findchar/skills.md) and [29-countchar skills.md](../29-countchar/skills.md): `c == old` checks if the current character equals the target.

### 4. `string(rune)` Conversion

From [14-firstchar skills.md](../14-firstchar/skills.md): convert a rune to a one-character string for appending:

```go
string('A')   // "A"
```

## Review If Stuck

- [17-toupper skills.md](../17-toupper/skills.md) — conditional string building with if/else
- [24-removespaces skills.md](../24-removespaces/skills.md) — filtering pattern (for comparison)

## You're Ready When You Can...

- [ ] Write a loop that appends either a replacement or the original character based on a condition
- [ ] Declare a function with three parameters including two runes
- [ ] Explain why there must be an `else` branch (unlike RemoveSpaces)

## Next Steps

- [32-printif](../32-printif/README.md)
- [32-printif README](../32-printif/README.md) — next challenge
