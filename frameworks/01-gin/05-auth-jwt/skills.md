# Skills for gin/05-auth-jwt

## What You'll Learn

**Previous:** [gin/04-json-crud-api](../04-json-crud-api/skills.md) | **Next:** [gin/06-testing-gin](../06-testing-gin/skills.md)

**Challenge:** Implement JWT authentication with password hashing, token signing, and protected middleware.

## Password Hashing with bcrypt

Never store plain-text passwords. Use bcrypt:

```go
import "golang.org/x/crypto/bcrypt"

// Hash
hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// Verify
err := bcrypt.CompareHashAndPassword(hash, []byte(password))
// err == nil means password matches
```

`bcrypt.DefaultCost` (10) balances security and speed. Higher cost = slower = harder to brute-force.

## JWT Basics

A JWT is three base64-encoded parts separated by `.`:

```
header.payload.signature
```

- **Header**: algorithm (`HS256`)
- **Payload**: claims (`user_id`, `exp`, custom fields)
- **Signature**: HMAC-SHA256 of header+payload with a secret key

The server validates the signature — if it matches, the payload is authentic.

## Creating a JWT

```go
import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

func createToken(userID int, secret string) (string, error) {
    claims := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
```

## Validating a JWT

```go
func parseToken(tokenStr, secret string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
        func(t *jwt.Token) (interface{}, error) {
            if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
            return []byte(secret), nil
        })
    if err != nil { return nil, err }
    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid { return nil, fmt.Errorf("invalid token") }
    return claims, nil
}
```

## JWT Auth Middleware

```go
func JWTAuth(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if !strings.HasPrefix(auth, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized,
                gin.H{"error": "missing token"})
            return
        }
        tokenStr := strings.TrimPrefix(auth, "Bearer ")
        claims, err := parseToken(tokenStr, secret)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized,
                gin.H{"error": "invalid token"})
            return
        }
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}

// In a protected handler:
userID := c.GetInt("user_id")
```

## Environment Variables

```go
import "os"

secret := os.Getenv("JWT_SECRET")
if secret == "" {
    secret = "dev-secret-change-in-production"
}
```

Set with: `JWT_SECRET=mykey go run .`

## Security Notes

- Use at least 32 random bytes for the JWT secret in production
- Never log tokens
- Use HTTPS in production (JWT payload is base64-encoded, not encrypted)
- Set short expiry (1h for APIs, 24h max)
- Implement token refresh if needed

## Documentation

- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [JWT.io — debugger and spec](https://jwt.io/)
- [OWASP JWT cheat sheet](https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_for_Java_Cheat_Sheet.html)

**Next:** [gin/06-testing-gin](../06-testing-gin/README.md)
