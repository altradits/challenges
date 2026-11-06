## rostring

### Instructions

Write a program that takes a `string` and displays this `string` after rotating it one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A word is a sequence of **alphanumerical** characters.

Words will be separated by only one space in the output.

If the number of arguments is different from 1, the program displays a newline.

### Usage

```console
$ go run . "abc   " | cat -e
abc$
$ go run . "Let there     be light"
there be light Let
$ go run . "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
$ go run . | cat -e
$
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
go run ./cmd/drill check 135
```

or run the tests directly:

```bash
go test ./135-rostring/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 135` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
