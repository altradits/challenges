# Skills for 05-onlyz

## What You'll Learn

**Previous:** [04-onlyf](../04-onlyf/skills.md) | **Next:** [07-functions](../07-functions/skills.md)

**Challenge:** Write a program that prints the character `z` and nothing else.

## Core Concept: Consolidating the Basic Program Pattern

### What Is It?

This is the last of the single-character print challenges. By now the pattern should be automatic. This challenge is about building *muscle memory* — being able to write a correct Go program from scratch without referring to notes.

### The Pattern You Know

```go
package main

import "fmt"

func main() {
    fmt.Println("z")
}
```

Output:
```
z
```

### What You Have Learned Across Challenges 01-05

| Challenge | Printed | Key Learning |
|-----------|---------|--------------|
| 01-only1 | `1` | `package main`, `import "fmt"`, `func main()`, `fmt.Println()` |
| 02-onlya | `a` | String literals use double quotes; `'a'` is a rune (int), not a string |
| 03-onlyb | `B` | Go is case-sensitive in both keywords and string values |
| 04-onlyf | `f` | `fmt.Println` vs `fmt.Print` vs `fmt.Printf` |
| 05-onlyz | `z` | Practice — write the full structure from memory |

### Self-Check: Write From Memory

Before looking at any reference, try to write this program from scratch:

1. Open a blank file
2. Write the package declaration
3. Write the import
4. Write the main function
5. Add the print statement for `z`
6. Run it with `go run .`

If you can do this without copying, you are ready for challenge 06 which introduces real logic.

### What Comes Next

Starting with challenge 06, programs will:
- Accept command-line arguments (`os.Args`)
- Convert strings to integers (`strconv.Atoi`)
- Make decisions (`if/else`)
- Use arithmetic operators (`%`, `>=`, etc.)

The simple skeleton you have practiced five times is the foundation for all of that.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `fmt.Println("Z")` uppercase | Prints `Z`, challenge expects lowercase `z` | Use `"z"` |
| Still copying from previous file | You are not building memory | Close the file, write from scratch |

## Solving This Challenge

### Algorithm

1. Declare `package main`
2. Import `fmt`
3. Inside `func main()`, call `fmt.Println("z")`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [07-functions](../07-functions/README.md)
