# Skills for 58-itoa

**Previous:** [57-cameltosnakecase](../57-cameltosnakecase/skills.md) | **Next:** [59-thirdchar](../59-thirdchar/skills.md)

**Challenge:** Convert an integer to its string representation without using `strconv.Itoa` or `fmt.Sprintf`.

## Core Concept: Digit Extraction with Modulo and Division

### What Is It?

Every integer is made of digits. The key insight is that the **modulo operator** (`%`) extracts the rightmost digit, and **integer division** (`/`) removes it. By repeating this process, you extract all digits — but in reverse order, so you must reverse the result at the end.

### Extracting Digits Right-to-Left

```
n = 12345

Step 1: digit = 12345 % 10 = 5,  n = 12345 / 10 = 1234
Step 2: digit = 1234  % 10 = 4,  n = 1234  / 10 = 123
Step 3: digit = 123   % 10 = 3,  n = 123   / 10 = 12
Step 4: digit = 12    % 10 = 2,  n = 12    / 10 = 1
Step 5: digit = 1     % 10 = 1,  n = 1     / 10 = 0  ← loop ends
Digits collected: [5, 4, 3, 2, 1]
After reversal: [1, 2, 3, 4, 5]  → "12345"
```

### Converting a Digit to Its Character

A digit `0-9` becomes its character by adding it to the rune `'0'`:

```go
digit := 5
char := rune('0') + rune(digit)  // '0'=48, 48+5=53='5'
fmt.Println(string(char))        // "5"
```

This works because ASCII assigns '0'=48, '1'=49, ..., '9'=57 — consecutive values.

### The Full Implementation

```go
func Itoa(n int) string {
    if n == 0 {
        return "0"
    }

    negative := n < 0
    if negative {
        n = -n
    }

    digits := []rune{}
    for n > 0 {
        digit := n % 10
        digits = append(digits, rune('0')+rune(digit))
        n = n / 10
    }

    // Digits are in reverse order — reverse them
    for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
        digits[i], digits[j] = digits[j], digits[i]
    }

    result := string(digits)
    if negative {
        result = "-" + result
    }
    return result
}
```

### Why Zero Must Be a Special Case

The loop condition is `n > 0`. If `n == 0`, the loop body never runs, and you'd return `""` instead of `"0"`. Always handle zero explicitly:

```go
if n == 0 {
    return "0"
}
```

### Handling Negative Numbers

1. Record the sign: `negative := n < 0`
2. Make `n` positive: `n = -n`
3. Extract digits from the positive value
4. Prepend `"-"` if negative was true

### Why Must You Reverse?

The loop extracts digits right-to-left (units place first). To get the correct string representation, you must reverse the collected digits.

```
12345: extracted order = 5,4,3,2,1
       reversed         = 1,2,3,4,5 ✓
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not handling `n == 0` | Returns `""` instead of `"0"` | Check zero first |
| Not negating before loop | Negative numbers produce wrong digits | `n = -n` before the loop |
| Forgetting to reverse | `12345` becomes `"54321"` | Reverse the digit slice |
| Using `string(digit)` | `string(5)` is NOT `"5"` — it's the string with the rune at code point 5 | Use `string(rune('0') + rune(digit))` |

## Algorithm

1. If `n == 0`, return `"0"`
2. If `n < 0`, set `negative = true`, set `n = -n`
3. While `n > 0`: extract `digit = n % 10`, append `rune('0')+rune(digit)`, set `n = n/10`
4. Reverse the collected digits
5. If `negative`, prepend `"-"`
6. Return the result string

## The Challenge

See [README.md](README.md).

**Next:** [59-thirdchar](../59-thirdchar/README.md)
