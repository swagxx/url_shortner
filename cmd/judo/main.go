package main

import (
	"judo/configs"
	"judo/internal/handlers"
	"judo/internal/link"
	"judo/internal/user"
	"judo/migrations"
	"judo/pkg/db"
	"judo/pkg/middleware"
	"net/http"
	"time"
)

func init() {

}

func main() {
	cfg := configs.MustLoad()

	dbInstance := db.NewDB(cfg)

	//repositories
	linkRepository := link.NewLinkRepository(dbInstance)
	userRepository := user.NewUserRepository(dbInstance)
	//auth service
	authService := user.NewAuthService(userRepository)

	//migrations
	migrations.RunMigrations(dbInstance.DB)

	mux := http.NewServeMux()

	//handlers
	handlerAuth := handlers.NewAuthHandler(cfg, authService)
	handlers.AuthRouter(handlerAuth, mux)

	handlerLink := handlers.NewLinkHandler(link.LinkRepository{
		DataBase: linkRepository.DataBase,
	})
	handlers.LinkRouter(handlerLink, mux)

	stack := middleware.Chain(middleware.CORS,
		middleware.Logging,
	)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        stack(mux),
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
