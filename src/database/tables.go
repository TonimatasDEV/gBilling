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
		password VARCHAR(255) NOT NULL
	    
	    first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		phone_number VARCHAR(20) NOT NULL,
		verified BOOLEAN NOT NULL DEFAULT FALSE,
		country VARCHAR(100) NOT NULL,
		country_state VARCHAR(100) NOT NULL,
		city VARCHAR(100) NOT NULL,
		postal_code VARCHAR(20) NOT NULL,
		address VARCHAR(255) NOT NULL,

		lang VARCHAR(20) NOT NULL,
		announcements BOOLEAN NOT NULL DEFAULT FALSE,

		organization VARCHAR(255),
		two_factor_auth BOOLEAN NOT NULL DEFAULT FALSE
	);`

	createTable(query)
}
