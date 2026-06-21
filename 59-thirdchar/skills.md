# Skills for 59-thirdchar

**Previous:** [58-itoa](../58-itoa/skills.md) | **Next:** [60-zipstring](../60-zipstring/skills.md)

**Challenge:** Return a string containing every 3rd character of the input (positions 3, 6, 9, ..., using 1-based counting), followed by a newline.

## Core Concept: Modulo Arithmetic for Periodic Selection

### What Is It?

The **modulo operator** `%` tells you the remainder after division. When you iterate with an index, `(i+1) % 3 == 0` is true exactly at positions that are multiples of 3 (1-based): position 3, 6, 9, etc.

```
Index (0-based):  0  1  2  3  4  5  6  7  8
1-based position: 1  2  3  4  5  6  7  8  9
(i+1) % 3:        1  2  0  1  2  0  1  2  0
Select when 0:          ↑           ↑           ↑
```

### The Implementation

```go
func ThirdChar(s string) string {
    result := ""
    for i, c := range s {
        if (i+1)%3 == 0 {
            result += string(c)
        }
    }
    return result + "\n"
}
```

### Why `(i+1) % 3 == 0` and Not `i % 3 == 0`?

`i` is 0-based. Position 3 (1-based) is index 2 (0-based).
- `i % 3 == 0` is true at i=0,3,6 → 1st, 4th, 7th characters (not what we want)
- `(i+1) % 3 == 0` is true at i=2,5,8 → 3rd, 6th, 9th characters (correct)

Alternatively: `i % 3 == 2` gives the same result:

```go
if i%3 == 2 {  // same as (i+1)%3 == 0
    result += string(c)
}
```

### Tracing Through `"123456789"`

| i | char | (i+1)%3 | Select? | result |
|---|------|---------|---------|--------|
| 0 | '1' | 1 | no | "" |
| 1 | '2' | 2 | no | "" |
| 2 | '3' | 0 | yes | "3" |
| 3 | '4' | 1 | no | "3" |
| 4 | '5' | 2 | no | "3" |
| 5 | '6' | 0 | yes | "36" |
| 6 | '7' | 1 | no | "36" |
| 7 | '8' | 2 | no | "36" |
| 8 | '9' | 0 | yes | "369" |

Return `"369\n"` — correct.

### The Trailing Newline

This function always returns its result with a `"\n"` at the end. Even for an empty string, it returns `"\n"`. This is a specific requirement of the challenge — don't forget it.

### Generalization: Every Nth Character

The same pattern selects every Nth character:

```go
// Every 4th character (positions 4, 8, 12, ...):
if (i+1)%4 == 0 { ... }

// Every 2nd character (positions 2, 4, 6, ...):
if (i+1)%2 == 0 { ... }
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `i % 3 == 0` | Selects 1st, 4th, 7th (off by 2) | Use `(i+1) % 3 == 0` or `i % 3 == 2` |
| Forgetting the `\n` suffix | Output does not match expected | Always append `"\n"` after the loop |
| Using `s[i]` instead of `c` | Works for ASCII but unsafe for Unicode | Use `c` from `for i, c := range s` |

## Algorithm

1. Set `result = ""`
2. For `i, c := range s`:
   - If `(i+1) % 3 == 0`: append `string(c)` to result
3. Return `result + "\n"`

## The Challenge

See [README.md](README.md).

**Next:** [60-zipstring](../60-zipstring/README.md)
