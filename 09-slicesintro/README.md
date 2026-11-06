## slicesintro

### Instructions

Write a function `Unique` that takes a slice of integers and returns a new slice containing only the unique values, in the order they first appeared.

### Expected function

```go
func Unique(nums []int) []int {

}
```

### Usage

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Unique([]int{1, 2, 2, 3, 1, 4})) // [1 2 3 4]
	fmt.Println(piscine.Unique([]int{}))                  // []
	fmt.Println(piscine.Unique([]int{5}))                 // [5]
}
```

### Test cases

| Input                      | Output        |
|---------------------------|---------------|
| `[]int{1, 2, 2, 3, 1, 4}` | `[]int{1, 2, 3, 4}` |
| `[]int{}`                  | `[]int{}`     |
| `[]int{5}`                 | `[]int{5}`    |
| `[]int{7, 7, 7}`           | `[]int{7}`    |
| `[]int{3, 1, 2, 1, 3}`    | `[]int{3, 1, 2}` |

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 9
```

or run the tests directly:

```bash
go test ./09-slicesintro/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 9` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
