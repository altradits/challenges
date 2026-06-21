# Skills for itoa-35

## What You'll Learn

**Previous:** [../40-iscapitalized/skills.md](../40-iscapitalized/skills.md) | **Next:** [../42-arrays/README.md](../42-arrays/README.md)

**Challenge:** Convert an integer to its string representation without using `strconv.Itoa` — implement the conversion manually.

## Core Concept: Manual Integer-to-String Conversion

### What Is It?
`strconv.Itoa` in the standard library converts an `int` to its decimal string. Here you implement that logic yourself using integer arithmetic: repeatedly take `n % 10` to get the last digit, prepend it to the result, then divide `n` by 10.

### How It Works

```go
func Itoa(n int) string {
    if n == 0 {
        return "0"
    }
    result := ""
    negative := false
    if n < 0 {
        negative = true
        n = -n
    }
    for n > 0 {
        digit := n % 10          // last digit: 0-9
        result = string(rune('0'+digit)) + result  // prepend
        n /= 10
    }
    if negative {
        result = "-" + result
    }
    return result
}
```

### How Digit Extraction Works
`n % 10` gives the ones digit. `n / 10` removes it. Repeating extracts all digits from right to left, so you prepend each one:

For `n = 12345`:
- 12345 % 10 = 5, prepend → "5";  12345/10 = 1234
- 1234 % 10 = 4, prepend → "45"; 1234/10 = 123
- 123 % 10 = 3, prepend → "345"; 123/10 = 12
- 12 % 10 = 2, prepend → "2345"; 12/10 = 1
- 1 % 10 = 1, prepend → "12345"; 1/10 = 0
- Loop ends. Result: `"12345"`

### Converting a Digit to a Character
```go
digit := 5
ch := string(rune('0' + digit)) // '0' is 48, +5 = 53 = '5'
```

`'0' + digit` gives the ASCII code of that digit character. `string(rune(...))` converts it to a one-character string.

### Handling Negatives
Detect negative, negate, convert the positive part, then prepend `"-"`:
```go
if n < 0 {
    negative = true
    n = -n
}
```

### Handling Zero
The loop `for n > 0` never runs when `n == 0`. Special-case it first:
```go
if n == 0 {
    return "0"
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not handling `n == 0` | Loop never runs, returns empty string | Return `"0"` before the loop |
| Appending instead of prepending | Digits come out in reverse order | Use `result = newDigit + result` |
| Not handling negative numbers | Negative input returns wrong result | Check `n < 0`, negate, add `"-"` at end |

## Solving This Challenge

### Algorithm
1. If `n == 0`, return `"0"`.
2. If `n < 0`, set `negative = true`, negate `n`.
3. While `n > 0`: prepend `string(rune('0' + n%10))` to result, then `n /= 10`.
4. If `negative`, prepend `"-"`.
5. Return result.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../42-arrays/README.md](../42-arrays/README.md)
