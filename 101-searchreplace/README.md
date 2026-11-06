# 101-searchreplace

## searchreplace

### Instructions

Write a program that takes 3 arguments, the first argument is a `string` in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

- If the number of arguments is different from 3, the program displays nothing.

- If the second argument is not contained in the first one (the string) then the program rewrites the `string` followed by a newline (`'\n'`).

### Usage

```console
$ go run . "hella there" "a" "o"
hello there
$ go run . "hallo thara" "a" "e"
hello there
$ go run . "abcd" "z" "l"
abcd
$ go run . "something" "a" "o" "b" "c"
$

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 101
```

or run the tests directly:

```bash
go test ./101-searchreplace/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 101` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
