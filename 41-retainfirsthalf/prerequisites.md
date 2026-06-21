# Prerequisites for 41-retainfirsthalf

## Before You Start

To solve this challenge you need to understand:

### 1. `len(s)` Returns the Byte Count

From [30-lastchar skills.md](../30-lastchar/skills.md): `len(s)` gives you an integer — the number of bytes in the string. For ASCII strings this equals the character count.

```go
len("Hello")   // 5
len("")        // 0
```

### 2. Integer Division Truncates

In Go, dividing two integers produces an integer result — the decimal part is dropped:

```go
5 / 2   // 2  (not 2.5)
4 / 2   // 2
7 / 3   // 2
```

This is different from languages that auto-convert to float. Go never secretly converts `int` to `float`.

### 3. String Slicing Syntax

You can extract a substring with `s[start:end]`:

```go
s := "Hello"
s[0:3]   // "Hel"  — indices 0, 1, 2
s[:3]    // "Hel"  — same: start defaults to 0
s[2:]    // "llo"  — end defaults to len(s)
```

`s[:n]` gives you the first `n` characters. If `n == 0`, you get `""`.

### 4. Conditional Guard Before the Main Logic

From [29-firstchar skills.md](../29-firstchar/skills.md) and [30-lastchar skills.md](../30-lastchar/skills.md): check edge cases first and return early:

```go
if len(s) <= 1 {
    return s
}
// main logic follows
```

## Review If Stuck

- [30-lastchar skills.md](../30-lastchar/skills.md) — `len(s)` and index arithmetic
- [29-firstchar skills.md](../29-firstchar/skills.md) — empty-string guard

## You're Ready When You Can...

- [ ] Call `len(s)` and use the result in arithmetic
- [ ] Explain what `7 / 2` equals in Go (integer division)
- [ ] Use `s[:n]` to extract the first `n` characters from a string
- [ ] Write an early-return guard for edge cases

## Next Steps

- [42-wordcount](../42-wordcount/README.md)
