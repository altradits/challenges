# Skills for slice

## What You'll Learn

**Previous:** [62-revconcatalternate](../62-revconcatalternate/skills.md) | **Next:** [64-findpairs](../64-findpairs/README.md)

**Challenge:** Implement a function `Slice(a []string, nbrs ...int)` that returns a sub-slice of `a` using Python-style indexing: negative indices count from the end, and one or two index arguments are supported.

## Core Concept: Variadic Functions and Negative Index Normalization

### What Is It?

This challenge teaches two things:
1. **Variadic parameters** (`nbrs ...int`) — a function that accepts a variable number of extra arguments.
2. **Negative index normalization** — converting negative indices (like `-1` meaning "last element") into positive ones.

### How It Works

**Step 1 — The variadic function signature:**

```go
func Slice(a []string, nbrs ...int) []string {
```

Inside the function, `nbrs` behaves like a `[]int`. Its length can be 0, 1, or 2.

**Step 2 — Normalize an index (handle negatives):**

A negative index `-n` means "from the end": index `-1` = last element, `-2` = second to last.

```go
func normalize(idx, length int) int {
    if idx < 0 {
        return length + idx  // e.g. -1 with len=5 → 4
    }
    return idx
}
```

**Step 3 — Determine start and end from `nbrs`:**

```go
func Slice(a []string, nbrs ...int) []string {
    length := len(a)

    if len(nbrs) == 0 {
        return a  // no args: return whole slice
    }

    start := normalize(nbrs[0], length)

    if len(nbrs) == 1 {
        if start < 0 || start >= length {
            return nil
        }
        return a[start:]
    }

    // Two arguments: start and end
    end := normalize(nbrs[1], length)

    if start < 0 || end <= 0 || start >= end || start >= length {
        return nil
    }
    if end > length {
        end = length
    }
    return a[start:end]
}
```

**Trace the examples with `a = {"coding","algorithm","ascii","package","golang"}`:**

| Call | start | end | Result |
|------|-------|-----|--------|
| `Slice(a, 1)` | 1 | — | `a[1:]` = `{"algorithm","ascii","package","golang"}` |
| `Slice(a, 2, 4)` | 2 | 4 | `a[2:4]` = `{"ascii","package"}` |
| `Slice(a, -3)` | 5+(-3)=2 | — | `a[2:]` = `{"ascii","package","golang"}` |
| `Slice(a, -2, -1)` | 5+(-2)=3 | 5+(-1)=4 | `a[3:4]` = `{"package"}` |
| `Slice(a, 2, 0)` | 2 | 0 | start>=end → `nil` |

### Variadic Function Internals

```go
func example(prefix string, nums ...int) {
    fmt.Println(len(nums))  // number of extra args
    for _, n := range nums {
        fmt.Println(n)
    }
}
example("x", 1, 2, 3)  // nums = [1, 2, 3]
example("y")            // nums = [] (length 0)
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not normalizing negative indices | `a[-3:]` panics in Go | Convert to `a[len(a)-3:]` |
| Returning `[]string{}` instead of `nil` | Output shows `[]string{}` vs `[]string(nil)` | Use `return nil` for invalid cases |
| Not clamping `end` to `length` | Panic on slice-of-slice with end > length | `if end > length { end = length }` |

## Solving This Challenge

### Algorithm
1. If `len(nbrs) == 0`, return `a`.
2. Normalize `start = nbrs[0]`; if negative, add `len(a)`.
3. If only one `nbrs`: validate `start`, return `a[start:]` or `nil`.
4. If two `nbrs`: normalize `end = nbrs[1]`. Validate `start < end && start >= 0 && end > 0`. Return `a[start:end]` or `nil`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [64-findpairs](../64-findpairs/README.md)
