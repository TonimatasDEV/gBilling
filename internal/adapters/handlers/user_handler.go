package handlers

import (
	"encoding/json"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/services"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	user := h.service.CreateUser(req.Email, req.Password)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
