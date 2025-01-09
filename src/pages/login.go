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

		if !auth.EmailExist(email) {
			http.Error(w, "Not account found with "+email+".", http.StatusBadRequest)
		}

		if auth.CheckPassword(email, password) {
			time.Sleep(25 * time.Millisecond)
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token", // TODO: Generate secure tokens.
				Value:   "some_secure_session_token",
				Expires: time.Now().Add(1 * time.Hour),
			})

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		/* TODO: Send a message to the login page.
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
