package handlers

import (
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleMain(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	domain.SendString(w, "Hello World!")
}
