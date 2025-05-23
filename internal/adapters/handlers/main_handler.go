package handlers

import (
	"github.com/TonimatasDEV/BillingPanel/internal/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleMain(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	util.SendString(w, "Hello World!")
}
