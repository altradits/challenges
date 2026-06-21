# Prerequisites for 110-reversestrcap

## Before You Start

### 1. `strings.Fields` — Split on Any Whitespace

`strings.Fields(s)` splits a string on any whitespace (spaces, tabs, newlines) and handles multiple consecutive spaces automatically:

```go
import "strings"

words := strings.Fields("  Hello   World  ")
fmt.Println(words)      // ["Hello" "World"]
fmt.Println(len(words)) // 2
```

### 2. `strings.Join` — Rejoin Words

`strings.Join(words, sep)` concatenates a slice of strings with `sep` between each:

```go
words := []string{"hello", "world"}
fmt.Println(strings.Join(words, " "))   // "hello world"
fmt.Println(strings.Join(words, ", "))  // "hello, world"
```

### 3. Converting a String to a Mutable Rune Slice

Strings are immutable in Go — you cannot modify individual characters. Convert to `[]rune` to get a mutable slice:

```go
word := "Hello"
runes := []rune(word)
runes[0] = 'h'          // modify the first character
fmt.Println(string(runes))  // "hello"
```

### 4. `unicode.ToLower` and `unicode.ToUpper`

```go
import "unicode"

fmt.Println(string(unicode.ToLower('H')))  // "h"
fmt.Println(string(unicode.ToUpper('a')))  // "A"
fmt.Println(string(unicode.ToLower('!')))  // "!" — non-letters unchanged
```

### 5. Accessing the Last Element of a Slice

```go
runes := []rune("Hello")
last := len(runes) - 1
fmt.Println(string(runes[last]))  // "o"
runes[last] = unicode.ToUpper(runes[last])
fmt.Println(string(runes))        // "HellO"
```

## Review If Stuck

- [105-cameltosnakecase](../105-cameltosnakecase/skills.md) — covers `unicode.IsUpper`, `unicode.IsLower` and case detection
- [104-join](../104-join/skills.md) — covers joining strings

## You're Ready When You Can...

- [ ] Use `strings.Fields` to split a string into words on any whitespace
- [ ] Convert a string to `[]rune` and modify individual characters
- [ ] Apply `unicode.ToLower` and `unicode.ToUpper` to individual runes
- [ ] Access the last element of a slice with `slice[len(slice)-1]`
- [ ] Rejoin words with `strings.Join`

## Next Steps

- [111-union](../111-union/README.md)
