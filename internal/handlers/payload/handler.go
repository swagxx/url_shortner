package payload

import (
	"judo/internal/base"
	"judo/pkg"
	"judo/pkg/middleware"
	"net/http"
)

func AuthRouter(auth *AuthHandler, router *http.ServeMux) {
	router.HandleFunc("/auth/register", auth.Register())
	router.HandleFunc("/auth/login", auth.Login())
}

func LinkRouter(link *base.LinkHandler, router *http.ServeMux) {
	router.Handle("POST /link", middleware.Bearer(link.Create(), link.Config))
	router.HandleFunc("GET /{hash}", link.Read())
	router.Handle("PATCH /link/{id}", middleware.Bearer(link.Update(), link.Config))
	router.Handle("DELETE /link/{id}", middleware.Bearer(link.Delete(), link.Config))
	router.Handle("GET /link-all/", middleware.Bearer(link.GetAll(), link.Config))
}

func StatRouter(stat *pkg.StatHandler, router *http.ServeMux) {
	router.HandleFunc("GET /stat/", stat.StatByDate())
}
