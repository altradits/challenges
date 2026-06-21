# 37-modules-and-tooling — Go Modules and Tooling

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
