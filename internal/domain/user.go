package domain

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	ID             int    `json:"id"`
	RoleID         int    `json:"roleId"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
}

type RawUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err
}

func CreateRawUser(r *http.Request) (RawUser, error) {
	var rawUser RawUser
	err := json.NewDecoder(r.Body).Decode(&rawUser)
	return rawUser, err
}
