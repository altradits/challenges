# Skills for concatalternate

## What You'll Learn

**Previous:** [104-chunk](../104-chunk/skills.md) | **Next:** [106-concatslice](../106-concatslice/skills.md)

**Challenge:** Build a new int slice by alternating elements from two input slices, leading with the larger slice's surplus elements.

## Core Concept: Alternating Merge of Two Slices

### What Is It?

An alternating merge picks one element from slice1, then one from slice2, then slice1 again, and so on. When one slice is longer, its extra elements come first (before the equal-length interleave begins). This pattern appears in merge-sort, playlist shuffling, and balanced round-robin scheduling.

### How It Works

**Step 1 — Measure both slices:**

```go
func ConcatAlternate(slice1, slice2 []int) []int {
    result := []int{}
    len1 := len(slice1)
    len2 := len(slice2)
```

**Step 2 — Drain the surplus from the longer slice first:**

```go
    if len2 > len1 {
        // slice2 has extra elements — they come first
        for i := 0; i < len2-len1; i++ {
            result = append(result, slice2[i])
        }
    } else if len1 > len2 {
        // slice1 has extra elements — they come first
        for i := 0; i < len1-len2; i++ {
            result = append(result, slice1[i])
        }
    }
```

**Step 3 — Interleave the equal-length tails:**

```go
    minLen := len1
    if len2 < minLen {
        minLen = len2
    }
    offset1 := len1 - minLen  // where the equal portion of slice1 starts
    offset2 := len2 - minLen  // where the equal portion of slice2 starts

    for i := 0; i < minLen; i++ {
        result = append(result, slice1[offset1+i])
        result = append(result, slice2[offset2+i])
    }
    return result
}
```

**Trace through the examples:**

| Input | Surplus | Interleave | Output |
|-------|---------|------------|--------|
| `{1,2,3}`, `{4,5,6}` | none (equal) | 1,4,2,5,3,6 | `[1 4 2 5 3 6]` |
| `{2,4,6,8,10}`, `{1,3,5,7,9,11}` | slice2 extra: 1 | 2,3,4,5,6,7,8,9,10,11 → see output | `[1 2 3 4 5 6 7 8 9 10 11]` |
| `{1,2,3}`, `{4,5,6,7,8,9}` | slice2 extra: 4,5,6 | 7,1,8,2,9,3 | `[4 1 5 2 6 3 7 8 9]` |
| `{1,2,3}`, `{}` | slice1 extra: 1,2,3 | nothing | `[1 2 3]` |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Starting interleave at index 0 for both slices | Wrong results when slices differ in length | Calculate `offset = len - minLen` to find where the equal portion starts |
| Forgetting the equal-length case | Offsets are both 0, interleave runs normally — just make sure your condition covers it | The `else if` covers this: no surplus means interleave starts at index 0 |
| Putting slice1's surplus first even when slice2 is longer | Output order wrong | Only slice1's surplus goes first when `len1 > len2` |

## Solving This Challenge

### Algorithm
1. Compute `len1 = len(slice1)`, `len2 = len(slice2)`.
2. If `len2 > len1`: append `slice2[0 : len2-len1]` to result.
3. If `len1 > len2`: append `slice1[0 : len1-len2]` to result.
4. Compute `minLen` = the smaller length.
5. Loop `i` from 0 to `minLen-1`: append `slice1[offset1+i]`, then `slice2[offset2+i]`.
6. Return result.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [106-concatslice](../106-concatslice/README.md)
