# Skills for notdecimal

## What You'll Learn

**Previous:** [129-fifthandskip](../129-fifthandskip/skills.md) | **Next:** [131-revconcatalternate](../131-revconcatalternate/skills.md)

**Challenge:** Given a float string like `"174.2"`, multiply by 10^n to remove the decimal point and return the integer string `"1742"`. If not a valid decimal with non-zero fractional part, return the input as-is (with newline).

## Core Concept: String-Based Decimal Manipulation

### What Is It?

Instead of using floating-point arithmetic (which has precision issues), this challenge works directly on the string representation. Find the decimal point, count how many significant digits follow it, remove the dot, and trim trailing zeros.

### How It Works

**Step 1 — Handle edge cases:**

```go
func NotDecimal(dec string) string {
    if dec == "" {
        return "\n"
    }
```

**Step 2 — Find the decimal point:**

```go
    dotIdx := -1
    for i, c := range dec {
        if c == '.' {
            dotIdx = i
            break
        }
    }
    // No dot, or not a valid number → return as-is
    if dotIdx == -1 {
        return dec + "\n"
    }
```

**Step 3 — Validate the input is a number:**

Check that all non-dot characters are digits (allowing a leading `-`).

```go
    for i, c := range dec {
        if i == 0 && c == '-' {
            continue
        }
        if c == '.' {
            continue
        }
        if c < '0' || c > '9' {
            return dec + "\n"  // contains non-digit, non-dot: not a number
        }
    }
```

**Step 4 — Build the integer string by removing the dot:**

```go
    // Integer part + fractional part without the dot
    intPart := dec[:dotIdx]
    fracPart := dec[dotIdx+1:]

    // Trim trailing zeros from fracPart
    end := len(fracPart)
    for end > 0 && fracPart[end-1] == '0' {
        end--
    }
    fracPart = fracPart[:end]

    if fracPart == "" {
        // Only zeros after dot (or no fractional part)
        return intPart + "\n"
    }
    return intPart + fracPart + "\n"
}
```

**Examples:**

| Input | dotIdx | intPart | fracPart (trimmed) | Output |
|-------|--------|---------|-------------------|--------|
| `"0.1"` | 1 | `"0"` | `"1"` | `"1\n"` |
| `"174.2"` | 3 | `"174"` | `"2"` | `"1742\n"` |
| `"0.1255"` | 1 | `"0"` | `"1255"` | `"1255\n"` |
| `"-19.525856"` | 3 | `"-19"` | `"525856"` | `"-19525856\n"` |
| `"1952"` | -1 | — | — | `"1952\n"` |
| `"-0.0f00d00"` | 2 | — | invalid char 'f' | `"-0.0f00d00\n"` |

**Why string manipulation instead of `strconv.ParseFloat`?**

`strconv.ParseFloat("1.20525856", 64)` would give floating-point imprecision. String manipulation is exact.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `float64` arithmetic | Precision loss: `1.20525856 * 10^8` may not be exactly `120525856` | Work on the string directly |
| Not trimming trailing zeros | `"174.20"` → `"174200"` instead of `"17420"` | Trim zeros from fracPart before concatenating |
| Forgetting the leading `-` during validation | Rejects negative numbers | Allow `-` at index 0 |

## Solving This Challenge

### Algorithm
1. Return `"\n"` for empty input.
2. Validate: all chars must be digits, one optional leading `-`, one optional `.`.
3. If no `.`, return `dec + "\n"`.
4. Split at `.`: `intPart = dec[:dotIdx]`, `fracPart = dec[dotIdx+1:]`.
5. Trim trailing zeros from `fracPart`.
6. If `fracPart` is empty after trimming, return `intPart + "\n"`.
7. Return `intPart + fracPart + "\n"`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [131-revconcatalternate](../131-revconcatalternate/README.md)
