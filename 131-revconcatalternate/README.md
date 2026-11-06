# 131-revconcatalternate

## revconcatalternate

### Instructions

Write a function `RevConcatAlternate()` that receives two slices of `int` as arguments and returns a new slice with alternated values of each slice in reverse order.

- The input slices can have different lengths.
- The new slice should start with the elements from the largest slice first and when they became equal size slices, it should add an element of the first given slice.
- If the slices are of equal length, the new slice should start with an element of the first slice.

> Note: you can check the examples bellow for more details.

### Expected function

```go
func RevConcatAlternate(slice1,slice2 []int) []int {

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
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
```

And its output:

```console
$ go run .
[3 6 2 5 1 4]
[9 8 7 3 6 2 5 1 4]
[8 9 3 2 5 1 4]
[3 2 1]
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 131
```

or run the tests directly:

```bash
go test ./131-revconcatalternate/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 131` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
