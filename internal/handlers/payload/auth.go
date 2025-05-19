package payload

import (
	"judo/configs"
	"judo/internal/user"
	"judo/pkg/dto"
	"judo/pkg/handlerset"
	"judo/pkg/request"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*user.AuthService
}

func NewAuthHandler(cfg *configs.Config, user *user.AuthService) *AuthHandler {
	return &AuthHandler{
		Config:      cfg,
		AuthService: user,
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !checkMethodHead(w, r) {
			return
		}
		defer r.Body.Close()
		reg, err := request.HandleBody[dto.RegisterRequest](w, r)
		if err != nil {
			return
		}

		token, err := h.AuthService.Register(reg.Email, reg.Password, reg.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, struct {
			Token string `json:"token"`
		}{Token: token}, http.StatusOK)

	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !checkMethodHead(w, r) {
			return
		}
		defer r.Body.Close()

		body, err := request.HandleBody[dto.LoginRequest](w, r)
		if err != nil {
			return
		}

		res, err := h.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, "Login failed", http.StatusUnauthorized)
			return
		}

		handlerset.HandlerSet(w, struct {
			Token string `json:"token"`
		}{
			Token: res,
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
