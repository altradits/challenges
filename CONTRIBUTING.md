# Contributing to Go Mastery Challenges

Thanks for helping make this the most complete “code Go from memory” course.
Please read this whole file before opening a PR.

## Ground rules

- **Code, tests and translations must be your own work.** Do not paste
  solutions from other courses, books, or AI output and call them yours.
- No signing key required. Plain commits are fine.
- One-line commit messages, plain text, no AI attribution. Example:
  `test(96-gcd): cover zero and equal inputs`
- Small, reviewable PRs: one topic per branch, rebase on `main` before asking
  for review. Never merge, push to `main`, or force-push; a maintainer merges.

## Pinned toolchain

- Go **1.26.10** (matches LND’s `go.mod`). Root module and every nested module
  declare `go 1.26.10`.
- CI runs `gofmt -l`, `go vet ./...`, and the full test suite with
  `-race -cover` on that exact version. Run the same locally before you push:
  `gofmt -l .`, `go vet ./...`, `go test -race -cover ./...` (after overlaying
  solutions, see below).

## How to add a challenge with tests

1. Pick the next free number in the right phase (see `LEARNING_PATH.md`). Never
   renumber an existing challenge; if you must, keep an old→new map in the PR.
2. Create `NN-name/` with exactly:
   - `README.md` — follow [docs/CHALLENGE_TEMPLATE.md](docs/CHALLENGE_TEMPLATE.md)
     (Task, Input/Output, Examples, Constraints, Time limit, Blind rule, How
     to check, Repeat schedule).
   - `skills.md` — concepts only, no solution code.
   - `prerequisites.md` — what to review when stuck.
   - `NN_name_test.go` — table-driven `go test` with clear failure messages
     and edge cases. Function challenges: `package piscine`. Program
     challenges (read `os.Args`, write stdout): `package main`, build the
     binary and compare exact bytes.
3. Add the reference solution under `solutions/NN-name/` (same package as the
   student file). The runner keeps it hidden until the student passes once:
   never link or import it from the challenge directory.
4. Verify with the overlay (this is what CI does):
   `scripts/overlay-solutions.sh && go test ./NN-name/ && git checkout -- NN-name/`
   (the script copies each reference solution into its challenge directory so
   the tests run against a known-good implementation, then you revert the
   overlay).
5. Update `LEARNING_PATH.md` (phase table + quick reference) and the phase
   table in `README.md`.

## Style

- `gofmt` clean, `go vet` clean, standard library only for the root module and
  the `drill` runner. New external dependencies need a written reason and
  belong in a nested module, never the root.
- Tests: table-driven, one case per row, `%q`-style diffs that show expected
  vs actual without opening any source.

## Contact

- Maintainer: Stanley (Altradits)
- Phone: 0707172370
- Email: 5tan@altradits.com
