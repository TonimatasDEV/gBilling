package main

import (
	"context"
	"errors"
	"github.com/TonimatasDEV/BillingPanel/src/pages"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", auth(src.IndexHandler))
	http.HandleFunc("/login", src.LoginHandler)

	server := &http.Server{
		Addr: ":8080",
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server crashed: %v\n", err)
		}
	}()

	log.Println("Server listening on :8080.")

	<-stop
	log.Println("Server shutting down.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error stopping the server: %v\n", err)
	}

	log.Println("Server stopped successfully.")
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value != "some_secure_session_token" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
