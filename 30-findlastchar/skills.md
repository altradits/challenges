# Skills for 30-findlastchar

## What You'll Learn

**Previous:** [29-countchar](../29-countchar/skills.md) | **Next:** [31-replacechar](../31-replacechar/skills.md)

**Challenge:** Write a function `FindLastChar(s string, c rune) int` that returns the byte index of the LAST occurrence of `c` in `s`, or `-1` if not found.

## Core Concept: Keep Updating Instead of Returning Early

### FindChar vs FindLastChar

| Function | Strategy | Stops when? |
|----------|----------|-------------|
| `FindChar` | Return on first match | First match found |
| `FindLastChar` | Record every match, keep going | Always scans entire string |

The key insight: to find the **last** occurrence, you must scan **all** the way to the end — never stop early.

### The "Keep Updating" Pattern

Start with `lastIndex := -1`. Every time a match is found, overwrite `lastIndex`. When the loop ends, you automatically have the position of the final match:

```go
func FindLastChar(s string, c rune) int {
    lastIndex := -1
    for i, ch := range s {
        if ch == c {
            lastIndex = i   // update every time — ends with the last match
        }
    }
    return lastIndex
}
```

Because the loop runs to completion, the last update to `lastIndex` is from the rightmost occurrence.

### Why Start with `-1`?

If `c` never appears in `s`, `lastIndex` is never updated. Returning `-1` correctly signals "not found".

### Contrast with `FindChar`

```go
// FindChar — stops at first match:
for i, ch := range s {
    if ch == c {
        return i   // immediate return
    }
}
return -1

// FindLastChar — scans everything:
lastIndex := -1
for i, ch := range s {
    if ch == c {
        lastIndex = i   // update, but keep going
    }
}
return lastIndex
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `return i` inside the loop | Returns first match, not last | Use `lastIndex = i` (update, not return) |
| Initialising `lastIndex := 0` | Returns 0 when not found (ambiguous) | Use `-1` |
| Forgetting `return lastIndex` | Compile error | Add the return after the loop |

## Solving This Challenge

### Algorithm

1. `lastIndex := -1`
2. For each `i, ch` in `s`:
   - If `ch == c`: `lastIndex = i`
3. Return `lastIndex`

### Trace Through an Example

Input: `FindLastChar("banana", 'a')`

| i | ch | ch=='a'? | lastIndex |
|---|----|---------|-----------|
| 0 | b | No | -1 |
| 1 | a | Yes | 1 |
| 2 | n | No | 1 |
| 3 | a | Yes | 3 |
| 4 | n | No | 3 |
| 5 | a | Yes | 5 |

Return `5` — the last `'a'`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [31-replacechar](../31-replacechar/README.md)
