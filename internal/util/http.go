package util

import (
	"net/http"
	"time"
)

func AddCookie(w http.ResponseWriter, name, value string, duration time.Duration) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(duration.Seconds()),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
}

func RemoveCookie(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}
