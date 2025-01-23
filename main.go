package main

import (
	"context"
	"errors"
	auth2 "github.com/TonimatasDEV/BillingPanel/src/auth"
	"github.com/TonimatasDEV/BillingPanel/src/database"
	"github.com/TonimatasDEV/BillingPanel/src/pages"
	"github.com/TonimatasDEV/BillingPanel/src/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	utils.LoadEnvFile()
	db := database.Connect()
	database.CreateTables()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", auth(src.IndexHandler))
	http.HandleFunc("/logout", auth(src.LogoutHandler))
	http.HandleFunc("/login", src.LoginHandler)

	server := &http.Server{
		Addr: ":8080",
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("Server crashed:", err)
		}
	}()

	log.Println("Server listening on :8080.")

	<-stop
	log.Println("Server shutting down.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error occurred closing the database:", err)
		}

		log.Println("Error stopping the server:", err)
	}

	log.Println("Server stopped successfully.")
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth2.CheckSession(w, r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
