# Skills for 37-modules-and-tooling

## What You'll Learn

**Previous:** [21-sort-and-math](../21-sort-and-math/skills.md) | **Next:** [23-testing-basics](../23-testing-basics/skills.md)

**Reference:** Go modules, dependency management, and essential CLI tools.

## Go Modules

A module is a collection of related Go packages with a shared version.

### Creating a Module

```bash
mkdir myapp && cd myapp
go mod init github.com/yourusername/myapp
```

This creates `go.mod`:
```
module github.com/yourusername/myapp

go 1.22
```

### `go.mod` Structure

```
module github.com/yourusername/myapp   ← module path (matches GitHub URL)

go 1.22                                 ← minimum Go version

require (
    github.com/stretchr/testify v1.8.4  ← direct dependency
)
```

### Adding Dependencies

```bash
go get github.com/stretchr/testify@latest  # add + download
go get github.com/some/pkg@v1.2.3          # pin to specific version
```

### Removing Unused Dependencies

```bash
go mod tidy  # removes unused, adds missing
```

Run `go mod tidy` before every commit.

### `go.sum`

Auto-generated checksum file. Never edit it manually. Commit it to version control.

## Essential CLI Tools

### `go run` — compile and run in one step

```bash
go run main.go
go run ./...      # run all packages
```

### `go build` — compile to a binary

```bash
go build ./...         # build all packages (check for errors)
go build -o myapp ./cmd/myapp  # build to specific output
```

### `go fmt` — format code

```bash
gofmt -w .      # format and write all .go files
go fmt ./...    # equivalent
```

Go has one canonical format. Run this before every commit.

### `go vet` — static analysis

```bash
go vet ./...
```

Catches common mistakes: unreachable code, bad format strings, locking errors.

### `go doc` — documentation

```bash
go doc strings.Contains     # docs for a function
go doc strings               # docs for a package
go doc -all fmt              # all exported symbols
```

### `go test` (preview — taught in depth in 38-testing-basics)

```bash
go test ./...           # run all tests
go test -v ./...        # verbose output
go test -run TestFoo    # run specific test
```

### `go install` — install a binary globally

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Standard Directory Layout

```
myapp/
├── cmd/
│   └── myapp/
│       └── main.go      ← entry point
├── internal/
│   └── mypackage/
│       └── logic.go     ← private packages (can't be imported by other modules)
├── pkg/
│   └── shared/
│       └── shared.go    ← public packages
├── go.mod
├── go.sum
└── README.md
```

`cmd/` contains your binary entry points.
`internal/` enforces package privacy at the module boundary.
`pkg/` contains reusable public packages.

## golangci-lint

A fast linter that runs many checks:

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run ./...
```

Common linters it runs: `errcheck` (unchecked errors), `gosimple`, `staticcheck`, `unused`.

## Special Patterns

### `init()`

Runs automatically before `main()` or before the first use of a package:

```go
var config Config

func init() {
    config = loadConfig()  // runs once at package initialization
}
```

Don't overuse. Prefer explicit initialization.

### Blank Identifier `_`

Discards a value you're forced to receive but don't need:

```go
_, err := fmt.Fprintln(os.Stderr, "error")  // discard byte count

for _, v := range slice { ... }  // discard index

import _ "net/http/pprof"  // side-effect import (registers handlers)
```

**Next:** [23-testing-basics](../23-testing-basics/README.md)
