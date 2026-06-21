# Skills for hashcode

## What You'll Learn

**Previous:** [81-gcd](../81-gcd/skills.md) | **Next:** [83-lastword](../83-lastword/skills.md)

**Challenge:** Transform every character in a string using a hash formula: `(ASCII value + string length) % 127`, with an adjustment for unprintable results.

## Core Concept: ASCII Arithmetic on Characters

### What Is ASCII?
Every character has a numeric code. ASCII maps characters 0–127 to numbers. In Go, you can get this number with `int(c)` and convert back with `string(rune(val))`.

```go
fmt.Println(int('A'))          // 65
fmt.Println(int('B'))          // 66
fmt.Println(string(rune(66)))  // "B"
```

### The Hash Formula
For each character `c` in a string of length `n`:
1. `hashed = (int(c) + n) % 127`
2. If `hashed < 32` (control/unprintable character), add 33 to shift it into the printable range.
3. Append `string(rune(hashed))` to the result.

```go
func HashCode(dec string) string {
    n := len(dec)
    result := ""
    for _, c := range dec {
        hashed := (int(c) + n) % 127
        if hashed < 32 {
            hashed += 33
        }
        result += string(rune(hashed))
    }
    return result
}
```

### Step-by-Step for `HashCode("A")`
- `n = 1`, `c = 'A'` → `int('A') = 65`
- `(65 + 1) % 127 = 66`
- `66 >= 32` so no adjustment
- `string(rune(66)) = "B"`
- Output: `"B"`

### Step-by-Step for `HashCode("AB")`
- `n = 2`
- `'A'=65`: `(65+2)%127 = 67` → `"C"`
- `'B'=66`: `(66+2)%127 = 68` → `"D"`
- Output: `"CD"`

### Why `% 127`?
Keeps the value in the range `[0, 126]`. ASCII printable characters start at 32 (space) and go to 126 (`~`). Values below 32 are control characters (tab, newline, etc.), which is why we add 33 to push them into printable range (33 = `!`).

### Converting Between rune, int, and string
```go
c := 'H'                  // rune — value 72
val := int(c) + 5         // int  — value 77
ch := string(rune(val))   // string — "M"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting `int(c)` before arithmetic | Rune arithmetic can overflow byte type | Cast explicitly: `int(c)` |
| Not checking `hashed < 32` | Unprintable control characters appear in output | Add `if hashed < 32 { hashed += 33 }` |
| Using string index `s[i]` instead of range | Gets bytes not runes | Use `for _, c := range dec` |

## Solving This Challenge

### Algorithm
1. Get `n := len(dec)`.
2. For each character `c` using `for _, c := range dec`:
   a. `hashed := (int(c) + n) % 127`
   b. If `hashed < 32`: `hashed += 33`
   c. Append `string(rune(hashed))` to result.
3. Return result.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [83-lastword](../83-lastword/README.md)
