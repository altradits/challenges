# 98-lastword

## lastword

### Instructions

Write a function `LastWord` that takes a `string` and returns its last word followed by a `\n`.

- A word is a section of `string` delimited by spaces or by the start/end of the `string`.

### Expected function

```go
func LastWord(s string) string{

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
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
```

And its output :

```console
$ go run . | cat -e
not$
lorem,ipsum$
$
$
```

### Notions

- [01-edu/z01](https://github.com/01-edu/z01)

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 98
```

or run the tests directly:

```bash
go test ./98-lastword/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 98` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
