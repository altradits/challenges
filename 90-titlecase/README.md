# 90-titlecase

## Title Case

### Instructions

Write a function that takes a string as an argument and returns the string converted to title case.

- Title case means the first letter of each word is uppercase, and all other letters are lowercase.
- A word is defined as a sequence of characters separated by spaces.
- If the string is empty, return an empty string.

### Expected function

```go
func TitleCase(s string) string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(TitleCase("hello world"))
	fmt.Println(TitleCase("Go is fun!"))
	fmt.Println(TitleCase(""))
	fmt.Println(TitleCase("multiple   spaces   here"))
	fmt.Println(TitleCase("ALL CAPS"))
}
```

And its output:

```console
$ go run .
Hello World
Go Is Fun!

Multiple   Spaces   Here
All Caps
$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 90
```

or run the tests directly:

```bash
go test ./90-titlecase/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 90` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
