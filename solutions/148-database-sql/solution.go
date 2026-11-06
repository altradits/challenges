package piscine

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	)`)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func InsertUser(db *sql.DB, name, email string) (int64, error) {
	res, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetUser(db *sql.DB, id int64) (User, error) {
	var u User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func DeleteUser(db *sql.DB, id int64) error {
	res, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("delete user %d: %w", id, sql.ErrNoRows)
	}
	return nil
}
