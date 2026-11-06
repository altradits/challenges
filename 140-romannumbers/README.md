# 140-romannumbers

## romannumbers

### Instructions

Write a program called `rn`. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.

The program should have a limit of `4000`. In case of an invalid number, for example `"hello"` or `0` the program should print `ERROR: cannot convert to roman digit`.

### Usage

```console
$ go run . hello
ERROR: cannot convert to roman digit
$ go run . 123
C+X+X+I+I+I
CXXIII
$ go run . 999
(M-C)+(C-X)+(X-I)
CMXCIX
$ go run . 3999
M+M+M+(M-C)+(C-X)+(X-I)
MMMCMXCIX
$ go run . 4000
ERROR: cannot convert to roman digit
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
go run ./cmd/drill check 140
```

or run the tests directly:

```bash
go test ./140-romannumbers/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 140` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
