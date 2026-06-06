# 29. Join

## What You'll Learn

This exercise teaches you **combining strings with separators**. By the end, you'll understand:
- How to join a slice of strings
- Adding separators between elements
- Handling edge cases (empty slices, empty separators)
- The inverse operation of split

## Theory: String Joining

### What is Joining?

Joining combines multiple strings into one with a separator:

```
["a", "b", "c"] joined with "," → "a,b,c"
["Hello", "World"] joined with " " → "Hello World"
```

### The Join Algorithm

```go
func join(elems []string, sep string) string {
    if len(elems) == 0 {
        return ""
    }
    result := elems[0]
    for i := 1; i < len(elems); i++ {
        result += sep + elems[i]
    }
    return result
}
```

### Empty Separator

When separator is empty, elements are concatenated directly:

```
["a", "b", "c"] joined with "" → "abc"
```

## Step-by-Step Approach

1. **Handle** empty slice → return ""
2. **Start** with first element
3. **Loop** through remaining elements
4. **Add** separator + element for each
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Adding separator before first | Would have leading sep | Start with first element |
| Forgetting empty slice | Might panic | Check `len(elems) == 0` |
| Wrong loop start | Would duplicate first element | Start from index 1 |

## The Challenge

Write a function that joins a slice of strings with a separator.

### Expected Function

```go
func Join(elems []string, sep string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `(["a", "b", "c"], ",")` | `"a,b,c"` | Basic join |
| `(["Hello", "World"], " ")` | `"Hello World"` | Space separator |
| `(["a", "b", "c"], "")` | `"abc"` | Empty separator |
| `([""], ",")` | `""` | Single empty string |
| `([], ",")` | `""` | Empty slice |

## Knowledge Check

Before coding, make sure you can answer:
1. Why start with the first element instead of adding sep before each?
2. What should happen with an empty slice?
3. How is Join the inverse of Split?

## Next Steps

After completing this, you'll be ready for:
- **30-cameltosnakecase**: Advanced string transformation
