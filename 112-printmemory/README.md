# 112-printmemory

## printmemory

### Instructions

Write a function that takes `(arr [10]byte)`, and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.

### Expected function

```go
func PrintMemory(arr [10]byte) {

}
```

### Usage

Example calls (the hidden tests invoke your function directly, in-package):

```go
package main


func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
```

And its output :

```console
$ go run . | cat -e
68 65 6c 6c$
6f 10 15 2a$
00 00$
hello..*..$
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
go run ./cmd/drill check 112
```

or run the tests directly:

```bash
go test ./112-printmemory/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 112` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
