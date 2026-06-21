# Skills for 35-regexp

## What You'll Learn

**Previous:** [19-fileio](../19-fileio/skills.md) | **Next:** [21-sort-and-math](../21-sort-and-math/skills.md)

**Challenge:** Match, extract, and replace text using regular expressions.

## The `regexp` Package

```go
import "regexp"
```

### Compiling a Pattern

```go
// MustCompile — panics if the pattern is invalid (use for static patterns)
re := regexp.MustCompile(`\d+`)

// Compile — returns an error (use when the pattern comes from user input)
re, err := regexp.Compile(`\d+`)
```

Compile once, reuse many times. Compiling is expensive; using the compiled pattern is cheap.

### Testing for a Match

```go
re := regexp.MustCompile(`^\d{3}-\d{4}$`)
fmt.Println(re.MatchString("123-4567"))  // true
fmt.Println(re.MatchString("abc"))       // false
```

### Finding Matches

```go
re := regexp.MustCompile(`\d+`)

// First match
s := re.FindString("abc123def456")   // "123"

// All matches
all := re.FindAllString("abc123def456", -1)  // ["123", "456"]
// -1 means "find all"; pass n to find at most n
```

### Replacing Matches

```go
re := regexp.MustCompile(`\d`)

// Replace all digits with *
result := re.ReplaceAllString("card: 1234 5678", "*")
// "card: **** ****"

// Replace with a function
result = re.ReplaceAllStringFunc("hello123", func(s string) string {
    return "[" + s + "]"
})
// "hello[1][2][3]"
```

### Capture Groups

```go
re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
match := re.FindStringSubmatch("user@example.com")
// match[0] = "user@example.com"  (full match)
// match[1] = "user"              (group 1)
// match[2] = "example"           (group 2)
// match[3] = "com"               (group 3)
```

### Common Patterns

| Pattern | Matches |
|---------|---------|
| `\d` | any digit 0–9 |
| `\w` | word char: letter, digit, underscore |
| `\s` | whitespace |
| `^` | start of string |
| `$` | end of string |
| `+` | one or more |
| `*` | zero or more |
| `?` | zero or one |
| `{n}` | exactly n |
| `[abc]` | a, b, or c |
| `[^abc]` | not a, b, or c |
| `.` | any character (except newline) |

### Solving the Challenge

```go
package piscine

import "regexp"

var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var numRe   = regexp.MustCompile(`\d+`)
var digitRe = regexp.MustCompile(`\d`)

func ValidateEmail(s string) bool {
    return emailRe.MatchString(s)
}

func ExtractNumbers(s string) []string {
    return numRe.FindAllString(s, -1)
}

func MaskPhone(s string) string {
    digits := []rune(s)
    for i := 0; i < len(digits)-4; i++ {
        if digits[i] >= '0' && digits[i] <= '9' {
            digits[i] = '*'
        }
    }
    return string(digits)
}
```

**Next:** [21-sort-and-math](../21-sort-and-math/README.md)
