package handlers

import (
	"encoding/json"
	"github.com/TonimatasDEV/BillingPanel/internal/domain"
	"github.com/TonimatasDEV/BillingPanel/internal/ports/services"
	"github.com/TonimatasDEV/BillingPanel/internal/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.service.CreateUser(req.Email, req.Password)

	if err != nil {
		if util.IsMysqlError(err, 1062) {
			domain.SendString(w, "This email already exists.")
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	domain.SendString(w, "User created successfully.")
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		println(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var resp struct {
		Token string `json:"token"`
	}

	resp.Token = token

	util.AddCookie(w, "session", token, time.Hour*24)
	_ = json.NewEncoder(w).Encode(resp)
}
