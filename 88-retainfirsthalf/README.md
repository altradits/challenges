## retainfirsthalf

### Instructions

Write a function called `RetainFirstHalf()` that takes a `string` as an argument and returns the first half of this `string`.

- If the length of the `string` is odd, round it down.
- If the `string` is empty, return an empty string.
- If the `string` length is equal to one, return the string.

### Expected function

```go
func RetainFirstHalf(str string) string {

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
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}
```

And its output:

```console
$ go run . | cat -e
This is the 1st half$
A$
$
Hello$

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 88
```

or run the tests directly:

```bash
go test ./88-retainfirsthalf/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 88` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
