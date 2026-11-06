# 22-modules-and-tooling — Go Modules and Tooling

## No Code Challenge

This is a study-and-practice lesson. Work through the tasks in `skills.md` in your terminal.

## Tasks

1. Create a new Go module in a temp directory: `go mod init github.com/yourname/demo`
2. Add a dependency: `go get github.com/stretchr/testify@latest`
3. Write a Go file that imports the dependency
4. Run `go mod tidy` and inspect `go.sum`
5. Run `go fmt ./...`, `go vet ./...`, `go build ./...`
6. Use `go doc strings.Contains` to read documentation in the terminal

Read `skills.md` for the full reference.

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

This is a manual challenge — there are no hidden tests. Run:

```bash
go run ./cmd/drill check 22
```

`drill check` walks a self-attestation checklist and records an honest
pass. Be honest; you are only cheating your own memory.

## Repeat schedule

No 1/3/7/21-day repeats — this is reference material or terminal
practice. Revisit whenever `prerequisites.md` of a later challenge points
back here.
