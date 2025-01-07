package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func SendTemplate(w http.ResponseWriter, templateName string, variables any, filenames ...string) {
	t, err := template.ParseFiles(filenames...)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing the template: %v", err), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, templateName, variables)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing the template: %v", err), http.StatusInternalServerError)
	}
}
