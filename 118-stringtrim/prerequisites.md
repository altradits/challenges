# Prerequisites for 118-stringtrim

## Before You Start

### 1. `unicode.IsSpace`

`unicode.IsSpace(r)` returns true for any Unicode whitespace character:

```go
import "unicode"

fmt.Println(unicode.IsSpace(' '))   // true — regular space
fmt.Println(unicode.IsSpace('\t'))  // true — tab
fmt.Println(unicode.IsSpace('\n'))  // true — newline
fmt.Println(unicode.IsSpace('a'))   // false — letter
fmt.Println(unicode.IsSpace('5'))   // false — digit
```

### 2. Scanning from the Left with a `start` Pointer

Move `start` forward as long as the character at `start` is whitespace:

```go
start := 0
for start < len(s) && unicode.IsSpace(rune(s[start])) {
    start++
}
// start now points to the first non-whitespace byte
```

### 3. Scanning from the Right with an `end` Pointer

Move `end` backward from the last character:

```go
end := len(s) - 1
for end >= 0 && unicode.IsSpace(rune(s[end])) {
    end--
}
// end now points to the last non-whitespace byte
```

### 4. Slicing Between Two Pointers

Once you have `start` and `end`:

```go
// Include both start and end:
result := s[start : end+1]
// Note: [a:b] is exclusive of b, so use end+1 to include end
```

### 5. All-Whitespace String

If the entire string is whitespace, `start` will equal `len(s)` after the left scan. Check for this:

```go
if start >= len(s) {
    return ""
}
```

Or equivalently, if `start > end` after both scans.

## Review If Stuck

- [100-countwords](../100-countwords/skills.md) — covers `unicode` package for character classification

## You're Ready When You Can...

- [ ] Use `unicode.IsSpace(r)` to detect whitespace
- [ ] Write a forward scan that advances a pointer past whitespace
- [ ] Write a backward scan that retreats a pointer past whitespace
- [ ] Extract the trimmed result with `s[start:end+1]`
- [ ] Handle the all-whitespace edge case

## Next Steps

- [119-stringcontains](../119-stringcontains/README.md)
