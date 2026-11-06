# 115-weareunique

## weareunique

### Instructions

Write a function that takes two `strings`'s and returns the number of characters that are not included in both, without repeating characters.

- If there is no unique characters return `0`.
- If both strings are empty return `-1`.

### Expected function

```go
func WeAreUnique(str1 , str2 string) int {

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
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
```

And its output:

```console
$ go run .
2
-1
6
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 115
```

or run the tests directly:

```bash
go test ./115-weareunique/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 115` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
