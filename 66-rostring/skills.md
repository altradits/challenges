# Skills for 66-rostring

## What You'll Learn

**Previous:** [65-revwstr](../65-revwstr/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Rostring

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String splitting and joining

Go's `strings` package provides split and join functions:
- `strings.Split(s, sep)` - split string into slice
- `strings.Join(slice, sep)` - join slice into string
- Manual implementation helps understand the logic

```go
// Manual split example
parts := []string{}
current := ""
for _, c := range s {
    if c == sep {
        parts = append(parts, current)
        current = ""
    } else {
        current += string(c)
    }
}
```

### 3. String filtering and cleaning

Filtering strings involves:
- Iterating through characters
- Checking conditions (is space? is digit? etc.)
- Building a new string with only wanted characters

```go
var result strings.Builder
for _, c := range s {
    if condition(c) {
        result.WriteRune(c)
    }
}
```

### 4. Conditional logic and boolean returns

Go uses `if/else` for conditional branching. The condition doesn't need parentheses:

```go
if condition {
    // do something
} else if otherCondition {
    // do something else
} else {
    // default case
}
```

Boolean operators: `&&` (AND), `||` (OR), `!` (NOT).

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [67-wordflip](../67-wordflip/skills.md) - Wordflip
