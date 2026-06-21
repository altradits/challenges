# Prerequisites for iscapitalized

## Before You Start

To solve this challenge you need to understand:

### 1. unicode.IsUpper and unicode.IsLetter
The `unicode` package classifies runes. Import it and use its functions to check character type.
```go
import "unicode"

fmt.Println(unicode.IsUpper('A'))  // true
fmt.Println(unicode.IsUpper('a'))  // false
fmt.Println(unicode.IsLetter('A')) // true
fmt.Println(unicode.IsLetter('3')) // false
fmt.Println(unicode.IsLetter('!')) // false
```

### 2. Converting a String to []rune
To safely access the first character of a string (even for Unicode), convert to a rune slice.
```go
word := "Hello"
runes := []rune(word)
firstChar := runes[0] // 'H'
```

### 3. strings.Fields for Word Splitting
Split a string into words, handling multiple spaces automatically.
```go
import "strings"

words := strings.Fields("Hello How Are You")
// words == []string{"Hello", "How", "Are", "You"}
```

### 4. Iterating Over Words and Early Return
Loop over words and return `false` as soon as any word fails the check.
```go
for _, word := range words {
    if !passesCheck(word) {
        return false  // fail fast
    }
}
return true
```

### 5. The Capitalization Condition
A word passes if its first character is NOT a letter, OR it IS an uppercase letter.
```go
first := []rune(word)[0]
if unicode.IsLetter(first) && !unicode.IsUpper(first) {
    return false  // letter but lowercase — fail
}
```

## Review If Stuck

- [../35-clean-the-list/skills.md](../35-clean-the-list/skills.md) — covers `unicode.ToUpper` and working with runes
- [../23-firstword/skills.md](../23-firstword/skills.md) — covers `strings.Fields` for splitting into words

## You're Ready When You Can...

- [ ] Import `unicode` and use `unicode.IsUpper` and `unicode.IsLetter`
- [ ] Convert a string to `[]rune` to safely access the first character
- [ ] Split a string into words with `strings.Fields`
- [ ] Return false early when a word's first character is a lowercase letter

## Next Steps

- [Next challenge](../41-itoa-35/README.md)
