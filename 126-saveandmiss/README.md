# 126-saveandmiss

## saveandmiss

### Instructions

Write a function called `SaveAndMiss()` that takes a `string` and an `int` as an argument. The function should move through the `string` in sets determined by the `int`, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the `string` is reached. Return a `string` containing the saved characters.

> If the `int` is `0` or a negative number return the original `string`.

### Expected function

```go
func SaveAndMiss(arg string, num int) string {

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
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
```

And its output:

```console
$ go run . | cat -e
123789$
abcghimnostuz$
$
hello you all ! $
what is your name?$
go Exercise Save and Miss$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 126
```

or run the tests directly:

```bash
go test ./126-saveandmiss/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 126` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
