package persistence

import (
	"database/sql"
	"errors"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type MariaDBSessionRepository struct {
	db *sql.DB
}

func NewMariaDBSessionRepository(db *sql.DB) repositories.SessionRepository {
	query := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		user_id INTEGER NOT NULL,
	    token VARCHAR(255) NULL UNIQUE,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		exp DATETIME NOT NULL,
	    FOREIGN KEY (user_id) REFERENCES users(id)
	        ON DELETE CASCADE
	        ON UPDATE CASCADE
	) ENGINE=InnoDB`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating the sessions table: %v", err)
	}

	return &MariaDBSessionRepository{db: db}
}

func (r *MariaDBSessionRepository) Create(userID int) (*domain.Session, error) {
	exp := time.Now().UTC().Add(time.Hour * 24)

	res, err := r.db.Exec("INSERT INTO sessions (user_id, exp) VALUES (?, ?)", userID, exp)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	token, err := generateToken(id, exp)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec("UPDATE sessions SET token = ? WHERE id = ?", token, id)
	if err != nil {
		return nil, err
	}

	session, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *MariaDBSessionRepository) Remove(token string) error {
	_, err := r.db.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

func generateToken(sessionID int64, exp time.Time) (string, error) {
	claims := jwt.MapClaims{
		"session_id": sessionID,
		"exp":        exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (r *MariaDBSessionRepository) Validate(tokenStr string) (*domain.Session, error) {
	token, err := getToken(tokenStr)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionID := claims["session_id"].(float64)

		session, err := r.GetByID(int64(sessionID))
		if err != nil {
			return nil, err
		}

		rawExp := claims["exp"].(float64)
		exp := time.Unix(int64(rawExp), 0)

		if !time.Now().After(exp) && session.Token == tokenStr {
			return session, nil
		}
	}

	return nil, errors.New("invalid token")
}

func getToken(tokenString string) (*jwt.Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secret, nil
	})
}

func (r *MariaDBSessionRepository) GetByID(id int64) (*domain.Session, error) {
	var createdAt string
	session := &domain.Session{}
	err := r.db.QueryRow("SELECT * FROM sessions WHERE id = ?", id).Scan(&session.ID, &session.UserID, &session.Token, &createdAt, &session.Exp)

	if err != nil {
		return nil, err
	}

	createdAtTime, err := time.Parse(time.DateTime, createdAt)
	if err != nil {
		return nil, err
	}

	session.CreatedAt = createdAtTime

	return session, err
}
