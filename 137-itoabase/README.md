# 137-itoabase

## itoabase

### Instructions

Write a function that:

- converts an `int` value to a `string` using the specified base in the argument
- and that returns this `string`

The base is expressed as an `int`, from 2 to 16. The characters comprising the base are the digits from 0 to 9, followed by uppercase letters from A to F.

For example, the base `4` would be the equivalent of "0123" and the base `16` would be the equivalent of "0123456789ABCDEF".

If the value is negative, the resulting `string` has to be preceded by a minus sign `-`.

Only valid inputs will be tested.

### Expected function

```go
func ItoaBase(value, base int) string {

}
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 137
```

or run the tests directly:

```bash
go test ./137-itoabase/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 137` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
