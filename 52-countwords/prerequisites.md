# Prerequisites for 52-countwords

## Before You Start

### 1. Iterating Over a String with `for...range`

In Go, `for...range` over a string yields **runes** (Unicode code points), not bytes. This is the correct way to process strings character by character.

```go
s := "Hello!"
for i, c := range s {
    fmt.Printf("index %d: %c\n", i, c)
}
// index 0: H
// index 1: e
// index 2: l
// index 3: l
// index 4: o
// index 5: !
```

### 2. The `rune` Type

A `rune` is Go's type for a single Unicode character (`int32` under the hood). When you range over a string, each character is a `rune`.

```go
var c rune = 'H'
fmt.Println(c >= 'A' && c <= 'Z') // true — uppercase letter
```

### 3. Boolean State Variables (Flags)

A boolean flag tracks a yes/no condition across loop iterations. Flip it when you detect a transition:

```go
inWord := false
for _, c := range s {
    if c >= 'a' && c <= 'z' {
        if !inWord {
            fmt.Println("new word starts")
            inWord = true
        }
    } else {
        inWord = false
    }
}
```

### 4. The `unicode` Package

`unicode.IsLetter(r)` returns true for any letter (including non-ASCII). `unicode.IsDigit(r)` returns true for digits.

```go
import "unicode"

fmt.Println(unicode.IsLetter('a'))  // true
fmt.Println(unicode.IsLetter('!'))  // false
fmt.Println(unicode.IsDigit('5'))   // true
fmt.Println(unicode.IsDigit('z'))   // false
```

### 5. Incrementing a Counter

A simple counter increments by 1 each time a condition is met:

```go
count := 0
for _, c := range "a b c" {
    if c == ' ' {
        count++
    }
}
fmt.Println(count) // 2
```

## Review If Stuck

- Prior string challenges (76-99 series) — covers `for...range` iteration and boolean logic

## You're Ready When You Can...

- [ ] Write a `for...range` loop over a string and print each character
- [ ] Check whether a rune is a letter using `unicode.IsLetter`
- [ ] Use a boolean variable to track state across loop iterations
- [ ] Explain why `"hello,world"` should count as 2 words

## Next Steps

- [53-findsubstring](../53-findsubstring/README.md)
