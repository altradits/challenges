package piscine

// This file is part of the hidden tests for 148-database-sql. It registers an
// in-memory SQL driver under the name "sqlite" so the challenge runs fully
// offline with the standard library only. In real use, swap the driver name
// for modernc.org/sqlite (see skills.md): the SQL in OpenDB works unchanged.

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"
)

// stubStore is one in-memory users table, keyed by DSN so each test database
// is isolated.
type stubStore struct {
	mu     sync.Mutex
	nextID int64
	users  map[int64]stubUser
}

type stubUser struct {
	id          int64
	name, email string
}

var (
	stubMu     sync.Mutex
	stubStores = map[string]*stubStore{}
)

func stubStoreFor(dsn string) *stubStore {
	stubMu.Lock()
	defer stubMu.Unlock()
	st, ok := stubStores[dsn]
	if !ok {
		st = &stubStore{users: map[int64]stubUser{}}
		stubStores[dsn] = st
	}
	return st
}

type stubDriver struct{}

func (stubDriver) Open(dsn string) (driver.Conn, error) {
	return &stubConn{store: stubStoreFor(dsn)}, nil
}

type stubConn struct{ store *stubStore }

func (c *stubConn) Prepare(query string) (driver.Stmt, error) {
	return &stubStmt{store: c.store, query: query}, nil
}

func (c *stubConn) Close() error              { return nil }
func (c *stubConn) Begin() (driver.Tx, error) { return stubTx{}, nil }
func (stubTx) Commit() error                  { return nil }
func (stubTx) Rollback() error                { return nil }

type stubTx struct{}

type stubStmt struct {
	store *stubStore
	query string
}

func (s *stubStmt) Close() error { return nil }

func (s *stubStmt) NumInput() int { return -1 }

func (s *stubStmt) Exec(args []driver.Value) (driver.Result, error) {
	q := strings.ToUpper(strings.TrimSpace(s.query))
	st := s.store
	st.mu.Lock()
	defer st.mu.Unlock()
	switch {
	case strings.HasPrefix(q, "CREATE TABLE"):
		return stubResult{}, nil
	case strings.HasPrefix(q, "INSERT INTO"):
		st.nextID++
		u := stubUser{id: st.nextID}
		if len(args) > 0 {
			u.name, _ = args[0].(string)
		}
		if len(args) > 1 {
			u.email, _ = args[1].(string)
		}
		st.users[u.id] = u
		return stubResult{id: u.id, n: 1}, nil
	case strings.HasPrefix(q, "DELETE FROM"):
		if len(args) == 0 {
			return stubResult{}, nil
		}
		id, _ := args[0].(int64)
		if _, ok := st.users[id]; !ok {
			return stubResult{}, nil
		}
		delete(st.users, id)
		return stubResult{n: 1}, nil
	default:
		return nil, fmt.Errorf("stub: unsupported exec %q", s.query)
	}
}

func (s *stubStmt) Query(args []driver.Value) (driver.Rows, error) {
	q := strings.ToUpper(strings.TrimSpace(s.query))
	st := s.store
	st.mu.Lock()
	defer st.mu.Unlock()
	if !strings.HasPrefix(q, "SELECT") {
		return nil, fmt.Errorf("stub: unsupported query %q", s.query)
	}
	var rows [][]driver.Value
	if strings.Contains(q, "WHERE") {
		if len(args) == 0 {
			return &stubRows{}, nil
		}
		id, _ := args[0].(int64)
		if u, ok := st.users[id]; ok {
			rows = append(rows, []driver.Value{u.id, u.name, u.email})
		}
		return &stubRows{rows: rows}, nil
	}
	for id := int64(1); id <= st.nextID; id++ {
		if u, ok := st.users[id]; ok {
			rows = append(rows, []driver.Value{u.id, u.name, u.email})
		}
	}
	return &stubRows{rows: rows}, nil
}

type stubResult struct {
	id int64
	n  int64
}

func (r stubResult) LastInsertId() (int64, error) { return r.id, nil }
func (r stubResult) RowsAffected() (int64, error) { return r.n, nil }

type stubRows struct {
	rows [][]driver.Value
	pos  int
}

func (r *stubRows) Columns() []string { return []string{"id", "name", "email"} }
func (r *stubRows) Close() error      { return nil }
func (r *stubRows) Next(dest []driver.Value) error {
	if r.pos >= len(r.rows) {
		return io.EOF
	}
	copy(dest, r.rows[r.pos])
	r.pos++
	return nil
}

func TestMain(m *testing.M) {
	sql.Register("sqlite", stubDriver{})
	m.Run()
}

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := OpenDB(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatalf("OpenDB: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	if err := db.PingContext(context.Background()); err != nil {
		t.Fatalf("ping: %v", err)
	}
	return db
}

func TestUserCRUD(t *testing.T) {
	db := openTestDB(t)
	id1, err := InsertUser(db, "Ada", "ada@example.com")
	if err != nil {
		t.Fatalf("InsertUser: %v", err)
	}
	id2, err := InsertUser(db, "Bo", "bo@example.com")
	if err != nil {
		t.Fatalf("InsertUser: %v", err)
	}
	if id1 == id2 {
		t.Fatalf("ids must differ: %d %d", id1, id2)
	}
	u, err := GetUser(db, id1)
	if err != nil {
		t.Fatalf("GetUser: %v", err)
	}
	if u.ID != id1 || u.Name != "Ada" || u.Email != "ada@example.com" {
		t.Errorf("GetUser = %+v", u)
	}
	all, err := ListUsers(db)
	if err != nil {
		t.Fatalf("ListUsers: %v", err)
	}
	if len(all) != 2 || all[0].ID != id1 || all[1].ID != id2 {
		t.Errorf("ListUsers = %+v, want ids [%d %d] in order", all, id1, id2)
	}
	if err := DeleteUser(db, id1); err != nil {
		t.Fatalf("DeleteUser: %v", err)
	}
	if _, err := GetUser(db, id1); err != sql.ErrNoRows {
		t.Errorf("GetUser after delete err = %v, want sql.ErrNoRows", err)
	}
	all, err = ListUsers(db)
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 1 || all[0].ID != id2 {
		t.Errorf("ListUsers after delete = %+v", all)
	}
}

func TestUserNotFound(t *testing.T) {
	db := openTestDB(t)
	if _, err := GetUser(db, 999); err != sql.ErrNoRows {
		t.Errorf("GetUser missing err = %v, want sql.ErrNoRows", err)
	}
	if err := DeleteUser(db, 999); err == nil {
		t.Error("DeleteUser of missing row should return an error")
	}
}
