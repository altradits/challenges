## arrays

### Instructions

Write a function `SumArray` that takes a fixed-size array of 5 integers and returns the sum of all 5 values.

### Expected function

```go
func SumArray(arr [5]int) int {

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
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(piscine.SumArray(arr)) // 15

	arr2 := [5]int{0, 0, 0, 0, 0}
	fmt.Println(piscine.SumArray(arr2)) // 0

	arr3 := [5]int{-1, -2, 3, 4, 5}
	fmt.Println(piscine.SumArray(arr3)) // 9
}
```

### Test cases

| Input                        | Output |
|-----------------------------|--------|
| `[5]int{1, 2, 3, 4, 5}`     | `15`   |
| `[5]int{0, 0, 0, 0, 0}`     | `0`    |
| `[5]int{-1, -2, 3, 4, 5}`   | `9`    |
| `[5]int{10, 20, 30, 40, 50}`| `150`  |
| `[5]int{-5, -5, -5, -5, -5}`| `-25`  |

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 11
```

or run the tests directly:

```bash
go test ./11-arrays/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 11` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
