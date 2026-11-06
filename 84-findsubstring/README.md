# 84-findsubstring

## findsubstring

### Instructions

Write a function that takes two strings as arguments: a `text` and a `substring`. The function should return the index of the first occurrence of `substring` in `text`.

- If the `substring` is not found, return -1.
- If the `substring` is empty, return 0.
- The search should be case-sensitive.

### Expected function

```go
func FindSubstring(text, substring string) int {

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
	fmt.Println(FindSubstring("Hello World", "World"))
	fmt.Println(FindSubstring("Hello World", "hello"))
	fmt.Println(FindSubstring("Hello World", ""))
	fmt.Println(FindSubstring("Hello World", "xyz"))
	fmt.Println(FindSubstring("banana", "ana"))
}
```

And its output:

```console
$ go run .
6
-1
0
-1
1
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
go run ./cmd/drill check 84
```

or run the tests directly:

```bash
go test ./84-findsubstring/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 84` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
