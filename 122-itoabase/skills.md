# Skills for itoabase

## What You'll Learn

**Previous:** [121-wordflip](../121-wordflip/skills.md) | **Next:** [123-options](../123-options/skills.md)

**Challenge:** Convert an integer to its string representation in any base from 2 to 16, using digits `0-9` followed by uppercase `A-F`.

## Core Concept: Converting an Integer to an Arbitrary Base

### What Is It?

The standard decimal-to-binary (or decimal-to-hex) algorithm uses repeated division: divide the number by the base, collect the remainders in reverse order. The remainders give you the digits from least-significant to most-significant, so you reverse them at the end.

The digit set is `"0123456789ABCDEF"` — index 10 maps to `A`, 11 to `B`, etc.

### How It Works

```go
func ItoaBase(value, base int) string {
    digits := "0123456789ABCDEF"

    // Handle zero explicitly
    if value == 0 {
        return "0"
    }

    // Handle negative values
    negative := value < 0
    if negative {
        value = -value
    }

    // Repeated division — collect digits in reverse
    result := ""
    for value > 0 {
        remainder := value % base
        result = string(digits[remainder]) + result  // prepend
        value /= base
    }

    if negative {
        result = "-" + result
    }
    return result
}
```

**Trace `ItoaBase(255, 16)` (decimal 255 → hex FF):**

| value | value % 16 | digit | result (prepending) |
|-------|-----------|-------|---------------------|
| 255 | 15 | `F` | `"F"` |
| 15 | 15 | `F` | `"FF"` |
| 0 | — | — | loop ends |

Result: `"FF"` ✓

**Trace `ItoaBase(42, 2)` (decimal 42 → binary 101010):**

| value | value % 2 | digit | result |
|-------|-----------|-------|--------|
| 42 | 0 | `0` | `"0"` |
| 21 | 1 | `1` | `"10"` |
| 10 | 0 | `0` | `"010"` |
| 5 | 1 | `1` | `"1010"` |
| 2 | 0 | `0` | `"01010"` |
| 1 | 1 | `1` | `"101010"` |

Result: `"101010"` ✓

**Trace `ItoaBase(-10, 16)` → `"-A"`:**
- Negative → work with 10
- 10 % 16 = 10 → `digits[10]` = `A`
- Prepend `-`
- Result: `"-A"` ✓

**Why prepend instead of append then reverse?**

Both work. Prepending with string concatenation is simpler to read. For performance on large numbers, you could use a `[]byte` buffer and reverse at the end:

```go
buf := []byte{}
for value > 0 {
    buf = append(buf, digits[value%base])
    value /= base
}
// reverse buf
for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
    buf[i], buf[j] = buf[j], buf[i]
}
return string(buf)
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not handling `value == 0` | The loop never executes, returns `""` | Return `"0"` before the loop |
| Using lowercase `abcdef` | Wrong output — challenge requires uppercase | Use `"0123456789ABCDEF"` |
| Appending instead of prepending | Digits come out in wrong order | Either prepend with `string(...) + result` or append then reverse |
| Not handling negative values | Returns wrong result or panics | Detect `value < 0`, work with absolute value, prepend `"-"` |

## Solving This Challenge

### Algorithm
1. If `value == 0`, return `"0"`.
2. If `value < 0`, set `negative = true`, `value = -value`.
3. While `value > 0`: prepend `digits[value % base]` to result; `value /= base`.
4. If negative, prepend `"-"`.
5. Return result.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [123-options](../123-options/README.md)
