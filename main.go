package main

import (
	"github.com/TonimatasDEV/EtheneBillingPanel/src/pages"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", auth(src.IndexHandler))
	http.HandleFunc("/login", src.LoginHandler)

	log.Println("Starting the server with the port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Server crashed: ", err)
		return
	}
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
