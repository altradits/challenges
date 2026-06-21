# Prerequisites for reversestrcap

## Before You Start

To solve this challenge you need to understand:

### 1. Case Conversion with `unicode` Package
`unicode.ToUpper(r)` and `unicode.ToLower(r)` convert a rune's case. `unicode.IsLetter(r)` checks if it's a letter.

```go
import "unicode"

fmt.Println(unicode.ToLower('A'))   // 97 ('a')
fmt.Println(unicode.ToUpper('a'))   // 65 ('A')
fmt.Println(unicode.IsLetter('3'))  // false
fmt.Println(unicode.IsLetter('x'))  // true
```

### 2. Converting a String to a `[]rune` for Mutation
Strings are immutable in Go. To modify individual characters, convert to a `[]rune` first.

```go
s := "hello"
runes := []rune(s)
runes[0] = unicode.ToUpper(runes[0])
result := string(runes)  // "Hello"
```

### 3. Look-Ahead — Checking the Next Element
To know if a character is "last in a word," look at the character at `i+1`. Always check bounds first.

```go
for i, r := range runes {
    isLastInWord := i+1 >= len(runes) || !unicode.IsLetter(runes[i+1])
    if unicode.IsLetter(r) && isLastInWord {
        runes[i] = unicode.ToUpper(r)
    }
}
```

### 4. `strings.ToLower` — Lowercase an Entire String
Before applying the "last letter uppercase" rule, lowercase the whole string.

```go
import "strings"

s := "First SMALL TesT"
s = strings.ToLower(s)  // "first small test"
```

### 5. Multiple Arguments with `os.Args[1:]`
The program takes one or more arguments. Loop over all of them.

```go
for _, arg := range os.Args[1:] {
    fmt.Println(process(arg))
}
```

## Review If Stuck
- [124-inter](../124-inter/skills.md) — covers iterating over strings with `for range`
- Prior unicode challenges — covers `unicode.IsLetter`, `unicode.ToUpper`, `unicode.ToLower`

## You're Ready When You Can...
- [ ] Convert a string to `[]rune` and modify individual runes
- [ ] Use `unicode.IsLetter`, `unicode.ToUpper`, `unicode.ToLower`
- [ ] Check `i+1 >= len(runes)` before accessing `runes[i+1]`
- [ ] Use `strings.ToLower` to normalize case before processing
- [ ] Loop over `os.Args[1:]` to process multiple arguments

## Next Steps

- [126-saveandmiss](../126-saveandmiss/README.md)
