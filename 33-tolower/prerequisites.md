# Prerequisites for 33-tolower

## Before You Start

To solve this challenge you need to understand everything from [32-toupper](../32-toupper/README.md). This challenge is structurally identical — only the direction of the ASCII arithmetic changes.

### 1. ASCII Case Conversion

From [32-toupper skills.md](../32-toupper/skills.md): uppercase and lowercase letters differ by 32 in ASCII:

```
'a' - 32 = 'A'   (toupper: subtract 32)
'A' + 32 = 'a'   (tolower: add 32)
```

### 2. Detecting a Specific Character Range

You already know how to detect lowercase letters:

```go
if c >= 'a' && c <= 'z' { /* lowercase */ }
```

For this challenge you need the uppercase range instead:

```go
if c >= 'A' && c <= 'Z' { /* uppercase */ }
```

### 3. Building a Result String in a Loop

From [32-toupper skills.md](../32-toupper/skills.md):

```go
result := ""
for _, c := range s {
    result += string(c)   // accumulate
}
return result
```

### 4. Converting a Rune to a String

After arithmetic on a rune, use `string()` to get text:

```go
var c rune = 'A'
string(c + 32)   // "a"
```

## Review If Stuck

- [32-toupper skills.md](../32-toupper/skills.md) — the identical structure for the opposite operation

## You're Ready When You Can...

- [ ] Write a `ToUpper`-style loop that processes every character
- [ ] Switch the range check from `'a'`–`'z'` to `'A'`–`'Z'`
- [ ] Change the arithmetic from `c - 32` to `c + 32`
- [ ] Verify that non-letter characters pass through unchanged

## Next Steps

- [34-countalpha](../34-countalpha/README.md)
