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

func RegisterUser(email, password, fistName, lastName, phoneNumber, country, countyState, city, zipcode, address, lang, organization string, announcements bool) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (email, password, first_name, last_name, phone_number, country, country_state, city, zipcode, address, lang, announcements, organization) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = database.DATABASE.Exec(query, email, hashedPassword, fistName, lastName, phoneNumber, country, countyState, city, zipcode, address, lang, announcements, organization)
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
