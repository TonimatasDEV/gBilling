package src

import (
	"github.com/TonimatasDEV/EtheneBillingPanel/src/utils"
	"net/http"
)

type PageVariables struct {
	Title   string
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:   "Billing Panel",
		Message: "¡Welcome to the home page of the Ethene Hosting billing panel!",
	}

	utils.SendTemplate(w, "index.html", pageVariables, "templates/index.html")
}
