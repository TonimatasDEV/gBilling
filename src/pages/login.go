package src

import (
	"net/http"
)

type Error struct {
	Text string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	/* TODO: Only send login data.
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		ok, err := auth.CheckPassword(email, password)

		if err != nil { // TODO: Send the error to the login page.
			http.Error(w, "Not account found with "+email+".", http.StatusBadRequest)
			return
		}

		if ok {
			time.Sleep(25 * time.Millisecond)

			token, err := auth.GenerateToken(email)
			if err == nil {
				http.SetCookie(w, &http.Cookie{
					Name:     "ethene_session",
					Value:    token,
					HttpOnly: true,
					Secure:   true,
					Expires:  time.Now().Add(time.Hour * 24 * 30),
					SameSite: http.SameSiteStrictMode,
				})
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized) // TODO: Send the error to the login page.
		return
	}

	if auth.CheckSession(w, r) {
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	*/
}
