# Prerequisites for piglatin

## Before You Start

To solve this challenge you need to understand:

### 1. String Slicing — Splitting at a Found Index
`word[n:]` gives everything from index `n` to the end; `word[:n]` gives the first `n` characters.

```go
word := "crunch"
firstVowel := 2  // 'u' is at index 2
fmt.Println(word[firstVowel:])   // "unch"
fmt.Println(word[:firstVowel])   // "cr"
fmt.Println(word[firstVowel:] + word[:firstVowel] + "ay")  // "unchcray"
```

### 2. Checking if a Character Is a Vowel
Write a helper function that checks both uppercase and lowercase vowels.

```go
func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
           c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
```

### 3. Finding the First Index Matching a Condition
Use a `for` loop that stops when the condition is true.

```go
firstVowel := 0
for firstVowel < len(word) && !isVowel(word[firstVowel]) {
    firstVowel++
}
// firstVowel is now the index of the first vowel (or len(word) if none)
```

### 4. Scanning for Existence of a Condition
Before transforming, check whether any vowel exists at all.

```go
hasVowel := false
for i := 0; i < len(word); i++ {
    if isVowel(word[i]) {
        hasVowel = true
        break
    }
}
```

### 5. String Concatenation for Transformation
Build the output by joining substrings.

```go
// Vowel-start: word + "ay"
// Consonant-start: word[firstVowel:] + word[:firstVowel] + "ay"
```

## Review If Stuck
- [61-notdecimal](../61-notdecimal/skills.md) — covers string scanning to find a specific character and slicing at that index
- [56-reversestrcap](../56-reversestrcap/skills.md) — covers character-by-character string scanning with conditions

## You're Ready When You Can...
- [ ] Write an `isVowel(c byte) bool` helper covering both cases
- [ ] Use a `for` loop to advance an index until a condition is true
- [ ] Use `word[:n]` and `word[n:]` to split a string at position `n`
- [ ] Concatenate the slices with `"ay"` to form the transformed word
- [ ] Handle the "no vowels" case by scanning the whole word first

## Next Steps
- [71-romannumbers](../71-romannumbers/README.md)
