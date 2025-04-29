package handlers

import (
	"io"
	"judo/configs"
	"judo/internal/handlers/payload"
	"judo/pkg/handlerset"
	"judo/pkg/request"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(cfg *configs.Config) *AuthHandler {
	return &AuthHandler{
		Config: cfg,
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !checkMethodHead(w, r) {
			return
		}
		defer r.Body.Close()
		reg, err := request.HandleBody[payload.RegisterRequest](w, r)
		if err != nil {
			return
		}

		handlerset.HandlerSet(w, struct {
			Email string `json:"email"`
		}{
			Email: reg.Email,
		}, http.StatusCreated)
		io.ReadAll(r.Body)
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !checkMethodHead(w, r) {
			return
		}
		defer r.Body.Close()

		log, err := request.HandleBody[payload.LoginRequest](w, r)
		if err != nil {
			return
		}

		handlerset.HandlerSet(w, struct {
			Token string `json:"token"`
		}{
			log.Password,
		}, http.StatusOK)
	}
}

func checkMethodHead(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return false
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid header Content-Type", http.StatusUnsupportedMediaType)
		return false
	}
	return true
}
