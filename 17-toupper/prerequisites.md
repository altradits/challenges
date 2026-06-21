# Prerequisites for 17-toupper

## Before You Start

To solve this challenge you need to understand:

### 1. The `for...range` Loop on Strings

From [13-stringlength skills.md](../13-stringlength/skills.md): iterating character by character:

```go
for _, c := range s {
    // c is a rune (int32) — one character
}
```

### 2. Building a New String by Accumulation

Go strings are immutable — you cannot change a character in place. Build a result by starting with `""` and appending:

```go
result := ""
for _, c := range s {
    result += string(c)   // add each character
}
```

You have seen this pattern in [14-firstchar](../14-firstchar/skills.md) and [15-lastchar](../15-lastchar/skills.md) (in a simpler form). Here you do it for every character.

### 3. Runes Are Integers — Arithmetic Works

A `rune` in Go is just `int32`. That means you can do arithmetic on it:

```go
var c rune = 'a'   // c = 97
c + 1              // = 98, which is 'b'
c - 1              // = 96, which is '`'
```

You will use this fact to convert between upper and lowercase.

### 4. Range Comparisons on Runes

To check whether a character falls in a range, compare directly:

```go
if c >= 'a' && c <= 'z' {
    // c is a lowercase ASCII letter
}

if c >= 'A' && c <= 'Z' {
    // c is an uppercase ASCII letter
}
```

`&&` means "and" — both conditions must be true.

### 5. Converting a Rune Back to String

After doing arithmetic on a rune, wrap it in `string()` to produce a one-character string:

```go
string('A')        // "A"
string(rune(65))   // "A"
```

## Review If Stuck

- [13-stringlength skills.md](../13-stringlength/skills.md) — `for...range` loop
- [14-firstchar skills.md](../14-firstchar/skills.md) — `string()` conversion
- [16-isempty skills.md](../16-isempty/skills.md) — boolean conditions and `==`

## You're Ready When You Can...

- [ ] Loop over every character in a string using `for...range`
- [ ] Check whether a rune falls in the range `'a'` to `'z'`
- [ ] Do arithmetic on a rune (e.g., `c - 32`) and convert the result back to a string
- [ ] Build a result string by appending inside a loop

## Next Steps

- [18-tolower](../18-tolower/README.md)
- [18-tolower README](../18-tolower/README.md) — next challenge
