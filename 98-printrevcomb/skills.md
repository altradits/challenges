# Skills for printrevcomb

## What You'll Learn

**Previous:** [97-printmemory](../97-printmemory/skills.md) | **Next:** [99-thirdtimeisacharm](../99-thirdtimeisacharm/skills.md)

**Challenge:** Print all unique combinations of three different digits in descending order, separated by `", "`.

## Core Concept: Nested Loops for Combinations

### What Is It?
A combination is a selection of items where order doesn't matter and each item is unique. For three digits a > b > c (where a, b, c are all different), use three nested loops with constraints.

### The Nested Loop Structure
The outer loop picks the largest digit, the middle picks the second, the inner picks the smallest. Each loop's range ensures the digits are strictly decreasing:

```go
import "fmt"

func main() {
    first := true
    for a := 9; a >= 2; a-- {
        for b := a - 1; b >= 1; b-- {
            for c := b - 1; c >= 0; c-- {
                if !first {
                    fmt.Print(", ")
                }
                fmt.Printf("%d%d%d", a, b, c)
                first = false
            }
        }
    }
    fmt.Println()
}
```

### Why a >= 2, b >= 1, c >= 0?
- `a` is the largest digit. It needs room for `b` and `c` below it, so minimum is 2.
- `b` is the middle digit. It needs room for `c` below it, so minimum is 1.
- `c` is the smallest digit, minimum is 0.

### Why b = a-1 and c = b-1?
Start each inner loop one below the outer loop's value to guarantee strictly decreasing digits. For example, when `a=9`, `b` starts at 8 and goes down to 1.

### Handling the Comma-Separator
The last combination should not have a trailing comma. Use a `first` boolean flag:
```go
first := true
// inside the innermost loop:
if !first {
    fmt.Print(", ")
}
fmt.Printf("%d%d%d", a, b, c)
first = false
```

Alternatively, track position and only print the comma before combinations that aren't the first.

### Step-by-Step (First Few)
- a=9, b=8, c=7 → print `987`
- a=9, b=8, c=6 → print `, 986`
- a=9, b=8, c=5 → print `, 985`
- ...
- a=3, b=1, c=0 → print `, 310`
- a=2, b=1, c=0 → print `, 210`
- Print newline

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `b := a; c := b` (not `a-1, b-1`) | Includes cases where digits repeat | Start inner loops at one below the outer |
| Printing a trailing comma | Output ends with `, ` | Use a `first` flag or check if it's the last combination |
| Printing with `Println` inside the loop | Adds extra newlines | Use `fmt.Print` inside, single `fmt.Println()` at the end |

## Solving This Challenge

### Algorithm
1. Use three nested loops: `a` from 9 down to 2, `b` from `a-1` down to 1, `c` from `b-1` down to 0.
2. Before printing each combination, print `", "` unless it's the first.
3. Print each combination as `"%d%d%d"`.
4. After all loops, print a newline.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [99-thirdtimeisacharm](../99-thirdtimeisacharm/README.md)
