## printrevcomb

### Instructions

Write a program that prints in descending order on a single line all unique combinations of three different digits so that the first digit is greater than the second and the second is greater than the third.

These combinations are separated by a comma and a space.

### Usage

Here is an **incomplete** output :

```console
$ go run . | cat -e
987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
$
```

`999` or `000` are not valid combinations because the digits are not different.

`789` should not be shown because the first digit is not greater than the second.

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 113
```

or run the tests directly:

```bash
go test ./113-printrevcomb/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 113` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
