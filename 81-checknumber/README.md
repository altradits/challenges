## checknumber

### Instructions

Write a function that takes a `string` as an argument and returns `true` if the string contains any number, otherwise return `false`.

### Expected function

```go
func CheckNumber(arg string) bool {
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
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

```

And its output:

```console
$ go run .
false
true
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
go run ./cmd/drill check 81
```

or run the tests directly:

```bash
go test ./81-checknumber/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 81` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
