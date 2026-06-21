# Prerequisites for strlenindex

## Basic Skills Needed

Before starting this challenge, you should know:

1. **Basic Go program structure**
   ```go
   package main

   import "fmt"

   func main() {
       // your code
   }
   ```

2. **How to declare a variable**
   ```go
   text := "stano"
   ```

3. **How to print output**
   ```go
   fmt.Println("hello")
   fmt.Printf("value: %d\n", 42)
   ```

4. **That strings are sequences (from hello/)**

   You have seen that a string is a sequence of characters at numbered positions. This challenge makes that concrete.

## You Do Not Need

- Any prior knowledge of how indexes work
- Experience with arrays or slices
- Knowledge of byte values or Unicode

This challenge teaches those things from scratch.

## Skills You Will Gain

After completing this challenge, you will be able to:

1. **Use string indexing** — access any character by position with `str[i]`
2. **Use len()** — get the total character count of a string
3. **Explain the len/index gap** — len counts from 1, indexes start at 0
4. **Use %c in fmt.Printf** — display a byte as a character instead of a number
5. **Calculate the last index** — `len(str) - 1` always gives the last valid position

## How This Connects to Bitcoin Development

String indexing is used constantly in Bitcoin code:

- Parsing transaction hex strings character by character
- Validating Bitcoin address formats (checking prefix characters)
- Navigating script byte sequences in Bitcoin Script

Understanding how len() and indexes relate is not optional knowledge — it is the foundation.

## What Comes Next

After completing this challenge, move to:
- [printlastchar/](../printlastchar/prerequisites.md) — Use len-1 to print the last character of any string
