package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var DATABASE *sql.DB

func Connect() *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalln("Error occurred connecting to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Error occurred connecting to the database:", err)
	}

	DATABASE = db
	log.Println("Connected to the database successfully.")
	return db
}
