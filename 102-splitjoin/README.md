## splitjoin

### Instructions

Write two functions:

1. `Split(s string, sep string) []string` - Splits a string by a separator and returns a slice of substrings.
2. `Join(arr []string, sep string) string` - Joins a slice of strings with a separator and returns a single string.

- If the separator is empty, each character becomes its own element in the slice.
- If the input string is empty, return a slice with one empty string element.
- Do not use the built-in `strings.Split` or `strings.Join` functions.

### Expected functions

```go
func Split(s string, sep string) []string {

}

func Join(arr []string, sep string) string {

}
```

### Usage

Here is a possible program to test your functions:

```go
package main

import (
	"fmt"
)

func main() {
	// Test Split
	fmt.Printf("%q\n", Split("a,b,c", ","))
	fmt.Printf("%q\n", Split("Hello World", " "))
	fmt.Printf("%q\n", Split("abc", ""))
	fmt.Printf("%q\n", Split("", ","))

	// Test Join
	fmt.Println(Join([]string{"a", "b", "c"}, ","))
	fmt.Println(Join([]string{"Hello", "World"}, " "))
	fmt.Println(Join([]string{"a", "b", "c"}, ""))
	fmt.Println(Join([]string{""}, ","))
}
```

And its output:

```console
$ go run .
["a" "b" "c"]
["Hello" "World"]
["a" "b" "c"]
[""]
a,b,c
Hello World
abc

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
go run ./cmd/drill check 102
```

or run the tests directly:

```bash
go test ./102-splitjoin/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 102` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
