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

func SendString(w http.ResponseWriter, str string) {
	create(str).send(w)
}

func create(str string) *Msg {
	return &Msg{Msg: str}
}

func (msg *Msg) send(w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(msg)
}
