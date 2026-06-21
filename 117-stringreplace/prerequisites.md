# Prerequisites for 117-stringreplace

## Before You Start

### 1. `strings.Index`

`strings.Index(s, substr)` returns the byte index of the first occurrence of `substr` in `s`. Returns -1 if not found.

```go
import "strings"

fmt.Println(strings.Index("hello world", "world"))  // 6
fmt.Println(strings.Index("hello world", "xyz"))    // -1
fmt.Println(strings.Index("abcabc", "bc"))          // 1
fmt.Println(strings.Index("", "abc"))               // -1
fmt.Println(strings.Index("hello", ""))             // 0
```

### 2. String Slicing Around a Match

Once you have the index where `old` starts, you can extract:
- The text before the match: `s[:idx]`
- The text after the match: `s[idx+len(old):]`

```go
s := "hello world"
idx := strings.Index(s, "world")  // 6
before := s[:idx]                  // "hello "
after  := s[idx+len("world"):]    // ""
```

### 3. `strings.Builder` (from 113-stringbuilder)

Efficient string building — write each piece with `b.WriteString`, get result with `b.String()`.

### 4. The "Remaining" Pattern

Process a string by repeatedly slicing off the front as you consume it:

```go
remaining := "foo bar foo"
for {
    idx := strings.Index(remaining, "foo")
    if idx == -1 {
        break  // no more matches
    }
    fmt.Println("match at", idx, "in:", remaining)
    remaining = remaining[idx+3:]  // advance past "foo"
}
```

### 5. Why Empty `old` Must Be Guarded

`strings.Index(s, "")` returns 0 (empty string is found at position 0 of everything). If `old == ""`, the remaining never advances past position 0, causing an infinite loop. Always check:

```go
if old == "" {
    return s
}
```

## Review If Stuck

- [102-replaceall](../102-replaceall/skills.md) — covers the same algorithm using manual character scanning
- [113-stringbuilder](../113-stringbuilder/skills.md) — covers `strings.Builder`

## You're Ready When You Can...

- [ ] Use `strings.Index(s, substr)` and understand that -1 means not found
- [ ] Extract text before/after a match with `s[:idx]` and `s[idx+len(old):]`
- [ ] Use `strings.Builder` to accumulate the result
- [ ] Write a loop that terminates when `strings.Index` returns -1

## Next Steps

- [118-stringtrim](../118-stringtrim/README.md)
