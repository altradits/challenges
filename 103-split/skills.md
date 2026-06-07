# Skills for 103-split

## What You'll Learn

**Previous:** [102-replaceall](../102-replaceall/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Split

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

### 4. Go function definition and usage

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

**Next:** [104-join](../104-join/skills.md) - Join
