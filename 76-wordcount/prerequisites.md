# Prerequisites for 76-wordcount

## Before You Start

To solve this challenge you need to understand `strings.Fields`, `len()` on slices, and optionally the state-tracking pattern with a boolean flag.

### 1. `strings.Fields` splits on whitespace, ignoring empty segments

```go
import "strings"

strings.Fields("Hello World")       // ["Hello", "World"]
strings.Fields("one  two   three")  // ["one", "two", "three"]
strings.Fields("   ")               // []
strings.Fields("")                  // []
```

### 2. `len()` on a slice counts elements

```go
words := strings.Fields("Hello World")
len(words)   // 2
```

### 3. The combined one-liner

```go
func WordCount(s string) int {
    return len(strings.Fields(s))
}
```

### 4. Boolean flag for tracking state (manual approach)
A boolean that changes based on what you saw last:

```go
inWord := false
for _, r := range s {
    if r != ' ' && !inWord {
        count++
        inWord = true
    } else if r == ' ' {
        inWord = false
    }
}
```

## Review If Stuck

- [75-titlecase](../75-titlecase/skills.md) — covers `strings.Fields`, `strings.Join`, and boolean state flags
- [67-count-character](../67-count-character/skills.md) — covers the counter pattern with `for...range`

## You're Ready When You Can...

- [ ] Call `strings.Fields(s)` and explain what it returns for edge cases like `""` and `"   "`
- [ ] Apply `len()` to a `[]string` to count its elements
- [ ] Write the `WordCount` function in one line using `strings.Fields`
- [ ] Explain the difference between `strings.Split(s, " ")` and `strings.Fields(s)`

## Next Steps

After completing this exercise, move to:
- [77-cameltosnakecase](../77-cameltosnakecase/README.md)
