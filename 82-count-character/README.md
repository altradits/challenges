# 82-count-character

## count-character

### Instructions

write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.

- if the character is not in the string return 0
- if the string is empty return 0

### Expected Function

```go
func CountChar(str string, c rune) int {
    ...
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
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}

```

And its output :

```console
$ go run .
3
0
3
1
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 82
```

or run the tests directly:

```bash
go test ./82-count-character/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 82` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
