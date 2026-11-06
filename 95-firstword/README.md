## firstword

### Instructions

Write a function that takes a string and return a string containing its first word, followed by a newline (`'\n'`).

- A word is a sequence of characters delimited by spaces or by the start/end of the argument.

### Expected Function

```go
func FirstWord(s string) string {
    // ...
}
```

### Usage

Here is a possible way to test your function:

```go
package main

import (
    "fmt"

    "piscine"
)

func main() {
    fmt.Print(piscine.FirstWord("hello there"))
    fmt.Print(piscine.FirstWord(""))
    fmt.Print(piscine.FirstWord("hello   .........  bye"))
}
```

And its output:

```console
$ go run .
hello

hello
$

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 95
```

or run the tests directly:

```bash
go test ./95-firstword/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 95` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
