# Prerequisites for gin/01-hello-gin

## Before You Start

- [147-http-basics](../../../147-http-basics/README.md) — understanding HTTP and net/http
- [146-encoding-json](../../../146-encoding-json/README.md) — JSON in Go
- [22-modules-and-tooling](../../../22-modules-and-tooling/README.md) — `go mod`, `go get`

## Install Gin

```bash
mkdir myproject && cd myproject
go mod init myproject
go get github.com/gin-gonic/gin
```

## You're Ready When You Can...

- [ ] Run `go run .` and get a running program
- [ ] Explain what `http.HandlerFunc` is
- [ ] Encode a struct as JSON with `encoding/json`
- [ ] Use `go get` to add a dependency

## Next Steps

- [gin/02-routing](../02-routing/README.md)
