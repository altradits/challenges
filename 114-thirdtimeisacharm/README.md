# 114-thirdtimeisacharm

## thirdtimeisacharm

### Instructions

Write a function `ThirdTimeIsACharm()` that takes a `string` as an argument and returns another `string` with every third character.

- Return the output followed by a newline `\n`.
- If the `string` is empty, return a newline `\n`.
- If there is no third character, return a newline `\n`.

### Expected function

```go
func ThirdTimeIsACharm(str string) string {

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
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
```

And its output:

```console
$ go run . | cat -e
369$
$
b e$
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
go run ./cmd/drill check 114
```

or run the tests directly:

```bash
go test ./114-thirdtimeisacharm/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 114` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
