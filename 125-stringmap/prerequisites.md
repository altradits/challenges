# Prerequisites for 125-stringmap

## Before You Start

### 1. `strings.Builder` for efficient string construction

```go
var b strings.Builder
b.WriteRune('H')
b.WriteString("ello")
result := b.String()  // "Hello"
```

Review: [113-stringbuilder](../113-stringbuilder/skills.md)

### 2. `for...range` on strings gives runes

```go
for _, r := range "hello" {
    fmt.Printf("%c\n", r)  // prints each character
}
```

### 3. Functions are values in Go

In Go, functions can be stored in variables and passed as arguments:

```go
f := unicode.ToUpper  // f is a variable holding a function
fmt.Println(f('a'))   // 'A'
```

### 4. Function type syntax

```go
func(rune) rune  // type: "a function taking a rune, returning a rune"
```

This is the type of `unicode.ToUpper`, `unicode.ToLower`, and any function you write with that signature.

## You're Ready When You Can...

- [ ] Use `strings.Builder` to build a string rune by rune
- [ ] Iterate over a string with `for...range`
- [ ] Write a function that takes another function as a parameter

## Next Steps

- [126-stringfilter](../126-stringfilter/README.md)
