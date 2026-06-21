# Skills for 48-union

**Previous:** [47-reversestrcap](../47-reversestrcap/skills.md) | **Next:** [49-inter](../49-inter/skills.md)

**Challenge:** Return a string containing every character that appears in either `s1` or `s2`, with no duplicates, preserving the order characters first appear (first from s1, then new ones from s2).

## Core Concept: Set Union with a Map for Deduplication

### What Is It?

A **set** is a collection with no duplicates. Go has no built-in set type, but you can simulate one with `map[rune]bool`. The map records "have I seen this character before?" so you never add the same character twice.

**Union** combines two sets: all characters from s1 plus any additional characters from s2 that aren't already in the result.

### The Algorithm

```go
func Union(s1, s2 string) string {
    seen := make(map[rune]bool)
    result := ""

    // First, process all of s1
    for _, c := range s1 {
        if !seen[c] {
            seen[c] = true
            result += string(c)
        }
    }

    // Then, add new characters from s2
    for _, c := range s2 {
        if !seen[c] {
            seen[c] = true
            result += string(c)
        }
    }

    return result
}
```

You can simplify by concatenating and iterating once:

```go
for _, c := range s1 + s2 {
    if !seen[c] {
        seen[c] = true
        result += string(c)
    }
}
```

Both are equivalent because `s1 + s2` processes s1 characters before s2 characters.

### How the Map Works as a Set

```go
seen := make(map[rune]bool)

seen['a'] = true      // add 'a' to the set
seen['b'] = true      // add 'b' to the set

fmt.Println(seen['a'])  // true  — 'a' is in the set
fmt.Println(seen['c'])  // false — 'c' not in set (zero value of bool)
```

Checking `!seen[c]` is safe even for keys that were never set, because the zero value of `bool` is `false`.

### Tracing Through `Union("abc", "bcd")`

| c | seen before | Action | seen after | result |
|---|------------|--------|------------|--------|
| 'a' (s1) | false | add | {a} | "a" |
| 'b' (s1) | false | add | {a,b} | "ab" |
| 'c' (s1) | false | add | {a,b,c} | "abc" |
| 'b' (s2) | true  | skip | {a,b,c} | "abc" |
| 'c' (s2) | true  | skip | {a,b,c} | "abc" |
| 'd' (s2) | false | add | {a,b,c,d} | "abcd" |

Result: `"abcd"` — all characters from both, no duplicates, s1 order first.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not using a seen map | Characters can appear twice | Always check `!seen[c]` before adding |
| Iterating s2 only | Misses s1-only characters | Process both strings |
| Using a `[]rune` to track seen | O(n²) lookup instead of O(1) | Use `map[rune]bool` for O(1) lookup |

## Algorithm

1. Create `seen = map[rune]bool{}`
2. For each rune `c` in `s1 + s2`:
   - If `!seen[c]`: set `seen[c] = true`, append `string(c)` to result
3. Return `result`

## The Challenge

See [README.md](README.md).

**Next:** [49-inter](../49-inter/README.md)
