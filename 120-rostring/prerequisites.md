# Prerequisites for rostring

## Before You Start

To solve this challenge you need to understand:

### 1. `strings.Fields` for Word Splitting
Splits on any whitespace, handles multiple spaces and leading/trailing spaces.

```go
words := strings.Fields("  Let  there be  light  ")
// words = ["Let", "there", "be", "light"]
```

### 2. `strings.Join` for Word Reassembly
Joins a string slice with a single-space separator.

```go
result := strings.Join([]string{"there", "be", "light", "Let"}, " ")
// "there be light Let"
```

### 3. Slice Re-slicing — `words[1:]`
`words[1:]` gives a new slice starting from index 1 to the end. Safe even when there is only one element (returns an empty slice).

```go
words := []string{"a", "b", "c"}
tail := words[1:]  // ["b", "c"]
```

### 4. `append(words[1:], words[0])` — Left Rotation
The standard Go idiom for rotating a slice one position left.

```go
words := []string{"a", "b", "c", "d"}
rotated := append(words[1:], words[0])
// rotated = ["b", "c", "d", "a"]
```

### 5. Handling Empty and Whitespace-Only Input
`strings.Fields` on an empty or all-space string returns an empty slice. Guard against indexing it.

```go
if len(words) == 0 {
    fmt.Println()
    return
}
```

## Review If Stuck
- [119-revwstr](../119-revwstr/skills.md) — teaches `strings.Fields`, `strings.Join`, and the two-pointer slice reversal pattern

## You're Ready When You Can...
- [ ] Use `strings.Fields` to split into tokens ignoring extra spaces
- [ ] Use `strings.Join(words, " ")` to reassemble
- [ ] Use `words[1:]` to get all elements after the first
- [ ] Use `append(words[1:], words[0])` to rotate left by one
- [ ] Guard against an empty `words` slice before indexing

## Next Steps
- [121-wordflip](../121-wordflip/README.md)
