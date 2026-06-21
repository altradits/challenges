# Prerequisites for 144-grouping

## Before You Start

### 1. Validating argument count and pattern structure

```go
if len(os.Args) != 3 {
    return
}
pattern := os.Args[1]
if len(pattern) < 3 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
    return
}
```

Review: [81-checknumber](../81-checknumber/skills.md) — byte indexing with `s[0]`, `s[len(s)-1]`.

### 2. `strings.Split` to handle the `|` alternation

```go
inner := pattern[1 : len(pattern)-1]   // strip ( and )
alternatives := strings.Split(inner, "|")
// "e|n" → ["e", "n"]
// "a"   → ["a"]
```

Review: [08-stringspackage](../08-stringspackage/skills.md) — strings.Split, strings.Fields, strings.Contains.

### 3. `strings.Fields` to split the sentence into words

```go
words := strings.Fields(sentence)
```

### 4. `strings.Contains` to test if a word contains the pattern

```go
if strings.Contains(word, alt) {
    count++
    fmt.Printf("%d: %s\n", count, cleanedWord)
}
```

### 5. `unicode.IsLetter` for stripping punctuation from words

```go
import "unicode"

for _, r := range word {
    if unicode.IsLetter(r) || unicode.IsDigit(r) {
        b.WriteRune(r)
    }
}
```

Review: [90-titlecase](../90-titlecase/skills.md) — unicode package usage.

### 6. Nested loops: for each alternative, scan all words

The output order is: for each alternative in order, list matching words in sentence order. This means a word can appear more than once if it matches multiple alternatives.

```go
for _, alt := range alternatives {
    for _, word := range words {
        if strings.Contains(cleanWord(word), alt) { ... }
    }
}
```

## You're Ready When You Can...

- [ ] Validate that a pattern starts with `(` and ends with `)`
- [ ] Strip the outer parentheses with slicing: `pattern[1:len(pattern)-1]`
- [ ] Split on `|` to get alternatives
- [ ] Use `strings.Fields` to split a sentence into words
- [ ] Use nested loops: outer over alternatives, inner over words
- [ ] Clean punctuation off words before matching

## Next Steps

- [153-financial-freedom-api](../153-financial-freedom-api/README.md)
