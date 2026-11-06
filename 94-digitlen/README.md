# 94-digitlen

## digitlen

### Instructions

Write a function `DigitLen()` that takes two integers as arguments and returns the times the first `int` can be divided by the second until it reaches zero.

- The second `int` must be between **_2_** and **_36_**. If not, the function returns `-1`.
- If the first `int` is negative, reverse the sign and count the digits.

### Expected function

```go
func DigitLen(n, base int) int {

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
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
```

And its output:

```console
$ go run . | cat -e
3$
7$
2$
-1$

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 94
```

or run the tests directly:

```bash
go test ./94-digitlen/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 94` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
