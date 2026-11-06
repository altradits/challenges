# 86-printifnot

## printifnot

### Instructions

Write a function that takes a string and returns:

- `"G\n"` if the string's length is less than 3 (including empty string).

- `"Invalid Input\n"` otherwise.

### Expected function

```go
func PrintIfNot(str string) string {

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
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
```

And its output:

```console
$ go run . | cat -e
Invalid Input$
Invalid Input$
G$
G$

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 86
```

or run the tests directly:

```bash
go test ./86-printifnot/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 86` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
