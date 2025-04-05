package main

import (
	"fmt"
	"github.com/TonimatasDEV/BillingPanel/internal/adapters/handlers"
	"github.com/TonimatasDEV/BillingPanel/internal/adapters/persistence"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/services"
	"net/http"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/ethene" // TODO
	repo := persistence.NewMariaDBUserRepository(dsn)

	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleMain)
	mux.HandleFunc("/users/create", handler.CreateUserHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
