package handlers

import (
	"encoding/json"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
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
	_ = json.NewDecoder(r.Body).Decode(&req)

	err := h.service.CreateUser(req.Email, req.Password)

	if err != nil {
		domain.SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	domain.SendString(w, "User created successfully.")
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var resp struct {
		Token string `json:"token"`
	}

	resp.Token = token

	_ = json.NewEncoder(w).Encode(resp)
}
