# Skills for romannumbers

## What You'll Learn

**Previous:** [70-piglatin](../70-piglatin/skills.md) | **Next:** [72-brackets](../72-brackets/README.md)

**Challenge:** Convert an integer (1–3999) to Roman numerals, printing both the additive/subtractive breakdown and the final numeral.

## Core Concept: Greedy Subtraction with a Value Table

### What Is It?

Roman numeral conversion uses a **greedy algorithm**: find the largest Roman numeral value that fits into the remaining number, subtract it, and repeat. The key insight is including the **subtractive pairs** (CM=900, CD=400, XC=90, XL=40, IX=9, IV=4) in the value table.

### How It Works

**Step 1 — Define the value table (largest to smallest, including subtractive pairs):**

```go
type romanPair struct {
    value  int
    symbol string
    prefix string  // for the breakdown notation
}

var romanTable = []romanPair{
    {1000, "M", "M"},
    {900, "CM", "(M-C)"},
    {500, "D", "D"},
    {400, "CD", "(D-C)"},  // wait, should be (C-D)? Check example
    {100, "C", "C"},
    {90, "XC", "(C-X)"},
    {50, "L", "L"},
    {40, "XL", "(L-X)"},   // check actual notation from example
    {10, "X", "X"},
    {9, "IX", "(X-I)"},
    {5, "V", "V"},
    {4, "IV", "(V-I)"},    // check
    {1, "I", "I"},
}
```

Looking at the example output for 999: `(M-C)+(C-X)+(X-I)` → `CMXCIX`. So CM → `(M-C)`, XC → `(C-X)`, IX → `(X-I)`. The notation shows (larger - smaller) even though the smaller comes first in the numeral. Let me define the table correctly:

```go
var vals = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var syms = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var breakdown = []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
```

Wait, checking: `CD` is 400 = 500-100 = D-C → `(D-C)`. But `CM` is 900 = 1000-100 = M-C → `(M-C)`. Let me verify with example `3999`: `M+M+M+(M-C)+(C-X)+(X-I)` = 3000 + 900 + 90 + 9 = 3999. Yes.

**Step 2 — Greedy algorithm:**

```go
func toRoman(n int) (string, string) {
    var symbols []string     // for the final numeral
    var breakdowns []string  // for the breakdown line

    for i, v := range vals {
        for n >= v {
            symbols = append(symbols, syms[i])
            breakdowns = append(breakdowns, breakdown[i])
            n -= v
        }
    }
    return strings.Join(breakdowns, "+"), strings.Join(symbols, "")
}
```

**Trace 123:**
- 100 fits → breakdown `C`, symbol `C`, n=23
- 10 fits → breakdown `X`, symbol `X`, n=13
- 10 fits → breakdown `X`, symbol `X`, n=3
- 1 fits × 3 → breakdown `I`, `I`, `I`, symbol `I`, `I`, `I`
- Breakdown: `C+X+X+I+I+I`, Symbol: `CXXIII` ✓

**Trace 999:**
- 900 fits → `(M-C)`, `CM`, n=99
- 90 fits → `(C-X)`, `XC`, n=9
- 9 fits → `(X-I)`, `IX`, n=0
- Breakdown: `(M-C)+(C-X)+(X-I)`, Symbol: `CMXCIX` ✓

**Main function:**

```go
func main() {
    if len(os.Args) != 2 {
        return
    }
    n, err := strconv.Atoi(os.Args[1])
    if err != nil || n <= 0 || n >= 4000 {
        fmt.Println("ERROR: cannot convert to roman digit")
        return
    }
    bd, sym := toRoman(n)
    fmt.Println(bd)
    fmt.Println(sym)
}
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Missing subtractive pairs (900, 400, 90, 40, 9, 4) | Greedy picks wrong symbols | Include all 13 pairs in the table |
| Table not sorted largest-to-smallest | Greedy picks wrong value first | Always sort descending |
| Accepting `n >= 4000` | Challenge limits to 3999 | `if n <= 0 || n >= 4000` → error |

## Solving This Challenge

### Algorithm
1. Parse `n` from args; validate `1 ≤ n ≤ 3999`.
2. For each `(value, symbol, breakdown_notation)` in the table (largest first): while `n >= value`, collect symbols and breakdown labels, subtract value.
3. Print breakdown labels joined by `+`.
4. Print symbols concatenated.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [72-brackets](../72-brackets/README.md)
