## hiddenp

### Instructions

Write a program named `hiddenp` that takes two `strings` as arguments. The program should check if the first string `s1` is hidden in the second `s2`.
`s1` is considered hidden in `s2` if it is possible to find each character from `s1` in `s2`, in the same order as they appear in `s1`, but not necessarily consecutively.

- If `s1` is hidden in `s2`, the program should display `1` followed by a newline.
- If `s1` is not hidden in `s2`, the program should display `0` followed by a newline.
- If `s1` is an empty string, it is considered hidden in any string.
- If the number of arguments is different from 2, the program should display nothing.

### Usage

```console
$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
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
go run ./cmd/drill check 123
```

or run the tests directly:

```bash
go test ./123-hiddenp/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 123` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
