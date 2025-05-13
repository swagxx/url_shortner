package handlers

import (
	"judo/pkg/middleware"
	"net/http"
)

func AuthRouter(auth *AuthHandler, router *http.ServeMux) {
	router.HandleFunc("/auth/register", auth.Register())
	router.HandleFunc("/auth/login", auth.Login())
}

func LinkRouter(link *LinkHandler, router *http.ServeMux) {
	router.HandleFunc("POST /link", link.Create())
	router.HandleFunc("GET /{hash}", link.Read())
	router.Handle("PATCH /link/{id}", middleware.Bearer(link.Update()))
	router.HandleFunc("DELETE /link/{id}", link.Delete())
	router.HandleFunc("GET /link-all", link.GetAll())
}
