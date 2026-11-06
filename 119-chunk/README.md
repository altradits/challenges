## chunk

### Instructions

Write a function called `Chunk` that receives as parameters a slice, `slice []int`, and a number `size int`. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of `size`.

- If the `size` is `0` it should print a newline (`'\n'`).

### Expected function

```go
func Chunk(slice []int, size int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
```

And its output :

```console
$ go run .
[]

[[0 1 2] [3 4 5] [6 7]]
[[0 1 2 3 4] [5 6 7]]
[[0 1 2 3] [4 5 6 7]]
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
go run ./cmd/drill check 119
```

or run the tests directly:

```bash
go test ./119-chunk/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 119` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
