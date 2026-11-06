# 129-fifthandskip

## fifthandskip

### Instructions

Write a function `FifthAndSkip()` that takes a `string` and returns another `string`. The function separates every five characters of the `string` with a space and removes the sixth one.

- If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a length of 5.
- If the `string` is less than 5 characters return `Invalid Input` followed by a newline `\n`.
- If the `string` is empty return a newline `\n`.

### Expected function

```go
func FifthAndSkip(str string) string {

}
```

### Usage

Example calls (the hidden tests invoke your function directly, in-package):

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
```

And its output:

```console
$ go run . | cat -e
abcde ghijk mnopq stuwx z$
Thisi ashor sente ce$
Invalid Input$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 129
```

or run the tests directly:

```bash
go test ./129-fifthandskip/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 129` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
