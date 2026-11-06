## forloops

### Instructions

Write a function called `CountDown` that takes an integer `n` as an argument and prints all integers from `n` down to `0`, each on its own line.

### Expected function

```go
func CountDown(n int) {
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	piscine.CountDown(5)
	fmt.Println("---")
	piscine.CountDown(0)
}
```

And its output:

```console
$ go run .
5
4
3
2
1
0
---
0
$
```

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 7
```

or run the tests directly:

```bash
go test ./07-forloops/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 7` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
