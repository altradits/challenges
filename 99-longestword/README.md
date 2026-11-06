# 99-longestword

## longestword

### Instructions

Write a function that takes a string and returns the longest word in the string.

- A word is a sequence of characters delimited by spaces.
- If there are multiple words with the same maximum length, return the first one.
- If the string is empty or contains only spaces, return an empty string.

### Expected function

```go
func LongestWord(s string) string {

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
	fmt.Println(LongestWord("hello world"))
	fmt.Println(LongestWord("the quick brown fox"))
	fmt.Println(LongestWord("a bb ccc dd eee"))
	fmt.Println(LongestWord(""))
	fmt.Println(LongestWord("   "))
	fmt.Println(LongestWord("single"))
}
```

And its output:

```console
$ go run .
hello
quick
ccc

single

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 99
```

or run the tests directly:

```bash
go test ./99-longestword/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 99` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
