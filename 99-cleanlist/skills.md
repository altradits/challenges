# Skills for 99-cleanlist

## What You'll Learn

**Previous:** [98-searchreplace](../98-searchreplace/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Cleanlist

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String filtering and cleaning

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

### 3. Slice manipulation and operations

Slices are dynamic, flexible views into arrays. They're the most common data structure in Go:

```go
// Create a slice
numbers := []int{1, 2, 3, 4, 5}

// Slice an existing slice
subset := numbers[1:4]  // [2, 3, 4]

// Append to a slice
numbers = append(numbers, 6)
```

Slices have length (current elements) and capacity (max elements without reallocation).

### 4. Map data structures for lookups

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

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [100-countwords](../100-countwords/skills.md) - Countwords
