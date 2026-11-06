# 145-time — The time Package

## Task

Implement four helpers in `package piscine` using the standard `time` package:

```go
// AgeInYears returns how many full years have passed since birthDate.
func AgeInYears(birthDate time.Time) int

// FormatLog formats a time as "2006-01-02 15:04:05" (the log timestamp format).
func FormatLog(t time.Time) string

// ParseDate parses a date string in "2006-01-02" format.
func ParseDate(s string) (time.Time, error)

// BusinessDays returns the number of weekdays between start and end (exclusive of end).
func BusinessDays(start, end time.Time) int
```

## Input / Output

- `AgeInYears`: a birth date → full years elapsed (birthday not yet reached
  this year does not count).
- `FormatLog`: a `time.Time` → `"2006-01-02 15:04:05"` in UTC.
- `ParseDate`: `"2006-01-02"` → that date at midnight UTC, or a non-nil error.
- `BusinessDays`: two dates → count of Monday–Friday days in
  `[start, end)`; `0` when `end` is not after `start`.

## Examples

```go
AgeInYears(time.Date(2000, 6, 15, 0, 0, 0, 0, time.UTC))  // full years since then

FormatLog(time.Date(2026, 11, 6, 9, 30, 5, 0, time.UTC))
// "2026-11-06 09:30:05"

ParseDate("2026-11-06")  // 2026-11-06 00:00:00 +0000 UTC, nil
ParseDate("06/11/2026")  // error (wrong layout)

// Monday 2026-11-02 → Monday 2026-11-09 counts Mon..Fri = 5 weekdays
BusinessDays(time.Date(2026, 11, 2, 0, 0, 0, 0, time.UTC),
             time.Date(2026, 11, 9, 0, 0, 0, 0, time.UTC))  // 5
```

## Constraints

- Use `time` layout strings (`"2006-01-02"`, `"2006-01-02 15:04:05"`), never
  hand-rolled formatting.
- `ParseDate` returns a non-nil error for any input outside the exact layout.
- `BusinessDays` counts with `Weekday()`; Saturday and Sunday never count.
- No external dependencies: standard library only.

## Time limit

**25 minutes** (round 1). Beat it on every repeat: aim for under 12 minutes
by round 4.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference
solution, turn off autocomplete if you can. Write every line from memory —
that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 145
```

or run the tests directly:

```bash
go test ./145-time/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 145` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
