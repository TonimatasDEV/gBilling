package handlers

import (
	"net/http"

	"github.com/TonimatasDEV/BillingPanel/internal/util"
	"github.com/julienschmidt/httprouter"
)

func HandleMain(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	util.SendString(w, http.StatusOK, "Hello World!")
}
