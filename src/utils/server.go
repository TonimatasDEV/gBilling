package utils

import (
	"github.com/TonimatasDEV/BillingPanel/src/auth"
	"net"
	"net/http"
)

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request), authEnabled bool) {
	if authEnabled {
		http.HandleFunc(pattern, checkProxy(checkAuth(handler)))
	} else {
		http.HandleFunc(pattern, checkProxy(handler))
	}
}

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.CheckSession(w, r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}

func checkProxy(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		if IsCloudflareIP(ip) || !proxyEnabled {
			next(w, r)
		} else {
			http.Error(w, "Sorry, you are not authorized to access this page.", http.StatusForbidden)
		}
	}
}
