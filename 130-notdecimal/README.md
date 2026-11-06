## notdecimal

### Instructions

Write a function called `NotDecimal()` that takes as an argument a `string` in form of a float number with the decimal point and returns a string converted to `int` without the decimal point (you will have to multiply it by 10^n to remove the `.`).

- If the number doesn't have a decimal point or there is only a zero after the `.` return the number followed by a newline `\n`.
- If the argument is empty return a newline `\n`.
- If the argument is not a number return it followed by a newline `\n`.

### Expected function

```go
func NotDecimal(dec string) string {

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
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
```

And its output:

```console
$ go run .  | cat -e
1$
1742$
1255$
120525856$
-0.0f00d00$
$
-19525856$
1952$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 130
```

or run the tests directly:

```bash
go test ./130-notdecimal/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 130` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
