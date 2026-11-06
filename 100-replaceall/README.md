# 100-replaceall

## replaceall

### Instructions

Write a function that takes three strings as arguments: `text`, `old`, and `new`. The function should return a new string where all occurrences of `old` in `text` are replaced with `new`.

- If `old` is empty, return `text` unchanged.
- The replacement should be case-sensitive.
- Do not use the built-in `strings.ReplaceAll` function. Build the result manually.

### Expected function

```go
func ReplaceAll(text, old, new string) string {

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
	fmt.Println(ReplaceAll("Hello World", "l", "L"))
	fmt.Println(ReplaceAll("banana", "na", "NA"))
	fmt.Println(ReplaceAll("Hello World", "", "x"))
	fmt.Println(ReplaceAll("abcabc", "abc", "xyz"))
	fmt.Println(ReplaceAll("no match here", "z", "Z"))
}
```

And its output:

```console
$ go run .
HeLLo WorLd
baNANA
Hello World
xyzxyz
no match here
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
go run ./cmd/drill check 100
```

or run the tests directly:

```bash
go test ./100-replaceall/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 100` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
