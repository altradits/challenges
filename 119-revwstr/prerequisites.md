# Prerequisites for revwstr

## Before You Start

To solve this challenge you need to understand:

### 1. `strings.Fields` — Split by Whitespace
Splits a string into words, automatically handling multiple spaces and leading/trailing whitespace.

```go
import "strings"

words := strings.Fields("  hello   world  ")
// words = ["hello", "world"]
fmt.Println(len(words)) // 2
```

### 2. `strings.Join` — Rejoin a Slice into a String
Join a `[]string` with a separator.

```go
words := []string{"contempt", "of", "time", "the"}
result := strings.Join(words, " ")
fmt.Println(result)  // "contempt of time the"
```

### 3. Two-Pointer In-Place Slice Reversal
Swap elements from both ends toward the center.

```go
s := []string{"a", "b", "c", "d"}
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
}
// s = ["d", "c", "b", "a"]
```

### 4. Simultaneous Assignment
Go allows swapping two variables without a temporary: `a, b = b, a`.

```go
x, y := 1, 2
x, y = y, x
fmt.Println(x, y)  // 2 1
```

### 5. Handling Empty String Input
An empty string argument produces an empty `words` slice. Print a bare newline.

```go
if s == "" {
    fmt.Println()
    return
}
```

## Review If Stuck
- [114-fifthandskip](../114-fifthandskip/skills.md) — covers string scanning and space handling
- [106-concatslice](../106-concatslice/skills.md) — covers slice operations

## You're Ready When You Can...
- [ ] Use `strings.Fields` to split a string into a word slice
- [ ] Use `strings.Join` to reassemble a slice into a string
- [ ] Reverse a `[]string` in-place using a two-pointer swap
- [ ] Use simultaneous assignment `a, b = b, a` without a temp variable
- [ ] Handle the empty-string edge case with a bare `fmt.Println()`

## Next Steps
- [120-rostring](../120-rostring/README.md)
