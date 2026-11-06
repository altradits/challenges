# 106-cleanstr

## cleanstr

### Instructions

Write a **program** that takes a `string`, and displays this `string` with exactly:

- one space between words.
- without spaces nor tabs at the beginning nor at the end.
- with the result followed by a newline ("`\n`").

A "word" is defined as a part of a `string` delimited either by spaces/tabs, or
by the start/end of the `string`.

If the number of arguments is not 1, or if there are no words to display, the
program displays a newline("`\n`").

### Usage :

```console
$ go run . "you see it's easy to display the same thing" | cat -e
you see it's easy to display the same thing$
$ go run . " only    it's  harder   "
only it's harder$
$ go run . " how funny" "Did you   hear Mathilde ?"

$ go run . ""

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
go run ./cmd/drill check 106
```

or run the tests directly:

```bash
go test ./106-cleanstr/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 106` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
