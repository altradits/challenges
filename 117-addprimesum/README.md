## addprimesum

### Instructions

Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline (`'\n'`).

- If the number of arguments is different from 1, or if the argument is not a positive number, the program displays `0` followed by a newline.

### Usage

```console
$ go run . 5
10
$ go run . 7
17
$ go run . -2
0
$ go run . 0
0
$ go run .
0
$ go run . 5 7
0
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
go run ./cmd/drill check 117
```

or run the tests directly:

```bash
go test ./117-addprimesum/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 117` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
