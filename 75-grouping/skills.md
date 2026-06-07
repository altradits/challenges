# Skills for 75-grouping

## What You'll Learn

**Previous:** [74-brainfuck](../74-brainfuck/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Grouping

## New Concepts Explained

### 1. String iteration and character access

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To access individual characters, you can also use indexing, but remember that `s[i]` returns a byte, not a rune. For UTF-8 safety, use `for...range`.

### 2. String searching and indexing

Go provides several ways to search within strings:
- `strings.Index()` - find first occurrence
- `strings.LastIndex()` - find last occurrence
- Manual iteration with `for...range` for custom search logic
- Compare runes or bytes directly

```go
// Manual search example
for i, c := range s {
    if c == target {
        return i
    }
}
return -1
```

### 3. Stack data structure for LIFO operations

A stack is a Last-In-First-Out (LIFO) data structure. In Go, you can implement a stack using a slice:

```go
// Push
stack = append(stack, item)

// Pop
item := stack[len(stack)-1]
stack = stack[:len(stack)-1]

// Peek
item := stack[len(stack)-1]
```

Stacks are useful for bracket matching, RPN evaluation, and backtracking algorithms.

### 4. Looping constructs (for, range)

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

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

**Next:** [76-stringlength](../76-stringlength/skills.md) - Stringlength
