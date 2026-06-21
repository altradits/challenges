# Skills for countrepeats

## What You'll Learn

**Previous:** [92-cameltosnakecase](../92-cameltosnakecase/skills.md) | **Next:** [94-digitlen](../94-digitlen/skills.md)

**Challenge:** Count the number of groups of consecutive repeated characters in a string.

## Core Concept: Detecting Consecutive Repetition Groups

### What Is It?
A "repeat group" is when the same character appears two or more times in a row. `"aaa"` is one group; `"aabb"` is two groups. The key skill is comparing each character to the one before it, and using a flag to count only the *start* of each group — not every matching pair.

### How It Works

Compare `s[i]` with `s[i-1]`. When you first enter a repeat (same as previous, and you weren't already in one), increment the counter. Use a boolean `inRepeat` flag to track state.

```go
func CountRepeats(s string) int {
    if len(s) < 2 {
        return 0
    }
    count := 0
    inRepeat := false
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            if !inRepeat {
                count++
                inRepeat = true
            }
        } else {
            inRepeat = false
        }
    }
    return count
}
```

Step-by-step for `"heelloo"` (`h`,`e`,`e`,`l`,`l`,`o`,`o`):
- i=1: `e==h`? No → inRepeat=false
- i=2: `e==e`? Yes, !inRepeat → count=1, inRepeat=true
- i=3: `l==e`? No → inRepeat=false
- i=4: `l==l`? Yes, !inRepeat → count=2, inRepeat=true
- i=5: `o==l`? No → inRepeat=false
- i=6: `o==o`? Yes, !inRepeat → count=3, inRepeat=true
- Result: 3

### Why the `inRepeat` Flag?
Without the flag, a run of `"aaaa"` would match at i=1, i=2, and i=3 — giving count=3 instead of 1. The flag ensures you only count the *transition into* a repeat, not every match within it.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Starting loop at `i=0` | Accesses `s[i-1]` which is `s[-1]` — panic | Start at `i := 1` |
| No `inRepeat` flag | `"aaaa"` gives 3 instead of 1 | Track whether you are already inside a group |
| No empty string check | Panic on `len(s) < 2` when trying `s[i-1]` | Guard with `if len(s) < 2 { return 0 }` |

## Solving This Challenge

### Algorithm
1. Return 0 if `len(s) < 2`.
2. Initialize `count := 0`, `inRepeat := false`.
3. Loop `i` from 1 to `len(s)-1`.
4. If `s[i] == s[i-1]` and `!inRepeat`: increment count, set inRepeat=true.
5. If `s[i] != s[i-1]`: set inRepeat=false.
6. Return count.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [94-digitlen](../94-digitlen/README.md)
