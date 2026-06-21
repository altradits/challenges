# Go Mastery — From Hello World to Bitcoin Open Source

A complete, Go-only learning path: **157 numbered directories**, each one building on the last.

Follow the numbers. Read `skills.md` first, then solve the challenge in `README.md`.

---

## Quick Start

```bash
git clone https://github.com/altradits/challenges.git
cd challenges
cat 01-only1/README.md
```

---

## Course Structure

| Phase | Directories | What You Learn |
|-------|------------|----------------|
| 1 | 01–05 | Hello World — the four-line Go skeleton |
| 2 | 06–27 | Foundations — every core Go concept: functions, loops, structs, pointers, interfaces, errors, file I/O, regexp, sort, modules, testing, goroutines, channels, sync, context |
| 3 | 28–51 | Beginner practice — one concept per exercise, build speed |
| 4 | 52–80 | Strings package mastery — every `strings`/`fmt`/`strconv` function |
| 5 | 81–144 | Challenge problems — piscine-style, multiple concepts combined |
| 6 | 145–157 | Backend bridge (time, json, http, sql, config, logging, generics) + Capstone projects |

See [LEARNING_PATH.md](LEARNING_PATH.md) for the full table with every directory listed.

---

## How to Use Each Directory

Every numbered directory has three files:

- **`README.md`** — the challenge: what to build and what it must do
- **`skills.md`** — the concepts you need: read this before you start
- **`prerequisites.md`** — what to review if you get stuck

Work the pattern: `skills.md` → attempt the challenge → check `prerequisites.md` if stuck.

---

## What's New in Phase 6 (145–151): Backend Bridge

These 7 lessons teach every backend-essential package before the capstone projects:

| Directory | Teaches |
|-----------|---------|
| `145-time` | `time.Time`, `Duration`, Parse/Format, Since/Until, weekday logic |
| `146-encoding-json` | `json.Marshal/Unmarshal`, struct tags, `omitempty`, Encoder/Decoder |
| `147-http-basics` | `net/http` handlers, `ServeMux`, request/response, `http.Client` |
| `148-database-sql` | `database/sql`, SQLite, QueryRow/Query, transactions, connection pool |
| `149-environment-and-config` | `os.Getenv`, `flag` package, 12-factor config struct pattern |
| `150-logging` | `log` package, `slog` structured logging (Go 1.21+), middleware |
| `151-generics` | Type parameters `[T any]`, `comparable`, union constraints `~T` |

---

## What's New in Phase 2 (13–27)

These foundation lessons cover every concept a Go developer needs before tackling real problems:

| Directory | Teaches |
|-----------|---------|
| `13-structs` | Struct types, field access, constructors |
| `14-pointers` | `&`, `*`, mutation, nil pointers |
| `15-methods` | Value vs pointer receivers |
| `16-interfaces` | Implicit satisfaction, `Stringer`, `io.Reader` |
| `17-errorhandling` | `errors.Is/As`, wrapping with `%w`, `panic/recover` |
| `18-typeassertions` | `.(T)`, type switches, `any` |
| `19-fileio` | `os.ReadFile/WriteFile`, `bufio.Scanner`, `io.Reader/Writer` |
| `20-regexp` | Pattern matching, extract, replace |
| `21-sort-and-math` | `sort.Slice`, `math.Sqrt/Pow/Pi` |
| `22-modules-and-tooling` | `go mod`, `go fmt/vet/doc`, `golangci-lint`, `init()` |
| `23-testing-basics` | Table-driven tests, `t.Run`, benchmarks |
| `24-goroutines` | `go` keyword, `sync.WaitGroup`, race detector |
| `25-channels` | Channel operations, pipeline pattern, worker pool |
| `26-select-and-sync` | `select`, `sync.Mutex`, `sync.Once`, `time.After` |
| `27-context` | Cancellation, deadlines, `context.WithTimeout` |

---

## Companion Repositories

This repo contains only Go. Other languages live in their own repos:
- Python data analysis: [altradits/python-data](https://github.com/altradits/python-data)
- JavaScript frontend + database: [altradits/js-mastery](https://github.com/altradits/js-mastery)
- Career development + LinkedIn: [altradits/developer-career](https://github.com/altradits/developer-career)

---

## Checkpoint Tiers

See [CHECKPOINT_TIERS.md](CHECKPOINT_TIERS.md) for which challenges test you at each 5%–100% proficiency tier.

---

**Start at `01-only1`. Follow the numbers to `157-bitcoin-oss`.**

```
01–05   Hello World     →  The four-line Go skeleton
06–27   Foundations     →  Every core Go concept
28–51   Beginner        →  One skill per exercise
52–80   Strings         →  Every strings/fmt/strconv function
81–144  Challenges      →  Hard piscine-style problems
145–151 Backend bridge  →  time, json, http, sql, config, logging, generics
152–157 Capstones       →  Real Go APIs + Bitcoin open source
```
