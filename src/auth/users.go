package auth

import (
	"database/sql"
	"errors"
	"github.com/TonimatasDEV/BillingPanel/src/database"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func InsertUser(email, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	_, err = database.DATABASE.Exec(query, email, hashedPassword)
	return err
}

func CheckPassword(email, password string) (bool, error) {
	var passwordHash string
	query := "SELECT password FROM users WHERE email = ?"
	err := database.DATABASE.QueryRow(query, email).Scan(&passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("user not found")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil, nil
}
