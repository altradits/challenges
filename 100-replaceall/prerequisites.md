# Prerequisites for replaceall

## Before You Start

To solve this challenge you need to understand:

### 1. String Slicing
Extract a portion of a string with `s[i:j]` — bytes from index `i` up to (not including) `j`.
```go
s := "banana"
fmt.Println(s[0:3]) // "ban"
fmt.Println(s[2:4]) // "na"
fmt.Println(s[4:])  // "na" (to end)
```

### 2. String Comparison with ==
Two strings (or string slices) compare equal with `==` when they have the same content.
```go
s := "banana"
fmt.Println(s[2:4] == "na") // true
fmt.Println(s[0:2] == "ba") // true
```

### 3. Index-Based For Loop
Use `i := 0; for i < len(text)` when you need to advance `i` by variable amounts (not just 1 each time).
```go
i := 0
for i < len(text) {
    // ... do work
    i++ // or i += len(old)
}
```

### 4. Building a Result String
Concatenate characters or substrings to a result variable.
```go
result := ""
result += "hello"
result += string(text[i]) // one character at a time
```

### 5. Bounds Checking Before Slicing
Before doing `text[i:i+n]`, ensure `i+n <= len(text)` to avoid a panic.
```go
if i+len(old) <= len(text) && text[i:i+len(old)] == old {
    // safe to slice
}
```

## Review If Stuck

- [../99-longestword/skills.md](../99-longestword/skills.md) — covers iterating over string content
- [../97-hashcode/skills.md](../97-hashcode/skills.md) — covers building a result string character by character

## You're Ready When You Can...

- [ ] Slice a string with `s[i:j]` and compare the result with `==`
- [ ] Write an index-based loop where `i` advances by different amounts
- [ ] Build a result string by appending to it in a loop
- [ ] Guard slice operations to avoid out-of-bounds panics

## Next Steps

- [101-searchreplace](../101-searchreplace/README.md)
