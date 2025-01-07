package src

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title   string
	Message string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:   "Billing Panel",
		Message: "¡Welcome to the home page of the Ethene Hosting billing panel!",
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al cargar la plantilla: %v", err), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "index.html", pageVariables)
}
