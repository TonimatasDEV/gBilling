package src

import (
	"github.com/TonimatasDEV/BillingPanel/src/auth"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("ethene_session")

	if err == nil {
		auth.RemoveToken(cookie.Value)
		auth.DeleteCookie(w)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
