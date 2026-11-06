package piscine

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"-"`
}

func MarshalUser(u User) ([]byte, error) {
	return json.Marshal(u)
}

func UnmarshalUser(data []byte) (User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	return u, err
}

func PrettyPrint(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
