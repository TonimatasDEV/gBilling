package persistence

import (
	"database/sql"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
	"github.com/golang-jwt/jwt/v5"
	"log"
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

func (r *MariaDBSessionRepository) Create(userId int) (*domain.Session, error) {
	exp := time.Now().UTC().Add(time.Hour * 24)

	res, err := r.db.Exec("INSERT INTO sessions (user_id, exp) VALUES (?, ?)", userId, exp)

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

	session, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func generateToken(sessionId int64, exp time.Time) (string, error) {
	claims := jwt.MapClaims{
		"session_id": sessionId,
		"exp":        exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("8c5b66db3219072f0809036a3fc0f0a4b375f1be7c100014ac06f3d0bac15de7")) // TODO: Use env JWT_SECRET
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (r *MariaDBSessionRepository) GetById(id int64) (*domain.Session, error) {
	var createdAt string
	session := &domain.Session{}
	err := r.db.QueryRow("SELECT * FROM sessions WHERE id = ?", id).Scan(&session.Id, &session.UserId, &session.Token, &createdAt, &session.Exp)

	createdAtTime, err := time.Parse(time.DateTime, createdAt)
	if err != nil {
		return nil, err
	}

	session.CreatedAt = createdAtTime

	return session, err
}
