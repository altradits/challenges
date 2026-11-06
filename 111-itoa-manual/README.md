# 111-itoa-manual

## itoa

### Instructions

- Write a function that simulates the behavior of the `Itoa` function in Go. `Itoa` transforms a number represented as an`int` in a number represented as a `string`.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

## Expected function

```go
func Itoa(n int) string {

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
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}
```

And its output :

```console
$ go run .
12345
0
-1234
987654321
$
```

### Notions

- [strconv/Itoa](https://pkg.go.dev/strconv#Itoa)

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 111
```

or run the tests directly:

```bash
go test ./111-itoa-manual/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 111` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
