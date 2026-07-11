# fiber/03-fiber-middleware — Fiber Middleware

## Challenge

Add middleware to your Product API:

1. **Logger** — use `github.com/gofiber/fiber/v2/middleware/logger`
2. **CORS** — allow all origins
3. **Recover** — convert panics to 500 responses
4. **RequestID** — add `X-Request-ID` header
5. **Compress** — gzip compress responses > 1KB
6. Custom **API version header** middleware — adds `X-API-Version: v1` to all responses

```bash
go get github.com/gofiber/fiber/v2
# middleware is included in the package
```

Read `skills.md` before you start.
