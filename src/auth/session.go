package auth

import (
	"errors"
	"github.com/TonimatasDEV/BillingPanel/src/database"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		println(err.Error())
		return "", err
	}

	err = saveToken(email, tokenString, expirationTime)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func saveToken(email string, tokenString string, expiresAt time.Time) error {
	query := "INSERT INTO sessions (email, token, expires_at) VALUES (?, ?, ?)"
	_, err := database.DATABASE.Exec(query, email, tokenString, expiresAt)
	return err
}

func validateToken(tokenString string) (bool, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	if checkTokenExpiration(tokenString) {
		return false, nil
	}

	return checkEmailMatch(claims.Email, tokenString), nil
}

func checkTokenExpiration(token string) bool {
	var expiresAt string
	var storedToken string

	query := "SELECT token, expires_at FROM sessions WHERE token = ?"
	err := database.DATABASE.QueryRow(query, token).Scan(&storedToken, &expiresAt)
	if err != nil {
		return true
	}

	parse, err := time.Parse(time.DateTime, expiresAt)
	if err != nil {
		return false
	}

	if parse.Before(time.Now()) {
		RemoveToken(token)
		return true
	}

	return false
}

func checkEmailMatch(email string, token string) bool {
	var storedEmail string
	query := "SELECT email FROM sessions WHERE token = ?"
	err := database.DATABASE.QueryRow(query, token).Scan(&storedEmail)
	if err != nil {
		return false
	}

	return email == storedEmail
}

func RemoveToken(token string) {
	query := "DELETE FROM sessions WHERE token = ?"
	_, _ = database.DATABASE.Exec(query, token)
}
