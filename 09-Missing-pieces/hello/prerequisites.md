# Prerequisites for hello

## Basic Skills Needed

Before starting this challenge, you should be comfortable with:

1. **Installing Go and running a program**

   You need Go installed and a terminal where you can run:
   ```bash
   go run main.go
   ```

2. **What a package declaration is**

   Every Go file starts with `package`. For executable programs:
   ```go
   package main
   ```

3. **What an import statement is**

   To use code from another package, you import it:
   ```go
   import "fmt"
   ```

4. **What a function is (basic idea)**

   A function is a named block of code. `main()` is the function Go looks for when it starts your program:
   ```go
   func main() {
       // code runs here
   }
   ```

## You Do Not Need

- Any knowledge of string indexing
- Any knowledge of arrays or slices
- Any experience with loops or conditions

This is a starting point. The goal is to run the program and read the message.

## Skills You Will Gain

After completing this challenge, you will understand:

1. The skeleton structure of a Go program
2. That strings are sequences of characters at numbered positions
3. What "sliding windows" means as a concept (preview only)

## What Comes Next

After `hello/`, move to:
- [strlenindex/](../strlenindex/prerequisites.md) — Learn how len() and indexes relate
- [printlastchar/](../printlastchar/prerequisites.md) — Apply that knowledge to get the last character
