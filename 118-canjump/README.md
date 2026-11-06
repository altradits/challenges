# 118-canjump — Can Jump

## Task

Given a slice of unsigned integers where each value is the exact number
of steps you must jump forward from that position, write in
`package piscine`:

```go
func CanJump(nums []uint) bool
```

Start at index 0 and follow the jumps. Return `true` if you land on or
beyond the last index, `false` if you get stuck on a `0` before the end.
A single element is already the last index (`true`); empty input is
`false`.

### Usage

Example calls (the hidden tests invoke your function directly, in-package):

```go
package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
```

And its output :

```console
$ go run .
true
false
true
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
go run ./cmd/drill check 118
```

or run the tests directly:

```bash
go test ./118-canjump/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 118` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
