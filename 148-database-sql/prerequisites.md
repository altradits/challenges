# Prerequisites for 148-database-sql

## Before You Start

- [13-structs](../13-structs/skills.md) — rows.Scan into struct fields
- [17-errorhandling](../17-errorhandling/skills.md) — sql.ErrNoRows is a sentinel error
- [19-fileio](../19-fileio/skills.md) — defer f.Close() pattern → defer rows.Close()
- [22-modules-and-tooling](../22-modules-and-tooling/skills.md) — go get modernc.org/sqlite
- [27-context](../27-context/skills.md) — QueryContext, ExecContext

## You're Ready When You Can...

- [ ] Open a database with `sql.Open` and verify with `db.Ping`
- [ ] Execute INSERT/UPDATE/DELETE with `db.Exec` and `?` placeholders
- [ ] Query a single row with `db.QueryRow(...).Scan`
- [ ] Iterate multiple rows with `rows.Next()`, `rows.Scan()`, `rows.Close()`
- [ ] Wrap multiple statements in a transaction with `db.Begin`, `tx.Commit`, `defer tx.Rollback`

## Next Steps

- [149-environment-and-config](../149-environment-and-config/README.md)
