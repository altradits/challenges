# Prerequisites for 32-errorhandling

## Before You Start

- [06-functions](../06-functions/skills.md) — multiple return values
- [31-interfaces](../31-interfaces/skills.md) — interfaces (error is an interface)
- [58-itoa](../58-itoa/skills.md) — strconv.Atoi (used in the challenge)

## You're Ready When You Can...

- [ ] Return `(value, error)` from a function
- [ ] Check `if err != nil` immediately after a call
- [ ] Use `errors.New` and `fmt.Errorf("%w", err)` to create errors
- [ ] Use `errors.Is` to check for a specific sentinel error

## Next Steps

- [18-typeassertions](../18-typeassertions/README.md)
