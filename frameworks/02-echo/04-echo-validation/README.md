# echo/04-echo-validation — Input Validation

## Challenge

Add a custom validator to your Echo application:

```go
type RegisterInput struct {
    Username  string `json:"username" validate:"required,min=3,max=20,alphanum"`
    Email     string `json:"email"    validate:"required,email"`
    Password  string `json:"password" validate:"required,min=8"`
    Age       int    `json:"age"      validate:"min=18,max=120"`
    Role      string `json:"role"     validate:"oneof=admin user moderator"`
}
```

1. Wire up the `go-playground/validator` with Echo's `Validator` interface
2. Return structured validation errors: `{ "errors": [{ "field": "email", "message": "must be a valid email" }] }`
3. Add a custom validator rule: `no_spaces` — fails if the string contains spaces
4. Add a custom validator: `strong_password` — requires at least one uppercase, one digit, one special character

## Test

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"al","email":"not-an-email","password":"weak"}'

# Response:
{
  "errors": [
    {"field": "username", "message": "must be at least 3 characters"},
    {"field": "email", "message": "must be a valid email"},
    {"field": "password", "message": "must be at least 8 characters"}
  ]
}
```

Read `skills.md` before you start.
