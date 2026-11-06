# 124-inter

## inter

### Instructions

Write a program that takes two `string` and displays, without doubles, the characters that appear in both `string`, in the order they appear in the first one.

- The display will be followed by a newline (`'\n'`).

- If the number of arguments is different from 2, the program displays nothing.

### Usage

```console
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
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
go run ./cmd/drill check 124
```

or run the tests directly:

```bash
go test ./124-inter/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 124` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
