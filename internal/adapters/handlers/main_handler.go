package handlers

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"net/http"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	domain.SendString(w, "Hello World!")
}
