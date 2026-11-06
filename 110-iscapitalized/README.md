## iscapitalized

### Instructions

Write a function `IsCapitalized()` that takes a `string` as an argument and returns `true` if each word in the `string` begins with either an uppercase letter or a non-alphabetic character.

- If any of the words begin with a lowercase letter return `false`.
- If the `string` is empty return `false`.

### Expected function

```go
func IsCapitalized(s string) bool {

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
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
```

And its output:

```console
$ go run .
false
true
true
true
true
false
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 110
```

or run the tests directly:

```bash
go test ./110-iscapitalized/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 110` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
