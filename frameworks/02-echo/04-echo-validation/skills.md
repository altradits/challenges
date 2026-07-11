# Skills for echo/04-echo-validation

## What You'll Learn

**Previous:** [echo/03-echo-middleware](../03-echo-middleware/skills.md) | **Next:** [echo/05-echo-auth](../05-echo-auth/skills.md)

**Challenge:** Wire go-playground/validator into Echo and create custom validation rules.

## Echo's Validator Interface

Echo has a `Validator` interface with a single method:

```go
type Validator interface {
    Validate(i interface{}) error
}
```

Implement it and assign to `e.Validator`:

```go
import "github.com/go-playground/validator/v10"

type CustomValidator struct {
    v *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.v.Struct(i)
}

// In main:
e.Validator = &CustomValidator{v: validator.New()}
```

Now `c.Bind(&input)` automatically validates — if validation fails, `Bind` returns an error.

## Using Validate Separately

```go
func register(c echo.Context) error {
    var input RegisterInput
    if err := c.Bind(&input); err != nil {
        return err
    }
    if err := c.Validate(input); err != nil {
        return err  // Echo passes this to the error handler
    }
    // input is valid
}
```

## Structured Validation Errors

```go
import "github.com/go-playground/validator/v10"

func formatValidationErrors(err error) []map[string]string {
    var errs []map[string]string
    for _, e := range err.(validator.ValidationErrors) {
        errs = append(errs, map[string]string{
            "field":   e.Field(),
            "message": validationMessage(e),
        })
    }
    return errs
}

func validationMessage(e validator.FieldError) string {
    switch e.Tag() {
    case "required":
        return "is required"
    case "min":
        return fmt.Sprintf("must be at least %s characters", e.Param())
    case "email":
        return "must be a valid email"
    case "oneof":
        return "must be one of: " + e.Param()
    default:
        return "is invalid"
    }
}
```

## Custom Validators

```go
v := validator.New()

// Register custom rule
v.RegisterValidation("no_spaces", func(fl validator.FieldLevel) bool {
    return !strings.Contains(fl.Field().String(), " ")
})

v.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
    p := fl.Field().String()
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(p)
    hasDigit := regexp.MustCompile(`[0-9]`).MatchString(p)
    hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(p)
    return hasUpper && hasDigit && hasSpecial
})
```

Use with struct tag:

```go
type Input struct {
    Username string `validate:"required,no_spaces"`
    Password string `validate:"required,strong_password"`
}
```

## Documentation

- [go-playground/validator](https://github.com/go-playground/validator)
- [Echo custom validator](https://echo.labstack.com/docs/request#validate-data)
- [Validator documentation](https://pkg.go.dev/github.com/go-playground/validator/v10)

**Next:** [echo/05-echo-auth](../05-echo-auth/README.md)
