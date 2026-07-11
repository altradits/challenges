# gin/03-middleware — Custom Middleware

## Challenge

Add the following middleware to your Gin server:

1. **Request Logger** — logs method, path, status code, and latency for every request
2. **CORS** — adds `Access-Control-Allow-Origin: *` headers for all responses
3. **RateLimit** — allows a maximum of 5 requests per second per IP (in-memory, approximate)
4. **Auth guard** — protects routes under `/api/v1/protected/*` with an API key header `X-API-Key: secret123`

```
GET /public      → always accessible
GET /api/v1/protected/data → requires X-API-Key header
```

## Test

```bash
# Public — works
curl http://localhost:8080/public

# Protected — fails without key
curl http://localhost:8080/api/v1/protected/data
# → 401

# Protected — works with key
curl -H "X-API-Key: secret123" http://localhost:8080/api/v1/protected/data
# → 200
```

Read `skills.md` before you start.
