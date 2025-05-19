package main

import (
	"judo/configs"
	"judo/internal/base"
	"judo/internal/handlers/payload"
	"judo/internal/link"
	"judo/internal/stat"
	"judo/internal/user"
	"judo/migrations"
	"judo/pkg"
	"judo/pkg/db"
	"judo/pkg/event"
	"judo/pkg/jwt"
	"judo/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func App(dbInstance *db.DB, cfg *configs.Config) http.Handler {
	j := jwt.NewJWT(cfg.Auth.Secret)
	eventBus := event.NewEventBus()

	//repositories
	linkRepository := link.NewLinkRepository(dbInstance)
	userRepository := user.NewUserRepository(dbInstance)
	statRepository := stat.NewStatRepository(dbInstance)

	//auth service
	authService := user.NewAuthService(userRepository, j)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//migrations
	migrations.RunMigrations(dbInstance.DB)

	mux := http.NewServeMux()

	//jwt

	//handlers
	handlerAuth := payload.NewAuthHandler(cfg, authService)
	payload.AuthRouter(handlerAuth, mux)

	handlerLink := base.NewLinkHandler(link.LinkRepository{
		DataBase: linkRepository.DataBase,
	}, cfg, eventBus)
	payload.LinkRouter(handlerLink, mux)

	handlerStat := pkg.NewStatHandler(cfg, statRepository)
	payload.StatRouter(handlerStat, mux)

	stack := middleware.Chain(middleware.CORS,
		middleware.Logging,
	)

	go statService.AddClick()
	return stack(mux)
}

func main() {
	cfg := configs.MustLoad()
	db := db.NewDB(cfg)
	app := App(db, cfg)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        app,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
