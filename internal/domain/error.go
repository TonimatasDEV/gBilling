package domain

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Msg string `json:"msg"`
}

func SendError(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(&Error{Msg: err.Error()})
}
