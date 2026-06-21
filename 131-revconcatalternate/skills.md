# Skills for revconcatalternate

## What You'll Learn

**Previous:** [130-notdecimal](../130-notdecimal/skills.md) | **Next:** [132-slice](../132-slice/skills.md)

**Challenge:** Like `ConcatAlternate` (51), but traverse both slices from their ENDS toward their starts, still starting with the longer slice's surplus.

## Core Concept: Reverse Alternating Merge

### What Is It?

`RevConcatAlternate` is the mirror of `ConcatAlternate`: instead of interleaving from the front, you interleave from the back of both slices. The longer slice's extra elements (its front portion) come first in the output, then you interleave the rear portions in reverse order.

### How It Works

Compare with ConcatAlternate (51):

| | ConcatAlternate | RevConcatAlternate |
|-|-----------------|-------------------|
| Surplus (longer slice) | front of longer slice, prepended | front of longer slice, prepended |
| Interleave direction | left-to-right (index 0 onward) | right-to-left (last index downward) |
| Which slice leads in interleave | slice1 leads | slice1 leads |

**Step 1 — Determine surplus and prepend it:**

```go
func RevConcatAlternate(slice1, slice2 []int) []int {
    result := []int{}
    len1, len2 := len(slice1), len(slice2)
    minLen := len1
    if len2 < minLen { minLen = len2 }

    if len2 > len1 {
        // slice2 surplus: its first (len2-len1) elements
        for i := 0; i < len2-len1; i++ {
            result = append(result, slice2[i])
        }
    } else if len1 > len2 {
        // slice1 surplus: its first (len1-len2) elements
        for i := 0; i < len1-len2; i++ {
            result = append(result, slice1[i])
        }
    }
```

**Step 2 — Interleave the equal-length tails in REVERSE order:**

Walk backwards from the last element of each slice's equal portion.

```go
    for i := minLen - 1; i >= 0; i-- {
        result = append(result, slice1[len1-minLen+i])
        result = append(result, slice2[len2-minLen+i])
    }
    return result
}
```

**Trace `{1,2,3}` and `{4,5,6}` (equal length, minLen=3):**

- No surplus.
- i=2: slice1[2]=3, slice2[2]=6
- i=1: slice1[1]=2, slice2[1]=5
- i=0: slice1[0]=1, slice2[0]=4
- Result: `[3 6 2 5 1 4]` ✓

**Trace `{1,2,3}` and `{4,5,6,7,8,9}` (slice2 longer, minLen=3):**

- Surplus from slice2: indices 0,1,2 → `4, 5, 6`
- Interleave (reverse of equal tails):
  - i=2: slice1[2]=3, slice2[5]=9
  - i=1: slice1[1]=2, slice2[4]=8
  - i=0: slice1[0]=1, slice2[3]=7
- Result: `[4 1 5 2 6 3 7 8 9]` → wait, expected `[9 8 7 3 6 2 5 1 4]`

Wait — re-reading the README: `{1,2,3}` and `{4,5,6,7,8,9}` → `[9 8 7 3 6 2 5 1 4]`. The surplus for the longer slice comes FIRST but we also interleave in reverse. Let me re-trace carefully:

The example output `[9 8 7 3 6 2 5 1 4]` means: 9(s2[5]), 8(s2[4]), 7(s2[3]) then 3(s1[2]),6(s2[2]), 2(s1[1]),5(s2[1]), 1(s1[0]),4(s2[0]).

So the surplus of s2 (its extra elements) are `s2[3], s2[4], s2[5]` — reversed! Then interleave the equal parts reversed.

```go
    if len2 > len1 {
        // surplus = s2[minLen..] in REVERSE
        for i := len2 - 1; i >= minLen; i-- {
            result = append(result, slice2[i])
        }
    } else if len1 > len2 {
        // surplus = s1[minLen..] in REVERSE
        for i := len1 - 1; i >= minLen; i-- {
            result = append(result, slice1[i])
        }
    }
    // Interleave equal parts in reverse
    for i := minLen - 1; i >= 0; i-- {
        result = append(result, slice1[i])
        result = append(result, slice2[i])
    }
```

**Trace `{1,2,3,9,8}` and `{4,5}` (slice1 longer, minLen=2):**
- surplus = s1[2..4] reversed = s1[4]=8, s1[3]=9, s1[2]=3 → `[8, 9, 3]`
- Interleave reversed: i=1: s1[1]=2,s2[1]=5; i=0: s1[0]=1,s2[0]=4
- Result: `[8 9 3 2 5 1 4]` ✓

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forward surplus (not reversed) | Wrong output for longer slice | Iterate surplus in reverse: `for i := len-1; i >= minLen; i--` |
| Using offsets from ConcatAlternate | Off-by-one in reverse iteration | Interleave equal parts: `slice1[i]` and `slice2[i]` where `i` goes from `minLen-1` down to `0` |

## Solving This Challenge

### Algorithm
1. Compute `minLen`.
2. Surplus (longer slice's elements beyond `minLen`): iterate from the last index down to `minLen` and append.
3. Interleave equal tails: for `i` from `minLen-1` down to 0, append `slice1[i]` then `slice2[i]`.
4. Return result.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [132-slice](../132-slice/README.md)
