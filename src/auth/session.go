package auth

import (
	"crypto/rand"
	"github.com/TonimatasDEV/BillingPanel/src/database"
	"math/big"
	"time"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$&*()_[]")

func GenerateToken(email string) (string, error) {
	expiration := time.Now().Add(time.Hour * 24 * 30)

	token := make([]rune, 128)

	for i := range token {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		token[i] = charset[idx.Int64()]
	}

	err := saveToken(email, string(token), expiration)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func saveToken(email string, tokenString string, expiresAt time.Time) error {
	query := "INSERT INTO sessions (email, token, expires_at) VALUES (?, ?, ?)"
	_, err := database.DATABASE.Exec(query, email, tokenString, expiresAt)
	return err
}

func validateToken(token string) bool {
	query := "SELECT token FROM sessions WHERE token = ?"
	_, err := database.DATABASE.Exec(query, token)

	if err != nil || checkTokenExpiration(token) {
		return false
	}

	return true
}

func checkTokenExpiration(token string) bool {
	var expiresAt string
	var storedToken string

	query := "SELECT token, expires_at FROM sessions WHERE token = ?"
	err := database.DATABASE.QueryRow(query, token).Scan(&storedToken, &expiresAt)
	if err != nil {
		println(err.Error())
		return true
	}

	parse, err := time.Parse(time.DateTime, expiresAt)
	if err != nil {
		return false
	}

	if parse.Before(time.Now()) {
		removeExpiredToken(token)
		return true
	}

	return false
}

func removeExpiredToken(token string) {
	query := "DELETE FROM sessions WHERE token = ?"
	_, _ = database.DATABASE.Exec(query, token)
}
