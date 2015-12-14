package sapi

import (
	"github.com/kless/osutil/user/crypt/sha512_crypt"
)

type User struct {
	ID    string `json:"id"`
	Hash  string `json:"hash"`
	Email string `json:"email"`
}

func userQueryNoAuth(id string) (*User, error) {
	c := User{ID: id}
	err := db.Get(&c, "SELECT * FROM users WHERE id = $1", c.ID)
	return &c, err
}

func (c *User) HashMatch(password string) bool {
	cry := sha512_crypt.New()
	err := cry.Verify(c.Hash, []byte(password))
	if err == nil {
		return true
	}
	return false
}
