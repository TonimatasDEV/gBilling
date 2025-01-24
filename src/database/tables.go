package database

import "log"

func CreateTables() {
	users()
	sessions()
	twoFactorAuth()
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
	CREATE TABLE IF NOT EXISTS ethene_user (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(320) NOT NULL UNIQUE,
        password VARCHAR(64) NOT NULL,
	    first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		phone_number VARCHAR(20) NOT NULL,
		verified BOOLEAN NOT NULL DEFAULT FALSE,
		lang VARCHAR(50) NOT NULL,
		announcements BOOLEAN NOT NULL DEFAULT FALSE,
		organization VARCHAR(255)
	);`

	/*
		country VARCHAR(100) NOT NULL,
		country_state VARCHAR(100) NOT NULL,
		city VARCHAR(100) NOT NULL,
		zipcode VARCHAR(20) NOT NULL,
		address VARCHAR(255) NOT NULL,
	*/

	createTable(query)
}

func sessions() {
	query := `
    CREATE TABLE IF NOT EXISTS ethene_session (
		id INT AUTO_INCREMENT PRIMARY KEY,
   		token TEXT NOT NULL,
        expires_at DATETIME NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   		id_user INT NOT NULL,
        CONSTRAINT session_id_user FOREIGN KEY (id_user) REFERENCES ethene_user(id) ON DELETE CASCADE
    );
    `

	createTable(query)
}

func twoFactorAuth() {
	query := `
    CREATE TABLE IF NOT EXISTS ethene_2fa (
		id INT AUTO_INCREMENT PRIMARY KEY,
    	secret VARCHAR(255) NOT NULL,
   		id_user INT NOT NULL,
        CONSTRAINT 2fa_id_user FOREIGN KEY (id_user) REFERENCES ethene_user(id) ON DELETE CASCADE
    );
    `

	createTable(query)
}
