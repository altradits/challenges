## expandstr

### Instructions

Write a program that takes a `string` and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

The `string` will be followed by a newline (`'\n'`).

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.

### Usage

```console
$ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
$ go run . "   only  it's harder   " | cat -e
only   it's   harder$
$ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
$ go run .
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
go run ./cmd/drill check 107
```

or run the tests directly:

```bash
go test ./107-expandstr/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 107` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
