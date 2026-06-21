# Skills for 28-findchar

## What You'll Learn

**Previous:** [27-wordcount](../27-wordcount/skills.md) | **Next:** [29-countchar](../29-countchar/skills.md)

**Challenge:** Write a function `FindChar(s string, c rune) int` that returns the byte index of the first occurrence of `c` in `s`, or `-1` if not found.

## Core Concept: Returning an Index — Using Both Values from `for...range`

### Two Values from `for...range`

Until now most loops used only the character (`_, c`). This challenge also needs the **position** (byte index):

```go
for i, c := range s {
    // i = byte index of this character in the string
    // c = the rune (character) at that position
}
```

Both `i` and `c` are available simultaneously.

### The Search Pattern

```go
func FindChar(s string, c rune) int {
    for i, ch := range s {
        if ch == c {
            return i   // found — return index immediately
        }
    }
    return -1   // not found after checking everything
}
```

This is early-return from [20-checknumber](../20-checknumber/skills.md), but now you return the **index** instead of `true`.

### Why `-1` for "Not Found"?

Valid indices are `0` and above. Using `-1` as a sentinel is a universal convention in Go (and most languages) for "the element was not found". The caller can check:

```go
idx := FindChar("Hello", 'z')
if idx == -1 {
    fmt.Println("not found")
}
```

Using `0` would be ambiguous (index 0 is the first character).

### The Function Signature Takes a `rune` Parameter

This challenge introduces a function with two parameters:

```go
func FindChar(s string, c rune) int
```

- `s string` — the string to search
- `c rune` — the character to find
- Returns `int` — the index, or `-1`

When calling it, pass a rune literal:

```go
FindChar("Hello", 'l')   // 'l' is a rune literal
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `for _, ch` (discarding index) | Cannot return the position | Use `for i, ch` |
| Returning `0` for "not found" | Ambiguous — 0 is a valid index | Return `-1` |
| Continuing after finding match | Returns last match, not first | `return i` immediately |
| `return i` outside the loop | Compile error if nothing was found | Add `return -1` after the loop |

## Solving This Challenge

### Algorithm

1. For each `i, ch` in `s`:
   - If `ch == c`: return `i`
2. Return `-1`

### Trace Through an Example

Input: `FindChar("banana", 'a')`

| i | ch | ch=='a'? | Action |
|---|----|---------|--------|
| 0 | b | No | continue |
| 1 | a | Yes | return 1 |

Result: `1` (first `'a'` is at index 1)

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [29-countchar](../29-countchar/README.md)
