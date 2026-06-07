# Skills for 125-stringmap

## What You'll Learn

**Previous:** [124-stringfield](../124-stringfield/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Stringmap

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. Map data structures for lookups

Maps are Go's built-in associative data type (hash tables):

```go
// Create a map
seen := make(map[rune]bool)
seen['a'] = true

// Check if key exists
if seen['b'] {
    // key exists
}

// Delete a key
delete(seen, 'a')
```

Maps are reference types - when you pass them to functions, modifications affect the original.

### 3. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [126-stringfilter](../126-stringfilter/skills.md) - Stringfilter
