package util

import (
	"encoding/json"
	"net/http"
)

type Msg struct {
	Msg string `json:"msg"`
}

func SendError(w http.ResponseWriter, err error) {
	create(err.Error()).send(w)
}

func SendString(w http.ResponseWriter, status int, str string) {
	w.WriteHeader(status)
	create(str).send(w)
}

func create(str string) *Msg {
	return &Msg{Msg: str}
}

func (msg *Msg) send(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(msg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SendJSON(w http.ResponseWriter, status int, payload any) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
