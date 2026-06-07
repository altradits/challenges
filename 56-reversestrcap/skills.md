# Skills for 56-reversestrcap

## What You'll Learn

**Previous:** [55-inter](../55-inter/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Reversestrcap

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String transformation and case conversion

Go's `unicode` package provides case conversion functions:
- `unicode.ToUpper(r)` - convert rune to uppercase
- `unicode.ToLower(r)` - convert rune to lowercase
- `unicode.IsUpper(r)` / `unicode.IsLower(r)` - check case

You can also use ASCII math: uppercase and lowercase letters differ by 32.

```go
// ASCII conversion
if c >= 'a' && c <= 'z' {
    c = c - 32  // to uppercase
}
```

### 3. Looping constructs (for, range)

Go has only one looping construct: the `for` loop. It can be used in several ways:

```go
// Traditional for loop
for i := 0; i < 10; i++ { }

// While-style loop
for condition { }

// Range loop (for collections)
for index, value := range collection { }
```

For strings, `for...range` iterates over runes, making it safe for UTF-8.

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

**Next:** [57-saveandmiss](../57-saveandmiss/skills.md) - Saveandmiss
