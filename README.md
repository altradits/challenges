# Go Mastery — Code Go From Memory, Then Ship Bitcoin Open Source

A complete, Go-only learning path: **numbered challenge directories** from
Hello World to a merged PR in btcd or LND — trained the way memory actually
works: **timed, blind, repeated**.

Follow the numbers. Read `skills.md` first, then solve the challenge in
`README.md` **from memory, against the clock, without looking anything up**.

---

## The Method: Why From Memory

Reading Go and writing Go are different skills. Open-source work — reading
LND, fixing a bug, defending a PR in review — happens with no safety net: you
either know the language cold or you stall. This course trains recall, not
recognition:

1. **Learn** — read `skills.md`, work the examples with the book open.
2. **Close everything** — start the drill timer. No docs, no search, no AI, no
   copying. Editor autocomplete off if you can.
3. **Write from memory** — every line from your head, against the challenge’s
   time limit.
4. **Check without looking up** — run the tests. Failure messages tell you
   what is wrong; your memory tells you how to fix it.
5. **Repeat on schedule** — redo each challenge from scratch after **1 day, 3
   days, 7 days, 21 days**, beating your previous time each round. That is
   spaced repetition: the schedule that moves syntax from short-term memory
   into permanent skill.

A challenge is done when you can solve it blind, inside the time limit, on the
21-day repeat.

## Quick Start

Prerequisites: Go **1.26.10** (matches LND’s `go.mod`), offline-friendly —
everything except the frameworks track uses the standard library only.

```bash
git clone https://github.com/altradits/challenges.git
cd challenges

# 1. Learn with the book open
cat 00-keywords/README.md
cat 06-functions/skills.md

# 2. Drill blind against the clock
go run ./cmd/drill start 06

# 3. Write your solution from memory, then check it
go run ./cmd/drill check 06

# 4. Come back tomorrow: see what is due, watch your streak
go run ./cmd/drill due
go run ./cmd/drill stats
```

No network, no accounts. Progress lives in `.drill/progress.json` (gitignored,
plain JSON).

## The Runner: `drill`

`cmd/drill` (standard library only) enforces the method:

| Command | What it does |
|---------|--------------|
| `drill start NN` | Starts the timer, shows **only the task** — skills and solutions stay hidden |
| `drill check NN` | Runs the challenge tests, records pass/fail and your time |
| `drill due` | Lists what spaced repetition says to redo today |
| `drill stats` | Streaks, best times, and your weak topics |
| `drill solution NN` | Shows the reference solution — only after you have passed NN once |

Manual challenges (00, 22, 83, 158) have no hidden tests: `drill check`
walks a self-attestation checklist instead. Be honest; you are only cheating
your own memory.

## Course Structure

| Phase | Directories | What You Learn |
|-------|-------------|----------------|
| Step 0 | 00 | The 25 keywords — the whole language surface on day one |
| 1 | 01–05 | Hello World — the four-line Go skeleton |
| 2 | 06–27 | Foundations — functions, loops, slices, maps, structs, pointers, methods, interfaces, errors, files, regexp, sort, modules, testing, goroutines, channels, sync, context |
| 3 | 28–51 | Beginner practice — one concept per exercise, build speed |
| 4 | 52–80 | Strings package mastery — every `strings`/`fmt`/`strconv` function |
| 5 | 81–144 | Challenge problems — piscine-style, multiple concepts combined |
| 6 | 145–158 | Backend bridge (time, json, http, sql, config, logging, generics, graceful shutdown) + capstones + first Bitcoin PR |
| 7 | 159–169 | Language depth — closures, defer/panic/recover, switch, variadic, conversions, iota, embedding, generics, io, mutexes, advanced testing |
| 8 | 176–185 | Data structures & algorithms — linked list, stack, queue, binary search, sorting, trees, graphs, DP, two pointers, sliding window |
| 9 | 170–175, 186+ | Bitcoin protocol — byte order, varints, hex, SHA-256d, Merkle root, block headers, TCP message codec |

See [LEARNING_PATH.md](LEARNING_PATH.md) for the full table with every
directory listed, and [docs/CHALLENGE_TEMPLATE.md](docs/CHALLENGE_TEMPLATE.md)
for the template every challenge follows (Task, Input/Output, Examples,
Constraints, Time limit, Blind rule, How to check, Repeat schedule).

The `frameworks/` track (Gin, Echo, Fiber, GORM, Chi) is a separate,
network-dependent extra after the core path — it is not part of the numbered
drills or the offline test suite.

## How to Use Each Directory

Every numbered directory has the same shape:

- **`README.md`** — the challenge: task, examples, constraints, time limit
- **`skills.md`** — the concepts you need: read this before you start
- **`prerequisites.md`** — what to review if you get stuck
- **`*_test.go`** — hidden tests with clear failure messages: the judge
- *You write the solution file* — nothing else goes in the directory

Reference solutions live separately under `solutions/` (one directory per
challenge). The runner keeps them hidden until you pass once. Do not open
them early — every peek is a rep stolen from your memory.

Work the pattern: `skills.md` → `drill start NN` → solve blind →
`drill check NN` → repeat on schedule.

## How to Track Progress

- `drill due` — your daily list, computed from the 1/3/7/21-day schedule.
- `drill stats` — day streak, per-challenge best times vs target, topics with
  the most failures (your weak list — drill those first).
- Raw data: `.drill/progress.json`. Back it up, delete it to restart, diff it
  to see yourself improve. It is never committed.

## Companion Repositories

This repo contains only Go. Other languages live in their own repos:

- Python data analysis: [altradits/python-data](https://github.com/altradits/python-data)
- JavaScript frontend + database: [altradits/js-mastery](https://github.com/altradits/js-mastery)
- Career development + LinkedIn: [altradits/developer-career](https://github.com/altradits/developer-career)

## Checkpoint Tiers

See [CHECKPOINT_TIERS.md](CHECKPOINT_TIERS.md) for which challenges test you
at each 5%–100% proficiency tier.

## Contributing

One-line commit messages, tests with every challenge, and your own work only.
Full rules: [CONTRIBUTING.md](CONTRIBUTING.md).

---

**Start at `00-keywords`. Follow the numbers. Code it from memory.**
