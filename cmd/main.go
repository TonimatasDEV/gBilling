package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TonimatasDEV/BillingPanel/internal/adapters/handlers"
	"github.com/TonimatasDEV/BillingPanel/internal/adapters/persistence"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found or failed to load.")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	userRepo := persistence.NewMariaDBUserRepository(db)
	sessionRepo := persistence.NewMariaDBSessionRepository(db)

	sessionService := services.NewSessionService(sessionRepo)
	userService := services.NewUserService(userRepo, sessionService)

	userHandler := handlers.NewUserHandler(userService)

	router := httprouter.New()

	router.GET("/", handlers.HandleMain)
	router.POST("/users/create", userHandler.CreateHandler)
	router.POST("/users/login", userHandler.LoginHandler)
	router.POST("/users/logout", userHandler.LogoutHandler)

	log.Printf("Server running on http://localhost:%s\n", os.Getenv("PORT"))

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
