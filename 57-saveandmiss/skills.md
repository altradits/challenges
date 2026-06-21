# Skills for saveandmiss

## What You'll Learn

**Previous:** [56-reversestrcap](../56-reversestrcap/skills.md) | **Next:** [58-union](../58-union/README.md)

**Challenge:** Walk through a string in groups of `num` characters — keep the first group, skip the second, keep the third, skip the fourth, and so on.

## Core Concept: Periodic Selection with a Counter

### What Is It?

"Save and miss" is a **periodic selection** pattern: you alternate between keeping `n` characters and discarding `n` characters. A counter tracks your position within each group of `n`. A boolean (or integer) flag tracks whether you're currently in a "save" or "miss" phase.

### How It Works

**Step 1 — Handle the edge case:**

```go
func SaveAndMiss(arg string, num int) string {
    if num <= 0 {
        return arg
    }
```

**Step 2 — Walk through the string with a counter:**

```go
    result := ""
    count := 0   // position within the current group
    save := true // are we in "save" mode?

    for _, c := range arg {
        if save {
            result += string(c)
        }
        count++
        if count == num {
            count = 0
            save = !save  // flip between save and miss
        }
    }
    return result
}
```

**Trace `"123456789"` with `num = 3`:**

| Char | count (before++) | save? | Action | count after | Flip? |
|------|-----------------|-------|--------|-------------|-------|
| `1` | 0 | true | keep | 1 | no |
| `2` | 1 | true | keep | 2 | no |
| `3` | 2 | true | keep | 3==num → 0 | YES → save=false |
| `4` | 0 | false | skip | 1 | no |
| `5` | 1 | false | skip | 2 | no |
| `6` | 2 | false | skip | 3==num → 0 | YES → save=true |
| `7` | 0 | true | keep | 1 | no |
| `8` | 1 | true | keep | 2 | no |
| `9` | 2 | true | keep | 3==num → 0 | YES → save=false |

Result: `"123789"` ✓

**Trace `"abcdefghijklmnopqrstuvwyz"` with `num = 3`:**
Groups: `abc` (save) `def` (miss) `ghi` (save) `jkl` (miss) `mno` (save) `pqr` (miss) `stu` (save) `vwy` (miss) `z` (save)
Result: `abcghimnostuz` ✓

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Starting `save = false` | Skips the first group instead of saving it | Always start with `save = true` |
| Checking `count == num` before incrementing | Off-by-one: group ends after `num` characters, not before | Increment first, then check `if count == num` |
| Not resetting `count` to 0 after flipping | Counter keeps growing past `num`, flip never triggers again | Set `count = 0` when flipping |
| Returning `""` when `num <= 0` | Should return the original string unchanged | Return `arg` directly when `num <= 0` |

## Solving This Challenge

### Algorithm
1. If `num <= 0`, return `arg` unchanged.
2. Initialize `count = 0`, `save = true`, `result = ""`.
3. For each rune `c` in `arg`: if `save`, append `c` to result. Increment `count`. If `count == num`, reset `count = 0` and flip `save`.
4. Return result.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [58-union](../58-union/README.md)
