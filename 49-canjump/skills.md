# Skills for canjump

## What You'll Learn

**Previous:** [../48-addprimesum/skills.md](../48-addprimesum/skills.md) | **Next:** [../50-chunk/README.md](../50-chunk/README.md)

**Challenge:** Simulate jumping through a slice where each value is the exact number of steps to advance, and determine if the last index is reachable.

## Core Concept: Slice Simulation / Reachability

### What Is It?
This is a simulation problem. You start at index 0. The value at each position tells you exactly how many steps forward to jump. You return true if you can reach the last index without going out of bounds.

### Key Insight: Exact Steps (Not Maximum)
This version requires *exactly* `slice[i]` steps — not "up to" that many. This means there is only one possible path from each position, making it deterministic (no greedy/dynamic programming needed).

```go
func CanJump(nums []uint) bool {
    if len(nums) == 0 {
        return false
    }
    if len(nums) == 1 {
        return true
    }
    pos := 0
    for {
        jump := int(nums[pos])
        pos += jump
        if pos == len(nums)-1 {
            return true   // landed on last index
        }
        if pos >= len(nums) {
            return false  // jumped past the end
        }
        if jump == 0 {
            return false  // stuck: value is 0, can't move
        }
    }
}
```

### Step-by-Step for `[]uint{2, 3, 1, 1, 4}`
- pos=0, jump=2 → pos=2
- pos=2, jump=1 → pos=3
- pos=3, jump=1 → pos=4
- pos=4 == len-1 (4) → return true

### Step-by-Step for `[]uint{3, 2, 1, 0, 4}`
- pos=0, jump=3 → pos=3
- pos=3, jump=0 → jump==0, return false

### Working With []uint
`nums[pos]` is of type `uint`. To add it to an `int` position variable, convert it:
```go
jump := int(nums[pos])
pos += jump
```

### Handling Stuck Positions (Value = 0)
If `nums[pos] == 0` and you're not at the last index, you're stuck. The loop would run forever. Guard with:
```go
if jump == 0 {
    return false
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not checking `jump == 0` | Infinite loop when stuck at a zero | Add `if jump == 0 { return false }` |
| Checking `pos > len(nums)-1` vs `pos >= len(nums)` | Off-by-one | `pos == len(nums)-1` is success; `pos >= len(nums)` is out-of-bounds failure |
| Not handling empty slice | Accessing `nums[0]` panics | Guard with `if len(nums) == 0 { return false }` |

## Solving This Challenge

### Algorithm
1. Return false for empty slice, true for single-element.
2. Start at `pos = 0`.
3. Loop: compute `jump = int(nums[pos])`.
4. If `jump == 0`, return false (stuck).
5. `pos += jump`. If `pos == len-1`, return true. If `pos >= len`, return false.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../50-chunk/README.md](../50-chunk/README.md)
