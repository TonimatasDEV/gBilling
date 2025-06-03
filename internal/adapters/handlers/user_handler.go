package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TonimatasDEV/BillingPanel/internal/ports/services"
	"github.com/TonimatasDEV/BillingPanel/internal/util"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendString(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(req.Email, req.Password); err != nil {
		if util.IsMysqlError(err, 1062) {
			w.WriteHeader(http.StatusConflict)
			util.SendString(w, http.StatusConflict, "This email already exists.")
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	util.SendString(w, http.StatusCreated, "User created successfully.")
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	validate, err := h.service.Auth(getCookieValue(r))

	var resp struct {
		Message string `json:"msg"`
		Token   string `json:"token"`
	}

	if err == nil {
		resp.Message = "Already logged in."
		resp.Token = validate.Token
		util.SendJSON(w, http.StatusOK, resp)
		return
	}

	rawUser, err := h.service.GetRawUser(r)
	if err != nil {
		util.SendString(w, http.StatusBadRequest, "Invalid user input.")
		return
	}

	token, err := h.service.Login(rawUser)

	if err != nil {
		util.SendString(w, http.StatusUnauthorized, "Email or password is incorrect.")
		return
	}

	resp.Message = "User logged in successfully."
	resp.Token = token

	util.AddCookie(w, "session", token, time.Hour*24)
	util.SendJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	validate, err := h.service.Auth(getCookieValue(r))

	if err != nil {
		util.RemoveCookie(w, "session")
		util.SendString(w, http.StatusBadRequest, "You are not logged in.")
		return
	}

	err = h.service.Logout(validate.Token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	util.RemoveCookie(w, "session")
	util.SendString(w, http.StatusOK, "Logged out successfully.")
}

func getCookieValue(r *http.Request) string {
	cookies := r.CookiesNamed("session")

	if len(cookies) == 0 {
		return ""
	}

	return cookies[0].Value
}
