# Prerequisites for 147-http-basics

## Before You Start

- [16-interfaces](../16-interfaces/skills.md) — http.Handler is an interface
- [17-errorhandling](../17-errorhandling/skills.md) — handlers return errors via HTTP status codes
- [146-encoding-json](../146-encoding-json/skills.md) — JSON in every handler
- [27-context](../27-context/skills.md) — http.NewRequestWithContext uses context

## You're Ready When You Can...

- [ ] Write an `http.HandlerFunc` that sets headers and encodes JSON
- [ ] Register routes on an `http.ServeMux`
- [ ] Read query parameters from `r.URL.Query()`
- [ ] Decode a JSON request body with `json.NewDecoder(r.Body).Decode`
- [ ] Make an HTTP GET request with `http.Get` and decode the response

## Next Steps

- [148-database-sql](../148-database-sql/README.md)
