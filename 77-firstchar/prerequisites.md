# Prerequisites for 77-firstchar

## Before You Start

To solve this challenge you need to understand:

### 1. Writing a Go Function

A function takes an input and returns a result. This challenge needs a function that takes a `string` and returns a `string`:

```go
func FirstChar(s string) string {
    return ""   // placeholder
}
```

If you haven't written a function yet, complete [76-stringlength](../76-stringlength/README.md) first.

### 2. The `for...range` Loop on Strings

From [76-stringlength](../76-stringlength/skills.md): `for...range` iterates over each character (rune) in a string:

```go
for _, c := range s {
    // c is a rune — one character
}
```

For this challenge, you only need the **first** iteration, so you can `return` immediately inside the loop.

### 3. Checking for an Empty String

An empty string has no characters. Before accessing any index, confirm the string is not empty:

```go
if s == "" {
    return ""
}
```

### 4. Converting a `rune` to a `string`

When `for...range` gives you a character, it is a `rune` (an integer representing the Unicode code point). To return it as a `string`, wrap it:

```go
var c rune = 'H'
fmt.Println(c)          // 72  (the number)
fmt.Println(string(c))  // "H" (the text)
```

## Review If Stuck

- [76-stringlength skills.md](../76-stringlength/skills.md) — covers `for...range` on strings and the basics of returning from a function

## You're Ready When You Can...

- [ ] Write a Go function that takes a `string` and returns a `string`
- [ ] Use `for...range` to iterate over a string and get runes
- [ ] Convert a `rune` to a `string` using `string(c)`
- [ ] Guard against an empty string before accessing characters

## Next Steps

- [77-firstchar skills.md](skills.md) — teaches byte indexing and type conversion
- [78-lastchar README](../78-lastchar/README.md) — next challenge
