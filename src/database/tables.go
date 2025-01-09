package database

import "log"

func CreateTables() {
	users()
	log.Println("Database tables successfully checked.")
}

func createTable(query string) {
	_, err := DATABASE.Exec(query)

	if err != nil {
		log.Fatalln("Error creating database table:", err)
	}
}

func users() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL
	);`

	createTable(query)
}
