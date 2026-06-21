# Skills for thirdtimeisacharm

## What You'll Learn

**Previous:** [../44-printrevcomb/skills.md](../44-printrevcomb/skills.md) | **Next:** [../46-weareunique/README.md](../46-weareunique/README.md)

**Challenge:** Return every third character from a string (positions 3, 6, 9, ...), followed by a newline.

## Core Concept: Selecting Every Nth Element Using Modulo

### What Is It?
Index every character and include only those at positions that are multiples of 3 (1-based counting: 3rd, 6th, 9th...). Using modulo on the loop index is the cleanest approach.

### How It Works

```go
func ThirdTimeIsACharm(str string) string {
    result := ""
    for i, c := range str {
        if (i+1)%3 == 0 {  // i is 0-based; (i+1) makes it 1-based
            result += string(c)
        }
    }
    return result + "\n"
}
```

Step-by-step for `"123456789"`:
- i=0 (char '1'): (0+1)%3 = 1 → skip
- i=1 (char '2'): (1+1)%3 = 2 → skip
- i=2 (char '3'): (2+1)%3 = 0 → include '3'
- i=3 (char '4'): (3+1)%3 = 1 → skip
- i=4 (char '5'): (4+1)%3 = 2 → skip
- i=5 (char '6'): (5+1)%3 = 0 → include '6'
- i=6 (char '7'): (6+1)%3 = 1 → skip
- i=7 (char '8'): (7+1)%3 = 2 → skip
- i=8 (char '9'): (8+1)%3 = 0 → include '9'
- Result: `"369\n"`

### Alternate: Index-Based Loop Stepping by 3
Start at index 2 (the 3rd character) and step by 3:
```go
func ThirdTimeIsACharm(str string) string {
    result := ""
    for i := 2; i < len(str); i += 3 {
        result += string(str[i])
    }
    return result + "\n"
}
```

### Handling Empty and Short Strings
- Empty string: loop doesn't run → `result = ""` → return `"\n"` ✓
- Fewer than 3 characters: no multiple of 3 found → return `"\n"` ✓
- Both cases handled automatically by the loop structure.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Using `i%3 == 0` (0-based) | Includes the 1st character (index 0) | Use `(i+1)%3 == 0` for 1-based counting |
| Using `i%3 == 2` instead of `(i+1)%3 == 0` | Works but confusing | Either is correct; be consistent |
| Not appending `"\n"` | Missing required newline | Append `+ "\n"` to the final return |

## Solving This Challenge

### Algorithm
1. Iterate over the string with `for i, c := range str`.
2. If `(i+1) % 3 == 0`, append `string(c)` to result.
3. Return `result + "\n"`.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../46-weareunique/README.md](../46-weareunique/README.md)
