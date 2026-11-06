## printif

### Instructions

Write a function that takes a `string` as an argument and returns the letter `G` followed by a newline `\n` if the argument length is more or equal than 3, otherwise returns `Invalid Input` followed by a newline `\n`.

- If it's an empty string return `G` followed by a newline `\n`.

### Expected function

```go
func PrintIf(str string) string {

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
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

```

And its output:

```console
$ go run . | cat -e
G$
G$
G$
Invalid Input$

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 85
```

or run the tests directly:

```bash
go test ./85-printif/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 85` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
