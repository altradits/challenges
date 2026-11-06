## countrepeats

### Instructions

Write a function that takes a string and returns the number of consecutive repeated characters in the string.

- A "repeat" is when the same character appears consecutively two or more times.
- Count each group of consecutive repeats as one.
- If the string is empty or has no consecutive repeats, return 0.

### Expected function

```go
func CountRepeats(s string) int {

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
	fmt.Println(CountRepeats("hello"))
	fmt.Println(CountRepeats("helloo"))
	fmt.Println(CountRepeats("heelloo"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats(""))
	fmt.Println(CountRepeats("aaa"))
}
```

And its output:

```console
$ go run .
1
2
2
0
0
1

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 93
```

or run the tests directly:

```bash
go test ./93-countrepeats/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 93` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
