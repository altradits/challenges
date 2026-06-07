# Skills for 01-only1

## What You'll Learn

This is your **first exercise** in Go programming. By the end, you'll understand:
- How Go programs are structured
- How to print output to the console
- The basics of the `main` package and `main` function
- How to run a Go program

## Prerequisites

Before starting this exercise, you should know:

### 1. How to create a basic Go program

```go
package main

import "fmt"

func main() {
    // Your code here
}
```

**Explanation:**
- `package main` - Declares this as an executable program (not a library)
- `import "fmt"` - Imports the "format" package for printing
- `func main()` - The entry point where execution begins

### 2. How to print to the console

```go
fmt.Println("Hello")
```

**Explanation:**
- `fmt.Println()` prints text followed by a newline
- The text must be inside double quotes `""`

### 3. How to run a Go program

```bash
go run .
```

**Explanation:**
- `go run` compiles and runs your program
- `.` means "run the program in the current directory"

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Write a complete Go program** from scratch
2. **Use `fmt.Println()`** to output text
3. **Understand program structure** (package, imports, main function)
4. **Run Go programs** using the command line

## How This Helps Your Capstone

This skill is used in:
- **All Projects** - Every program needs a main function and output
- **Budget Planner** - Display results to users
- **Savings Calculator** - Show calculated values
- **Investment Tracker** - Print portfolio summaries
- **Net Worth Tracker** - Output financial reports

## Step-by-Step Approach

1. **Create the package declaration** - `package main`
2. **Import the fmt package** - `import "fmt"`
3. **Create the main function** - `func main() { }`
4. **Add the print statement** - `fmt.Println("1")`
5. **Run the program** - `go run .`

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Forgetting `package main` | Program won't compile | Always start with `package main` |
| Missing `import "fmt"` | `fmt` won't be recognized | Import packages you use |
| Using `print` instead of `Println` | Go is case-sensitive | Use exact function names |
| Forgetting quotes around text | Syntax error | `"1"` not `1` |

## Practice Tips

- Type the program character by character to build muscle memory
- Read error messages carefully - they tell you exactly what's wrong
- Run `go run .` after every small change to test

## The Challenge

Write a program that displays a `1` character on the standard output and nothing else.

### Expected Program

```go
package main

import "fmt"

func main() {
    fmt.Println("1")
}
```

### Expected Output

```
1
```

### Usage Example

```bash
$ go run .
1
$ go run . "a" "b"
1
$ go run . "a" "b" "c"
1
```

## Knowledge Check

Before coding, make sure you can answer:
1. What does `package main` mean?
2. Why do we need `import "fmt"`?
3. What is the purpose of `func main()`?

## Next Steps

After completing this, you'll be ready for:
- **02-onlya**: Printing a different character
- **03-onlyb**: Printing another character
- **04-onlyf**: Continuing with character output
- **05-onlyz**: Mastering basic output

---

## Building Your Skills

### Level 1: Absolute Basics (01-05)
These exercises teach you the fundamentals of Go program structure and output.

| Exercise | Skill | What You Practice |
|----------|-------|-------------------|
| 01-only1 | Basic output | `fmt.Println("1")` |
| 02-onlya | Character output | `fmt.Println("a")` |
| 03-onlyb | More output | `fmt.Println("B")` |
| 04-onlyf | Continued practice | `fmt.Println("f")` |
| 05-onlyz | Mastery of basics | `fmt.Println("z")` |

### Level 2: Conditional Logic (06-13)
After mastering output, you'll learn:
- Checking conditions with `if` statements
- Using boolean logic
- Printing conditionally

### Level 3: String Operations (15-19)
Then you'll move to:
- String manipulation
- Counting characters
- Reversing strings

### Level 4: Advanced Patterns (20-33)
Building on basics:
- Pattern matching
- Word analysis
- Complex string processing

## Your Learning Journey

```
01-only1 (You are here)
    ↓
02-onlya → 03-onlyb → 04-onlyf → 05-onlyz
    ↓
06-checknumber → 07-countalpha → 08-count-character
    ↓
...and so on through 138 challenges
```

**Remember:** Every expert was once a beginner. Take your time, understand each concept, and don't skip ahead. The skills build on each other!
