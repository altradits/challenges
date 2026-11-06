# 109-fromto

## fromto

### Instructions

Write a function that takes two `integers` and returns a `string` showing the range of numbers from the first to the second.

- The numbers must be separated by a comma and a space.
- If any of the arguments is bigger than `99` or less than `0`, the function returns `Invalid` followed by a newline `\n`.
- Prepend a `0` to any number that is less than `10`.
- Add a new line `\n` at the end of the `string`.

### Expected function

```go
func FromTo(from int, to int) string {

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
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
```

And its output:

```console
$ go run . | cat -e
01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
10$
Invalid$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 109
```

or run the tests directly:

```bash
go test ./109-fromto/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 109` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
