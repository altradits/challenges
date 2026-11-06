# 128-wdmatch

## wdmatch

### Instructions

Write a program that takes two `string` and checks whether it is possible to write the first `string` with characters from the second `string`. This rewrite must respect the order in which these characters appear in the second `string`.

If it is possible, the program displays the `string` followed by a newline (`'\n'`), otherwise it simply displays nothing.

If the number of arguments is different from 2, the program displays nothing.

### Usage

```console
$ go run . 123 123
123
$ go run . faya fgvvfdxcacpolhyghbreda
faya
$ go run . faya fgvvfdxcacpolhyghbred
$ go run . error rrerrrfiiljdfxjyuifrrvcoojh
$ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux
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
go run ./cmd/drill check 128
```

or run the tests directly:

```bash
go test ./128-wdmatch/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 128` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
