# Skills for hello

## What You'll Learn

**Challenge:** Hello World — Preview of what strings really are

## New Concepts Explained

### 1. The basic structure of a Go program

Every Go program you will ever write follows this pattern:

```go
package main

import "fmt"

func main() {
    // your code here
}
```

- `package main` — tells Go this file is an executable program, not a library
- `import "fmt"` — loads the `fmt` package so you can print output
- `func main()` — the entry point; execution always starts here

### 2. Printing output

`fmt.Println` prints a value and moves to the next line:

```go
fmt.Println("Hello World.")
```

The string between the quotes is printed exactly as written.

### 3. Strings are sequences

This is the key idea to carry forward. A string like `"Bitcoin"` is not just a word — it is a sequence of individual characters laid out one after another:

```
B  i  t  c  o  i  n
0  1  2  3  4  5  6
```

Each character sits at a numbered position called an index. Indexes start at 0, not 1. This matters enormously in the challenges that follow.

### 4. Preview: Sliding windows

The message printed in this program — "Ready to learn about sliding windows" — refers to one of the most important string algorithms. A sliding window moves through a string one position at a time, looking at a fixed-size slice of characters.

To understand sliding windows, you first need to understand that strings have positions and that you can navigate those positions with indexes. That foundation starts here.

## The Big Idea

Before you write any algorithm, you need to see strings the way computers see them: not as words, but as arrays of characters sitting at numbered positions.

**Next:** [strlenindex](../strlenindex/skills.md) — Understanding len() and index positions
