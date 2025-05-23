package main

import (
	"database/sql"
	"fmt"
	"github.com/TonimatasDEV/BillingPanel/internal/adapters/handlers"
	"github.com/TonimatasDEV/BillingPanel/internal/adapters/persistence"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/services"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/ethene" // TODO
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	userRepo := persistence.NewMariaDBUserRepository(db)
	sessionRepo := persistence.NewMariaDBSessionRepository(db)

	service := services.NewUserService(userRepo, sessionRepo)
	handler := handlers.NewUserHandler(service)

	router := httprouter.New()

	router.GET("/", handlers.HandleMain)
	router.POST("/users/create", handler.CreateUserHandler)
	router.POST("/users/login", handler.LoginUserHandler)

	fmt.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
