# 97-hashcode

## hashcode

### Instructions

Write a function called `HashCode()` that takes a `string` as an argument and returns a new **hashed** `string`.

- The hash equation is computed as follows:

`(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.`

- If the resulting character is unprintable add `33` to it.

### Expected function

```go
func HashCode(dec string) string {
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
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
```

And its output:

```console
$ go run .
B
CD
EDF
Spwwz+bz}wo

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 97
```

or run the tests directly:

```bash
go test ./97-hashcode/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 97` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
