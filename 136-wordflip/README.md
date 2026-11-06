# 136-wordflip

## wordflip

### Instructions

Write a function `WordFlip()` that takes a `string` as input and returns it in reverse order.

- The output should be followed by a newline `\n`.
- If the string is empty, return `Invalid Output`.
- Ignore multiple spaces between words and trim any leading or trailing spaces in the string.

### Expected function

```go
func WordFlip(str string) string {

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
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
```

And its output:

```console
$ go run . | cat -e
last second First$
Invalid Output$
$
you! of all hello$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 136
```

or run the tests directly:

```bash
go test ./136-wordflip/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 136` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
