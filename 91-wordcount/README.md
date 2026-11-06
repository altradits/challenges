# 91-wordcount

## wordcount

### Instructions

Write a function that takes a string as an argument and returns the number of words in the string.

- A word is defined as a sequence of non-space characters separated by one or more spaces.
- Leading and trailing spaces should be ignored.
- If the string is empty or contains only spaces, return 0.

### Expected function

```go
func WordCount(s string) int {

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
	fmt.Println(WordCount("Hello World"))
	fmt.Println(WordCount("Go is fun!"))
	fmt.Println(WordCount(""))
	fmt.Println(WordCount("   "))
	fmt.Println(WordCount("One  two   three"))
}
```

And its output:

```console
$ go run .
2
3
0
0
3
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
go run ./cmd/drill check 91
```

or run the tests directly:

```bash
go test ./91-wordcount/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 91` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
