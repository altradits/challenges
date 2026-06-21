# Skills for 75-grouping

## What You'll Learn

**Previous:** [74-brainfuck](../74-brainfuck/skills.md) | **Next:** [76-stringlength](../76-stringlength/skills.md)

**Challenge:** Given a regex-like group pattern (e.g. `(a)` or `(e|n)`) and a string, print each word that contains the pattern, numbered in order of appearance.

## Core Concept: Word Filtering with Substring Matching and the `|` Alternation

### What the Program Does

```
$ go run . "(a)" "I'm heavy, jumpsuit is on steady"
1: heavy
2: steady
```

The pattern `(a)` means: find all words that contain the letter `a`.

```
$ go run . "(e|n)" "I currently have 4 windows opened"
1: currently
2: currently
3: have
4: windows
...
```

The pattern `(e|n)` means: find all words that contain `e` OR `n`. A word matching both is listed **twice** — once per matching alternative.

### Step 1: Validate the Pattern

A valid pattern must start with `(` and end with `)`:
```go
if len(pattern) < 3 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
    return // invalid — print nothing
}
inner := pattern[1 : len(pattern)-1]  // strip the parentheses
```

### Step 2: Split on `|`

```go
import "strings"

alternatives := strings.Split(inner, "|")
// "(e|n)" → inner = "e|n" → alternatives = ["e", "n"]
// "(a)"   → inner = "a"   → alternatives = ["a"]
```

### Step 3: Split the String into Words

```go
words := strings.Fields(sentence)
```

### Step 4: For Each Alternative, Find Matching Words

The key: for each alternative, scan all words in order. This ensures words appear in sentence order for each alternative, and a word matching both `e` and `n` appears in both passes.

```go
count := 0
for _, alt := range alternatives {
    for _, word := range words {
        // strip punctuation? No — strings.Contains handles substrings
        if strings.Contains(word, alt) {
            count++
            fmt.Printf("%d: %s\n", count, word)
        }
    }
}
```

Wait — but the examples show words stripped of punctuation (e.g., `"heavy,"` becomes `"heavy"`). You may need to clean each word:

```go
import "unicode"

func cleanWord(w string) string {
    var b strings.Builder
    for _, r := range w {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            b.WriteRune(r)
        }
    }
    return b.String()
}
```

### Complete Implementation

```go
package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

func cleanWord(w string) string {
    var b strings.Builder
    for _, r := range w {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            b.WriteRune(r)
        }
    }
    return b.String()
}

func main() {
    if len(os.Args) != 3 {
        return
    }
    pattern := os.Args[1]
    sentence := os.Args[2]

    if sentence == "" {
        return
    }
    if len(pattern) < 3 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
        return
    }

    inner := pattern[1 : len(pattern)-1]
    if inner == "" {
        return
    }

    alternatives := strings.Split(inner, "|")
    words := strings.Fields(sentence)

    count := 0
    for _, alt := range alternatives {
        for _, word := range words {
            cleaned := cleanWord(word)
            if strings.Contains(cleaned, alt) {
                count++
                fmt.Printf("%d: %s\n", count, cleaned)
            }
        }
    }
}
```

### Key Go Techniques

**`strings.Split(s, "|")`** — splits on the pipe character:
```go
strings.Split("e|n", "|")  // ["e", "n"]
strings.Split("a", "|")    // ["a"]
```

**`strings.Fields(s)`** — splits on any whitespace:
```go
strings.Fields("hello world")  // ["hello", "world"]
```

**`strings.Contains(word, sub)`** — true if word contains sub:
```go
strings.Contains("heavy", "a")  // true
strings.Contains("light", "a")  // false
```

**`unicode.IsLetter(r)`** — true for alphabetic runes, false for punctuation:
```go
unicode.IsLetter('h')  // true
unicode.IsLetter(',')  // false
```

**`fmt.Printf("%d: %s\n", count, word)`** — numbered output.

### Understanding the Double-Count Behavior

For `(e|n)` applied to the word `"currently"`:
- "currently" contains `e` → match 1
- "currently" contains `n` → match 2

The two alternatives are processed separately, so one word can appear multiple times. This is intentional — you're matching patterns, not unique words.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [76-stringlength](../76-stringlength/skills.md) — Start of the main exercises series
