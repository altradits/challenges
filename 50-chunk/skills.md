# Skills for chunk

## What You'll Learn

**Previous:** [../49-canjump/skills.md](../49-canjump/skills.md) | **Next:** [../51-concatalternate/README.md](../51-concatalternate/README.md)

**Challenge:** Split a `[]int` slice into sub-slices of a fixed `size` and print them.

## Core Concept: Slice Chunking With Sub-Slicing

### What Is It?
Divide a slice into groups of `size` elements each. The last group may be smaller if the slice length is not evenly divisible by `size`. Each group is a sub-slice created with `slice[i:j]` syntax.

### How It Works

```go
import "fmt"

func Chunk(slice []int, size int) {
    if size == 0 {
        fmt.Println()
        return
    }
    result := [][]int{}
    for i := 0; i < len(slice); i += size {
        end := i + size
        if end > len(slice) {
            end = len(slice)  // don't go past the end
        }
        result = append(result, slice[i:end])
    }
    fmt.Println(result)
}
```

### Sub-Slicing: slice[i:j]
A sub-slice shares the underlying array — it's a view of elements from index `i` to `j-1`:
```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7}
fmt.Println(s[0:3]) // [0 1 2]
fmt.Println(s[3:6]) // [3 4 5]
fmt.Println(s[6:8]) // [6 7]  — last chunk, smaller
```

### Advancing by `size` Each Iteration
The loop variable increases by `size` each step:
```go
for i := 0; i < len(slice); i += size {
    end := i + size
    if end > len(slice) {
        end = len(slice)
    }
    // slice[i:end] is the current chunk
}
```

### Step-by-Step for `Chunk([]int{0,1,2,3,4,5,6,7}, 3)`
- i=0: end=3 → chunk = `[0 1 2]`
- i=3: end=6 → chunk = `[3 4 5]`
- i=6: end=9 > 8 → end=8 → chunk = `[6 7]`
- result = `[[0 1 2] [3 4 5] [6 7]]`

### The size == 0 Case
The challenge says: if `size == 0`, print a newline. Guard this first:
```go
if size == 0 {
    fmt.Println()
    return
}
```

### Empty Slice
If `slice` is empty, `len(slice) == 0`, the loop doesn't run, and `fmt.Println(result)` prints `[]`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not clamping `end` to `len(slice)` | Slice out of bounds panic on last chunk | `if end > len(slice) { end = len(slice) }` |
| Not handling `size == 0` | Infinite loop (i never advances past 0) | Special-case `size == 0` first |
| Using `append(result, slice[i:end]...)` | Spreads elements instead of appending a sub-slice | Use `append(result, slice[i:end])` without `...` |

## Solving This Challenge

### Algorithm
1. If `size == 0`, print newline and return.
2. Initialize `result := [][]int{}`.
3. Loop with `i += size`: compute `end = min(i+size, len(slice))`, append `slice[i:end]`.
4. Print `result`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../51-concatalternate/README.md](../51-concatalternate/README.md)
