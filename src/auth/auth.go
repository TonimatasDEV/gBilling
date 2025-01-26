package auth

import "net/http"

func Check(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if checkSession(w, r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
