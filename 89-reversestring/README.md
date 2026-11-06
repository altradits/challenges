## reversestring

### Instructions

Write a function that takes a string as an argument and returns the string reversed.

- If the string is empty, return an empty string.
- Do not use any built-in reverse functions. Build the reversed string manually.

### Expected function

```go
func ReverseString(s string) string {

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
	fmt.Println(ReverseString("Hello"))
	fmt.Println(ReverseString("Go is fun!"))
	fmt.Println(ReverseString(""))
	fmt.Println(ReverseString("a"))
}
```

And its output:

```console
$ go run .
olleH
!nuf si oG

a
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
go run ./cmd/drill check 89
```

or run the tests directly:

```bash
go test ./89-reversestring/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 89` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
