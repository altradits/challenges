# 12-recursion

## recursion

### Instructions

Write a function that takes a `uint` and returns its factorial using recursion.

The factorial of a non-negative integer `n` is defined as:

- `0! = 1` (base case)
- `1! = 1` (base case)
- `n! = n × (n-1)!` for `n > 1` (recursive case)

Your function **must** use recursion — that is, it must call itself.

### Expected function

```go
func Factorial(n uint) uint {

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
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(1))
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(10))
}
```

And its output:

```console
$ go run .
1
1
120
3628800
$
```

### Test Cases

| Input | Expected output |
|-------|----------------|
| `Factorial(0)` | `1` |
| `Factorial(1)` | `1` |
| `Factorial(5)` | `120` |
| `Factorial(10)` | `3628800` |

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 12
```

or run the tests directly:

```bash
go test ./12-recursion/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 12` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
