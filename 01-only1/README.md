# 01-only1 — Print 1

## Task

Write a program in `package main` that prints the single character `1` to
standard output, followed by a newline. It must ignore any command-line
arguments and print nothing else.

## Input / Output

- Input: none (any arguments are ignored).
- Output: exactly `1\n` on stdout.

## Examples

```console
$ go run .
1
$ go run . "a" "b"
1
$ go run . "a" "b" "c"
1
```

## Constraints

- Output is exactly two bytes: `1` and `\n`. No extra spaces, no extra lines.
- Must behave identically with zero, two, or three arguments.

## Time limit

**5 minutes** (round 1). Beat it on every repeat: the round-4 target is under
a minute — this skeleton must be pure muscle memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference
solution, turn off autocomplete if you can. Write every line from memory —
that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 01
```

or run the tests directly:

```bash
go test ./01-only1/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 01` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
