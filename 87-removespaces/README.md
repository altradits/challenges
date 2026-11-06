# 87-removespaces

## removespaces

### Instructions

Write a function that takes a string as an argument and returns a new string with all spaces removed.

- If the string is empty, return an empty string.
- Only remove space characters (' '), not other whitespace like tabs or newlines.

### Expected function

```go
func RemoveSpaces(s string) string {

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
	fmt.Println(RemoveSpaces("Hello World"))
	fmt.Println(RemoveSpaces("Go is fun!"))
	fmt.Println(RemoveSpaces(""))
	fmt.Println(RemoveSpaces("NoSpacesHere"))
}
```

And its output:

```console
$ go run .
HelloWorld
Goisfun!

NoSpacesHere
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
go run ./cmd/drill check 87
```

or run the tests directly:

```bash
go test ./87-removespaces/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 87` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
