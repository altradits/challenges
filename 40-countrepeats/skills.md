# Skills for 40-countrepeats

## What You'll Learn

**Previous:** [39-removespaces](../39-removespaces/skills.md) | **Next:** [41-retainfirsthalf](../41-retainfirsthalf/skills.md)

**Challenge:** Write a function `CountRepeats(s string) int` that counts the number of groups of consecutive repeated characters. `"heelloo"` has three groups (`ee`, `ll`, and `oo`) so it returns `3`. `"aaa"` is one group, so it returns `1`.

## Core Concept: State Tracking Across Iterations

### What Is State Tracking?

Sometimes you need to remember something from the previous iteration to make a decision in the current one. A **state variable** holds information that persists across loop iterations.

### The Problem

Simply counting "every time the current char matches the previous char" gives the wrong answer for `"aaa"`:
- `a` == `a` (index 0→1): count becomes 1
- `a` == `a` (index 1→2): count becomes 2

But `"aaa"` is only **one** group. You need a flag to say "I am already inside a repeat — do not count again."

### The Solution: A `prevChar` and an `inRepeat` Flag

```go
func CountRepeats(s string) int {
    count := 0
    inRepeat := false
    var prevChar rune

    for _, c := range s {
        if c == prevChar {
            if !inRepeat {
                count++       // first time this repeat starts
                inRepeat = true
            }
        } else {
            inRepeat = false  // new character, end of any repeat
        }
        prevChar = c
    }
    return count
}
```

Step by step:
1. `prevChar` starts as the zero value for `rune` (`0`), which is not a printable character — safe as a sentinel
2. For each character `c`:
   - If `c == prevChar` and we are NOT already in a repeat: increment count, set `inRepeat = true`
   - If `c == prevChar` and we ARE in a repeat: do nothing (same group)
   - If `c != prevChar`: clear `inRepeat`, this group has ended
3. Update `prevChar = c` at the end of each iteration

### Tracing `"heelloo"`

| c | prevChar | c==prev? | inRepeat before | Action | count | inRepeat after |
|---|----------|----------|----------------|--------|-------|---------------|
| h | 0 | No | false | reset | 0 | false |
| e | h | No | false | reset | 0 | false |
| e | e | Yes | false | count++, set | 1 | true |
| l | e | No | true | reset | 1 | false |
| l | l | Yes | false | count++, set | 2 | true |
| o | l | No | true | reset | 2 | false |
| o | o | Yes | false | count++, set | 3 | true |

The trace ends at count 3: `ee`, `ll`, and `oo` are three separate groups. Any example that claims otherwise is wrong — trust the trace.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| No `inRepeat` flag | `"aaa"` counted as 2 groups instead of 1 | Track whether already in a repeat |
| Not updating `prevChar` | Always comparing to zero — never matches | `prevChar = c` at end of each iteration |
| Forgetting `var prevChar rune` declaration | Compile error | Declare before the loop |

## Solving This Challenge

### Algorithm

1. `count := 0`, `inRepeat := false`, `var prevChar rune`
2. For each rune `c` in `s`:
   - If `c == prevChar && !inRepeat`: `count++`, `inRepeat = true`
   - If `c != prevChar`: `inRepeat = false`
   - `prevChar = c`
3. Return `count`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [41-retainfirsthalf](../41-retainfirsthalf/README.md)
