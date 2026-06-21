# Skills for 99-cleanlist

## What You'll Learn

**Previous:** [98-searchreplace](../98-searchreplace/skills.md) | **Next:** [100-countwords](../100-countwords/skills.md)

**Challenge:** Remove duplicate strings from a slice while preserving order.

## Core Concept: Map as a Set for Deduplication

### What Is It?

A **set** is a collection where each element can appear only once. Go doesn't have a built-in set type, but `map[string]bool` works perfectly — the keys are the unique values.

### How It Works

```go
func CleanList(inputs []string) []string {
    seen := make(map[string]bool)
    result := []string{}
    for _, s := range inputs {
        if !seen[s] {         // not seen yet?
            seen[s] = true    // mark as seen
            result = append(result, s)  // keep it
        }
    }
    return result
}
```

Walk-through with `["a", "b", "a", "c", "b"]`:
- `"a"`: not seen → add, seen={"a"}
- `"b"`: not seen → add, seen={"a","b"}
- `"a"`: already seen → skip
- `"c"`: not seen → add, seen={"a","b","c"}
- `"b"`: already seen → skip
- Result: `["a", "b", "c"]`

### Why `map[string]bool` Works as a Set

Accessing a missing key in a Go map returns the zero value (`false` for bool), so `!seen[s]` is `true` for any key that hasn't been added yet. No need to check `ok` separately.

```go
seen := make(map[string]bool)
seen["hello"]        // false (zero value, not found)
seen["hello"] = true
seen["hello"]        // true (now found)
```

### Alternative: `map[string]struct{}`

Using `struct{}` instead of `bool` saves a tiny amount of memory (empty struct has zero size):

```go
seen := make(map[string]struct{})
if _, ok := seen[s]; !ok {
    seen[s] = struct{}{}
    result = append(result, s)
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Sorting instead of map | Changes order | Use map to preserve insertion order |
| Not pre-allocating result | Minor — still correct | `make([]string, 0, len(inputs))` for efficiency |
| Using `map[string]int` | Works but wasteful | `map[string]bool` is cleaner for sets |

## Algorithm

1. Create `seen := make(map[string]bool)` and `result := []string{}`
2. For each string `s` in input:
   - If `!seen[s]`: mark `seen[s] = true`, append `s` to result
3. Return `result`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [100-countwords](../100-countwords/skills.md)
