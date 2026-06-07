# Skills for 131-savings-runway-calculator

## What You'll Learn

**Previous:** [130-index-fund-tracker](../130-index-fund-tracker/skills.md)

If you're stuck, review the previous exercise's skills.md to strengthen your foundation.

**Challenge:** Savings Runway Calculator

## Prerequisite Knowledge

Before starting, you should understand:
- Basic Go syntax and program structure
- How to define and call functions
- String manipulation basics in Go
- Control flow (if/else, loops)

## New Concepts Explained

### 1. String manipulation and processing

In Go, strings are immutable sequences of bytes encoded in UTF-8. You can iterate over them using `for...range` which gives you runes (Unicode code points) rather than bytes.

```go
for _, char := range myString {
    // char is a rune (int32)
}
```

To build new strings, concatenate using `+` or use `strings.Builder` for efficiency.

### 2. Go function definition and usage

Functions in Go are defined using the `func` keyword. They can take parameters and return values:

```go
func FunctionName(param1 type1, param2 type2) returnType {
    // function body
    return result
}
```

The `main()` function is special - it's where program execution begins.

### 3. Looping constructs (for, range)

Go has only one looping construct: the `for` loop. It can be used in several ways:

```go
// Traditional for loop
for i := 0; i < 10; i++ { }

// While-style loop
for condition { }

// Range loop (for collections)
for index, value := range collection { }
```

For strings, `for...range` iterates over runes, making it safe for UTF-8.

### 4. Conditional logic and boolean returns

Go uses `if/else` for conditional branching. The condition doesn't need parentheses:

```go
if condition {
    // do something
} else if otherCondition {
    // do something else
} else {
    // default case
}
```

Boolean operators: `&&` (AND), `||` (OR), `!` (NOT).

## Complete Solution Example

```go
package main

import "fmt"

ParseExpense implementation here

func main() {
    // Test your function
    fmt.Println(ParseExpense())
}
```

### Line-by-Line Explanation

1. `package main` - Declares this as an executable program
2. `import "fmt"` - Imports the formatting package
3. Function implementation follows the expected signature
4. `main()` - Tests the function with sample inputs

## How to Run Your Program

1. Save the file as `main.go`
2. Open a terminal in the same directory
3. Run: `go run main.go`
4. Verify the output matches expected results

## Skills You'll Learn

After completing this exercise, you'll be able to:
1. Apply string manipulation and processing to solve challenges
2. Build on previous skills without repeating them
3. Progress to the next challenge with confidence

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Text processing and validation
- **Savings Calculator** - Input handling and formatting
- **Investment Tracker** - Data parsing and analysis
- **Net Worth Tracker** - String manipulation for reports

## Step-by-Step Approach

1. **Review previous skills** if needed - click the link above
2. **Understand the challenge** - Read the README carefully
3. **Plan your solution** - Sketch the logic before coding
4. **Implement incrementally** - Build and test piece by piece
5. **Verify edge cases** - Test with empty, single, and boundary inputs

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Skipping prerequisites | You'll miss foundational concepts | Review previous skills.md first |
| Overcomplicating | Simple problems have simple solutions | Start with the simplest approach |
| Off-by-one errors | Index math is tricky | Draw it out on paper |
| Wrong return type | Type mismatch errors | Check function signature |

## Practice Tips

- **Draw it out**: Visualize the problem with diagrams
- **Trace manually**: Follow your code step by step with sample input
- **Test boundaries**: Empty string, single char, max values
- **Review previous**: If stuck, go back to the previous skills.md

## The Challenge

See [README.md](README.md) for the full challenge description, expected function, and test cases.

## Knowledge Check

Before coding, make sure you can answer:
1. What new concept does this exercise introduce?
2. Which previous skill does this build on? (Click the link above)
3. What edge cases should I test?

## Reference Links

For continued learning, explore these resources:

- [Go Documentation](https://go.dev/doc/)
- [Go Package Documentation](https://pkg.go.dev/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://go.dev/doc/effective_go)
- [strings package](https://pkg.go.dev/strings)
- [math package](https://pkg.go.dev/math)
- [regexp package](https://pkg.go.dev/regexp)

**Next:** [132-portfolio-website](../132-portfolio-website/skills.md) - 132 Portfolio Website

---

## Need Help?

If you're stuck:
1. **Review previous skills** - Click the "Previous" link above
2. **Re-read the README** - It contains all the theory you need
3. **Draw the logic** on paper - trace through examples manually
4. **Test with simple input** first, then edge cases

**Remember:** Every expert was once a beginner. Take your time, understand each concept, and don't skip ahead!
