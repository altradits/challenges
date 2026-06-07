# Skills for hiddenp

## What You'll Learn

This exercise is part of the **Expert Checkpoints** series. By completing it, you'll master the logic needed to solve this type of problem and be prepared for similar challenges.

**Challenge:** Write a program named `hiddenp` that takes two `strings` as arguments. The program should check if the first string `s1` is hidden in the second `s2`.
`s1` is considered hidden in `s2` if it is possible to find each character from `s1` in `s2`, in the same order as they appear in `s1`, but not necessarily consecutively.
- If `s1` is hidden in `s2`, the program should display `1` followed by a newline.
- If `s1` is not hidden in `s2`, the program should display `0` followed by a newline.
- If `s1` is an empty string, it is considered hidden in any string.
- If the number of arguments is different from 2, the program should display nothing.
### Usage
```console
$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
$ go run .
$

## Prerequisites

Before starting this exercise, review these fundamental concepts:

### 1. Go Function Basics
```go
func MyFunction(parameter string) int {
    // Your code here
    return 0
}
```

### 2. String Iteration
```go
for i, c := range s {
    // i is index, c is rune
}
```

### 3. Conditional Logic
```go
if condition {
    // do something
} else {
    // do something else
}
```

## Core Concepts to Master

### The Problem Pattern
Understanding the pattern behind this challenge:

1. **Input Analysis** - What type of data are you receiving?
2. **Transformation Logic** - What operation needs to be performed?
3. **Edge Cases** - What happens with empty input, single elements, or boundary values?
4. **Output Format** - What should the function return?

### Key Techniques
- **Iteration**: Processing each element systematically
- **State Tracking**: Remembering previous values or flags
- **Early Returns**: Handling edge cases immediately
- **String Building**: Constructing results character by character

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Analyze problem patterns** - Recognize similar challenges
2. **Implement core algorithms** - Write the logic from scratch
3. **Handle edge cases** - Deal with empty, single, or boundary inputs
4. **Optimize solutions** - Write efficient, readable code

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Text processing and validation
- **Savings Calculator** - Input handling and formatting
- **Investment Tracker** - Data parsing and analysis
- **Net Worth Tracker** - String manipulation for reports

## Step-by-Step Approach

1. **Understand the requirements** - What exactly is being asked?
2. **Identify the pattern** - What algorithm or technique applies?
3. **Plan your solution** - Sketch the logic before coding
4. **Implement incrementally** - Build and test piece by piece
5. **Verify edge cases** - Test with empty, single, and boundary inputs

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Skipping edge cases | Empty/single inputs break | Always test boundaries first |
| Overcomplicating | Simple problems have simple solutions | Start with the simplest approach |
| Off-by-one errors | Index math is tricky | Draw it out on paper |
| Wrong return type | Type mismatch errors | Check function signature |

## Practice Tips

- **Draw it out**: Visualize the problem with diagrams
- **Trace manually**: Follow your code step by step with sample input
- **Test boundaries**: Empty string, single char, max values
- **Refactor**: Can you make it simpler or more efficient?

## The Challenge

Write a program named `hiddenp` that takes two `strings` as arguments. The program should check if the first string `s1` is hidden in the second `s2`.
`s1` is considered hidden in `s2` if it is possible to find each character from `s1` in `s2`, in the same order as they appear in `s1`, but not necessarily consecutively.
- If `s1` is hidden in `s2`, the program should display `1` followed by a newline.
- If `s1` is not hidden in `s2`, the program should display `0` followed by a newline.
- If `s1` is an empty string, it is considered hidden in any string.
- If the number of arguments is different from 2, the program should display nothing.
### Usage
```console
$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
$ go run .
$

### Expected Function
```go

```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| *(See README.md for specific test cases)* | | |

### Usage Example
```go
// See README.md for usage examples
```

## Knowledge Check

Before coding, make sure you can answer:
1. What is the core algorithm needed?
2. What edge cases must be handled?
3. How would you explain this solution to someone else?

## Next Steps

After completing this, you'll be ready for:
- **Next in sequence**: Check the LEARNING_PATH.md for progression
- **Similar patterns**: Look for exercises with the same technique
- **Advanced variations**: Try modifying the challenge constraints

---

## Mastering This Pattern

### Similar Challenges
Look for these related patterns in other exercises:
- String iteration and manipulation
- Conditional logic and early returns
- State tracking across iterations
- Building results incrementally

### Transferable Skills
The techniques you learn here apply to:
- Text processing and validation
- Data parsing and transformation
- Algorithmic problem solving
- Writing clean, maintainable code

### Real-World Applications
This pattern is used in:
- Form validation (checking input formats)
- Data cleaning (removing/transforming characters)
- Text analysis (counting, finding, replacing)
- File processing (parsing line by line)
