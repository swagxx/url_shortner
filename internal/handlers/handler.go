package handlers

import "net/http"

func InitRouters(auth *AuthHandler, router *http.ServeMux) {
	router.HandleFunc("/auth/register", auth.Register())
	router.HandleFunc("/auth/login", auth.Login())
}
