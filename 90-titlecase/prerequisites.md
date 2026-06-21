# Prerequisites for 90-titlecase

## Before You Start

To solve this challenge you need to understand `strings.Fields`, `unicode.ToUpper`/`unicode.ToLower`, rune slices, and how to track state across a loop.

### 1. `strings.Fields` splits on any whitespace

```go
import "strings"

strings.Fields("hello world")         // ["hello", "world"]
strings.Fields("  extra   spaces  ")  // ["extra", "spaces"]
strings.Fields("")                    // []
```

### 2. `unicode.ToUpper` and `unicode.ToLower`

```go
import "unicode"

unicode.ToUpper('h')   // 'H'
unicode.ToLower('H')   // 'h'
unicode.ToLower('!')   // '!'  (non-letter unchanged)
```

### 3. Rune slices allow mutation

```go
runes := []rune("hello")
runes[0] = unicode.ToUpper(runes[0])   // runes[0] is now 'H'
string(runes)  // "Hello"
```

### 4. Boolean flag to track state across iterations
A flag variable that changes during the loop controls what happens next:

```go
capitalizeNext := true
for _, r := range s {
    if r == ' ' {
        capitalizeNext = true
    } else if capitalizeNext {
        // do the capitalizing...
        capitalizeNext = false
    }
}
```

### 5. `strings.Join` to reassemble words

```go
import "strings"

words := []string{"Hello", "World"}
strings.Join(words, " ")   // "Hello World"
```

## Review If Stuck

- [87-removespaces](../87-removespaces/skills.md) — covers `strings.Builder` and character filtering
- [89-reversestring](../89-reversestring/skills.md) — covers `[]rune` conversion and indexed mutation

## You're Ready When You Can...

- [ ] Use `strings.Fields` and explain what it does with multiple spaces
- [ ] Convert a string to `[]rune`, modify an element, convert back with `string()`
- [ ] Call `unicode.ToUpper(r)` and `unicode.ToLower(r)` on a rune
- [ ] Use a boolean flag variable that gets set inside a loop
- [ ] Use `strings.Join` to combine a `[]string` with a separator

## Next Steps

- [91-wordcount](../91-wordcount/README.md)
