## concatslice

### Instructions

Write a function `ConcatSlice()` that takes two slices of integers as arguments and returns the concatenation of the two slices.

### Expected function

```go
func ConcatSlice(slice1, slice2 []int) []int {

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
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
}
```

And its output:

```console
$ go run .
[1 2 3 4 5 6]
[4 5 6 7 8 9]
[1 2 3]
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 121
```

or run the tests directly:

```bash
go test ./121-concatslice/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 121` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
