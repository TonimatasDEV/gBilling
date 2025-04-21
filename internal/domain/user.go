package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int    `json:"id"`
	RoleID         int    `json:"roleId"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err
}
