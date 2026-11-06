# 116-zipstring

## zipstring

### Instructions

Write a function that takes a `string` and returns a new `string` that replaces every character with the number of duplicates and the character itself, deleting the extra duplications.

- The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.

### Expected function

```go
func ZipString(s string) string {

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
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
```

And its output:

```console
$ go run .
1Y1o3u1n1g1F1e4l1a1s
1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
1H1e2l2o1 1T1h1e2r1e1!
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 116
```

or run the tests directly:

```bash
go test ./116-zipstring/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 116` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
