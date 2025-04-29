package main

import (
	"judo/configs"
	"judo/internal/handlers"
	"judo/pkg/db"
	"net/http"
	"time"
)

func init() {

}

func main() {
	cfg := configs.MustLoad()
	_ = db.NewDB(cfg)
	mux := http.NewServeMux()
	handler := handlers.NewAuthHandler(cfg)
	handlers.InitRouters(handler, mux)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
