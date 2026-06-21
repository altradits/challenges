# Prerequisites for hiddenp

## Before You Start

To solve this challenge you need to understand:

### 1. `os.Args` and Argument Count Checking
The program takes exactly 2 string arguments. Access them via `os.Args` and validate the count.

```go
import "os"

if len(os.Args) != 3 {  // os.Args[0] is program name
    return
}
s1 := os.Args[1]
s2 := os.Args[2]
```

### 2. `for range` Over a String (Iterating as Runes)
When you range over a string, each iteration gives you a `rune` (Unicode code point). This is the safe way to iterate character by character.

```go
for _, c := range "hello" {
    // c is a rune: 'h', 'e', 'l', 'l', 'o'
    fmt.Printf("%c ", c)
}
```

### 3. Indexing a String as Bytes and Converting to Rune
`s[i]` returns a `byte`. To compare it with a `rune` from `range`, convert it: `rune(s[i])`.

```go
s := "abc"
fmt.Println(s[0])         // 97 (byte value of 'a')
fmt.Println(rune(s[0]))   // 97 (rune value of 'a')
fmt.Println(s[0] == 'a')  // true
```

### 4. A Single Advancing Pointer
Track your progress through `s1` with a single integer index that only moves forward.

```go
i := 0
for _, c := range s2 {
    if i < len(s1) && c == rune(s1[i]) {
        i++
    }
}
```

### 5. `len()` on a String
`len(s)` returns the number of bytes (for ASCII strings, also the number of characters).

```go
fmt.Println(len("hello")) // 5
fmt.Println(len(""))      // 0
```

## Review If Stuck
- [53-fprime](../53-fprime/skills.md) — covers `os.Args` validation and looping patterns
- Prior string iteration challenges — covers `for range` over strings

## You're Ready When You Can...
- [ ] Check `len(os.Args)` and return early if the count is wrong
- [ ] Use `for _, c := range s` to iterate over a string as runes
- [ ] Index a string with `s[i]` and compare with `rune(s[i])`
- [ ] Use a single integer index that only increments on a match
- [ ] Check `i == len(s1)` after the loop to decide true/false

## Next Steps
- [55-inter](../55-inter/README.md)
