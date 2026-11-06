## Can Jump

Given an array of non-negative integers representing the number of steps you can take forward from each position, implement the function `CanJump()` which takes a slice of unsigned integers `[]uint` as input and returns a `boolean` value. This function should determine if it's possible to reach and stay at the last index of the array starting from the first index, based on the steps you need to advance. Be aware that:

- Each value represents the exact number of steps you must take forward from that position.
- The function should return `true` if it's possible to reach and stay at the last index without stepping out of the array, and `false` otherwise.
- If the input has only one element, that is the last position in the array so the function will return `true` but if the array is empty it returns `false`.

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(piscine.CanJump(input3))
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
