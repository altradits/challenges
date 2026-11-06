# 120-concatalternate

## concatalternate

### Instructions

Write a function `ConcatAlternate()` that receives two slices of an `int` as arguments and returns a new slice with the result of the alternated values of each slice.

- The input slices can be of different lengths.
- The new slice should start with an element of the largest slice.
- If the slices are of equal length, the new slice should return the elements of the first slice first and then the elements of the second slice.

### Expected function

```go
func ConcatAlternate(slice1,slice2 []int) []int {

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
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
```

And its output:

```console
$ go run .
[1 4 2 5 3 6]
[1 2 3 4 5 6 7 8 9 10 11]
[4 1 5 2 6 3 7 8 9]
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
go run ./cmd/drill check 120
```

or run the tests directly:

```bash
go test ./120-concatalternate/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 120` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
