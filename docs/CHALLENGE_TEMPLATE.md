# Challenge README Template

Every numbered challenge directory keeps the same three files (`README.md`,
`skills.md`, `prerequisites.md`) plus a test file. This document defines the
one template every challenge `README.md` follows.

Read `skills.md` to learn, then solve the challenge in `README.md` blind.

## The template

```markdown
# NN-name — Title

## Task

One paragraph: what to build, in which package (`package piscine` for
functions, `package main` for programs), and the exact names the tests call.

## Input / Output

How the solution receives input and produces output (function arguments and
return values, or program arguments and stdout).

## Examples

Concrete input → output pairs. Copy the exact values into your head, then
close the page.

## Constraints

Edge cases, limits, forbidden shortcuts (`len`, `sort.Slice`, …), performance
notes.

## Time limit

**M minutes** (round 1). Beat it on every repeat: each spaced round targets a
shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference
solution, turn off autocomplete if you can. Write every line from memory —
that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check NN

or run the tests directly:

    go test ./NN-name/

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start NN` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
```

## Time limits (round 1)

| Numbers | Phase | Minutes |
|---------|-------|---------|
| 00 | Step 0 (keywords preface, manual) | 15 |
| 01–05 | Phase 1 (hello world) | 5 |
| 06–27 | Phase 2 (foundations) | 15 |
| 28–51 | Phase 3 (beginner practice) | 10 |
| 52–80 | Phase 4 (strings mastery) | 10 |
| 81–140 | Phase 5 (challenge problems) | 20 |
| 141–144 | Phase 5 (hard: stack, RPN, brainfuck, grouping) | 30 |
| 145–152 | Phase 6 (backend bridge) | 25 |
| 153–155 | Phase 6 (capstone services) | 120 |
| 156 | Phase 6 (API course: each sub-exercise) | 45 |
| 157 | Phase 6 (analytics capstone) | 240 |
| 158 | Phase 6 (first merged Bitcoin PR, manual) | — |
| 159–175 | Phase 7 (language depth) | 20 |
| 176–185 | Phase 8 (data structures & algorithms) | 30 |
| 186–191 | Phase 9 (Bitcoin protocol) | 30 |
| 192 | Phase 9 (tooling practice, manual) | 30 |

Manual challenges (00, 22, 83, 158, 192) replace “How to check” with a
self-attestation checklist: `drill check NN` walks it and only records a pass
on an honest yes.

## Rules for authors

- The `## Task` names must match the test file exactly: same function names,
  same signatures, same program arguments and output bytes.
- Examples in the README must be a subset of the cases in the test file.
- Constraints must list every edge case the tests cover (empty input, zero,
  negative, not-found, …).
- Time limit comes from the table above unless the challenge has a written
  reason to differ; note the reason in the PR.
- Never paste the solution (or a near-solution) into `README.md`, `skills.md`,
  or `prerequisites.md`. Reference solutions live only under `solutions/`.
