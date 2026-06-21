# Prerequisites for wdmatch

## Before You Start

To solve this challenge you need to understand:

### 1. Subsequence Matching — Single Pointer Technique
Move a pointer through `word` forward-only as you scan `s`. Only advance the pointer when you find the next expected character.

```go
i := 0
for _, c := range s {
    if i < len(word) && c == rune(word[i]) {
        i++
    }
}
// i == len(word) means full match
```

### 2. `rune(word[i])` — Comparing a Byte from Indexing to a Rune from Range
`word[i]` returns a `byte`; ranging over `s` gives `rune` values. Convert with `rune(word[i])` to compare them.

```go
word := "abc"
fmt.Println(word[0])        // 97 (byte)
fmt.Println(rune(word[0]))  // 97 (rune) — same value, compatible type
```

### 3. Printing Nothing on Failure
If the match fails, the program outputs nothing — not even a newline.

```go
if i == len(word) {
    fmt.Println(word)
}
// no else — silence on failure
```

### 4. `os.Args` Validation
Exactly two string arguments are required.

```go
if len(os.Args) != 3 {
    return
}
word := os.Args[1]
s    := os.Args[2]
```

## Review If Stuck
- [108-hiddenp](../108-hiddenp/skills.md) — teaches exactly this subsequence matching algorithm
- [112-union](../112-union/skills.md) — covers `for range` string iteration and `os.Args`

## You're Ready When You Can...
- [ ] Use a single integer pointer that advances only on a match
- [ ] Compare `rune(word[i])` with a rune obtained from `for _, c := range s`
- [ ] Check `i == len(word)` to confirm a complete match
- [ ] Print the word on success and nothing on failure
- [ ] Validate `len(os.Args) != 3` and return early

## Next Steps
- [114-fifthandskip](../114-fifthandskip/README.md)
