package src

import (
	"github.com/TonimatasDEV/EtheneBillingPanel/src/utils"
	"net/http"
	"time"
)

var validUser = map[string]string{
	"admin": "admin", // username: password
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if validPassword(username, password) {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   "some_secure_session_token", // In a real application, use a secure token
				Expires: time.Now().Add(1 * time.Hour),
			})

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
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

func validPassword(username, password string) bool {
	expectedPassword, ok := validUser[username]
	return ok && expectedPassword == password
}
