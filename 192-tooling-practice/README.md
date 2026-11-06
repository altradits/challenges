# 192-tooling-practice — Race, Fuzz, Vet, and Profile

## No Code Challenge

This is a manual challenge — there are no hidden tests. Work through the
tasks in `skills.md` in your terminal, using code you already wrote.

## Tasks

1. Run your `174-worker-pool` and `175-sync-atomic` solutions under the
   race detector: `go test -race ./174-worker-pool/ ./175-sync-atomic/`
2. Write a fuzz test for `184-two-pointers` proving
   `Reverse(Reversed(x)) == x`, seed it, and run `go test -fuzz` for
   60 seconds with zero crashes
3. Benchmark `180-sorting-algorithms` (`-bench . -benchtime 100x`),
   capture `-cpuprofile`, and read the top 5 functions with
   `go tool pprof -top`
4. Run `go vet ./...` from the repo root and explain every finding in
   your own words (or confirm zero findings)

Read `skills.md` for the full reference.

---

## Time limit

**30 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

This is a manual challenge — there are no hidden tests. Run:

```bash
go run ./cmd/drill check 192
```

`drill check` walks a self-attestation checklist and records an honest
pass. Be honest; you are only cheating your own memory.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
