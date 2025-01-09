package src

import (
	"github.com/TonimatasDEV/BillingPanel/src/auth"
	"github.com/TonimatasDEV/BillingPanel/src/utils"
	"net/http"
	"time"
)

type Error struct {
	Text string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if validPassword(email, password) {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   "some_secure_session_token", // In a real application, use a secure token
				Expires: time.Now().Add(1 * time.Hour),
			})

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		/*
			err := Error{
				Text: "Invalid credentials.",
			}

			utils.SendTemplate(w, "login.html", err, "templates/login.html")
		*/
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	cookie, err := r.Cookie("session_token")

	if err != nil || cookie.Value != "some_secure_session_token" {
		utils.SendTemplate(w, "login.html", nil, "templates/login.html")
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func validPassword(email, password string) bool {
	expectedPassword, ok := auth.Users[email]
	return ok && expectedPassword == password
}
