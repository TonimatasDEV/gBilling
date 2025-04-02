package persistence

import (
	"database/sql"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/repositories"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MariaDBUserRepository struct {
	db *sql.DB
}

func NewMariaDBUserRepository(dsn string) repositories.UserRepository {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(100) NOT NULL UNIQUE,
		hashed_password VARCHAR(255) NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error creando la tabla: %v", err)
	}

	return &MariaDBUserRepository{db: db}
}

func (r *MariaDBUserRepository) Save(user domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (email, hashed_password) VALUES (?, ?)", user.Email, user.HashedPassword)
	return err
}

func (r *MariaDBUserRepository) GetByEmail(email string) (domain.User, error) {
	var id int
	var getEmail string
	var hashedPassword string

	err := r.db.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&id, &getEmail, &hashedPassword)
	user := domain.User{ID: id, Email: getEmail, HashedPassword: hashedPassword}

	return user, err
}
