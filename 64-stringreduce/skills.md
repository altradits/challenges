# Skills for 64-stringreduce

## What You'll Learn

**Previous:** [63-stringfilter](../63-stringfilter/skills.md) | **Next:** [65-stringformat](../65-stringformat/skills.md)

**Challenge:** Implement `Reduce(s string, f func(rune, rune) rune) rune` — fold a string to a single value.

## Core Concept: The Fold/Reduce Pattern

### What Is Reduce?

Reduce (also called fold or accumulate) takes a collection and **collapses it to a single value** by applying a combining function repeatedly.

```
"abc"  →  f(f('a', 'b'), 'c')  →  single rune
```

### How It Works

```go
func Reduce(s string, f func(rune, rune) rune) rune {
    if len(s) == 0 {
        return 0  // empty string: return zero value
    }
    runes := []rune(s)
    result := runes[0]           // start with the first element
    for _, r := range runes[1:] { // combine with each remaining element
        result = f(result, r)
    }
    return result
}
```

Walk-through: `Reduce("abc", max)` where max returns the larger rune:
- `result = 'a'` (97)
- `result = max('a', 'b')` = `'b'` (98)
- `result = max('b', 'c')` = `'c'` (99)
- Return `'c'`

### The Accumulator Pattern

The accumulator holds the running result. You update it in each step:

```go
accumulator := initialValue
for each element {
    accumulator = combine(accumulator, element)
}
return accumulator
```

This is the same pattern as summing numbers — `sum += n` — generalized to any combining function.

### Calling Reduce with Different Functions

```go
// Find max character
Reduce("hello", func(a, b rune) rune {
    if a > b { return a }
    return b
})  // 'o' (111, the largest)

// Sum all rune values (numeric ASCII sum)
Reduce("123", func(a, b rune) rune {
    return a + b
})  // '1'(49) + '2'(50) + '3'(51) = 150 = 'ö'

// Find min character
Reduce("hello", func(a, b rune) rune {
    if a < b { return a }
    return b
})  // 'e' (101, the smallest)
```

### Map → Filter → Reduce Together

These three patterns form the foundation of functional programming:

```
Map:    transform each element → same count
Filter: select some elements  → fewer elements
Reduce: combine all elements  → single value
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Starting accumulator at `0` | Wrong for max/min operations | Initialize to `runes[0]` |
| Not handling empty string | Would panic on `runes[0]` | Return `0` if `len(s) == 0` |
| Starting loop at index 0 | Applies function to first element twice | Start at index 1: `runes[1:]` |

## Algorithm

1. Return `0` if `s` is empty
2. Convert `s` to `[]rune`
3. Set `result = runes[0]`
4. For each `r` in `runes[1:]`: `result = f(result, r)`
5. Return `result`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [65-stringformat](../65-stringformat/README.md)
