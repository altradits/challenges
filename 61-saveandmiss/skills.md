# Skills for 61-saveandmiss

**Previous:** [60-zipstring](../60-zipstring/skills.md) | **Next:** [62-reversestrcap](../62-reversestrcap/skills.md)

**Challenge:** Given a string `s` and integer `n`, keep the first `n` characters of every `2n`-character chunk and discard the next `n`. Return the kept characters concatenated.

## Core Concept: Cyclic Position Tracking with a Counter

### What Is It?

A **cyclic position tracker** resets to 0 after completing each full cycle. This is the same idea as a clock: after 12 comes 1 again (not 13). Here, each cycle is `2n` characters long: the first `n` are saved, the next `n` are missed, then the cycle repeats.

### The Pattern

```
n=3, string="123456789"

Cycle:    SAVE SAVE SAVE | MISS MISS MISS | SAVE SAVE SAVE
char:       1    2    3  |  4    5    6   |  7    8    9
pos:        0    1    2  |  3    4    5   |  0    1    2
                              (pos resets to 0 after reaching 2n=6)
Result: "123789"
```

### The Implementation

```go
func SaveAndMiss(s string, n int) string {
    if n <= 0 {
        return s
    }
    result := ""
    pos := 0
    for _, c := range s {
        if pos < n {
            result += string(c)  // in the SAVE zone
        }
        // pos >= n: in the MISS zone, skip character
        pos++
        if pos == 2*n {
            pos = 0  // completed one full cycle, reset
        }
    }
    return result
}
```

### Why Not Use Modulo Directly?

You could write `if i%(2*n) < n { save }`, but the manual counter is clearer and avoids recalculating `2*n` at every iteration. Both approaches are correct.

Using modulo:
```go
for i, c := range s {
    if i%(2*n) < n {
        result += string(c)
    }
}
```

### Tracing `"abcdefghijklmno"` with n=3

Cycle length = 6. Save positions 0,1,2. Miss positions 3,4,5.

| i | char | i%(2*3)=i%6 | < 3? | result |
|---|------|-------------|------|--------|
| 0 | a | 0 | yes | "a" |
| 1 | b | 1 | yes | "ab" |
| 2 | c | 2 | yes | "abc" |
| 3 | d | 3 | no | "abc" |
| 4 | e | 4 | no | "abc" |
| 5 | f | 5 | no | "abc" |
| 6 | g | 0 | yes | "abcg" |
| 7 | h | 1 | yes | "abcgh" |
| 8 | i | 2 | yes | "abcghi" |

### Edge Cases

- `n <= 0`: The cycle length would be 0 or negative. Return the original string unchanged.
- Empty string: The loop never runs; returns `""`.
- String shorter than `n`: All characters are in the save zone; returns the whole string.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `<= n` instead of `< n` | Saves `n+1` characters per cycle instead of `n` | Use strict `< n` |
| Not resetting pos at `2*n` | pos keeps growing forever | Reset `pos = 0` when `pos == 2*n` |
| Not handling `n <= 0` | Division by zero or weird behavior | Return original string if `n <= 0` |

## Algorithm

1. If `n <= 0`, return `s`
2. Set `result = ""`, `pos = 0`
3. For each rune `c` in `s`:
   - If `pos < n`: append `string(c)` to result
   - Increment `pos`
   - If `pos == 2*n`: set `pos = 0`
4. Return `result`

## The Challenge

See [README.md](README.md).

**Next:** [62-reversestrcap](../62-reversestrcap/README.md)
