# Skills for digitlen

## What You'll Learn

**Previous:** [../21-countrepeats/skills.md](../21-countrepeats/skills.md) | **Next:** [../23-firstword/README.md](../23-firstword/README.md)

**Challenge:** Count how many times `n` can be divided by `base` until it reaches zero — effectively counting the number of digits in that base.

## Core Concept: Repeated Integer Division

### What Is It?
The number of digits a positive integer `n` has in base `b` equals the number of times you can divide it by `b` before it hits zero. This is pure integer arithmetic — no string conversion needed.

For example, `100` in base 10: 100 → 10 → 1 → 0 = 3 steps = 3 digits.

### How It Works

```go
func DigitLen(n, base int) int {
    if base < 2 || base > 36 {
        return -1
    }
    if n < 0 {
        n = -n
    }
    count := 0
    for n > 0 {
        n /= base
        count++
    }
    return count
}
```

Step-by-step for `DigitLen(100, 10)`:
- n=100: 100/10=10, count=1
- n=10:  10/10=1,  count=2
- n=1:   1/10=0,   count=3
- n=0: loop ends
- return 3

Step-by-step for `DigitLen(100, 2)`:
- 100→50→25→12→6→3→1→0 (7 steps) → return 7

### Integer Division in Go
The `/` operator on integers truncates toward zero — no decimal part:
```go
fmt.Println(7 / 2)    // 3, not 3.5
fmt.Println(1 / 10)   // 0  ← this ends the loop
fmt.Println(-7 / 2)   // -3 (truncated toward zero)
```

### Handling Negative Input
The challenge says treat negative numbers by using their absolute value:
```go
if n < 0 {
    n = -n
}
```

### Validating the Base
The base must be between 2 and 36. Check this before doing anything else:
```go
if base < 2 || base > 36 {
    return -1
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not validating base first | Wrong results for base 0 or 1 (infinite loop for base 1) | Check `base < 2 \|\| base > 36` at the start |
| Not negating negative `n` | Loop `n > 0` never runs for negative numbers | Add `if n < 0 { n = -n }` before the loop |
| Using `n /= base` with float | Precision errors | Use integer types only — no `float64` |

## Solving This Challenge

### Algorithm
1. If `base < 2` or `base > 36`, return `-1`.
2. If `n < 0`, set `n = -n`.
3. Initialize `count := 0`.
4. While `n > 0`: divide `n` by `base`, increment `count`.
5. Return `count`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../23-firstword/README.md](../23-firstword/README.md)
