package main

import (
	"log"
	"net/http"
	"room-reservation/internal/config"
	"room-reservation/internal/handler"
	"room-reservation/internal/middleware"
	"room-reservation/internal/render"
	"room-reservation/internal/repository"
	"room-reservation/internal/routes"
)

func main() {
	app := config.NewAppConfig()
	app.SessionManager = newSessionManager()

	repo := repository.NewRepository(app)
	middleware.SetRepo(repo)
	handler.SetRepo(repo)

	SetPortAndEnv(app)
	SetTemplateCache(app)
	render.NewRenderer(app)

	server := &http.Server{
		Addr:    ":" + app.GetPort(),
		Handler: routes.Routes(app),
	}

	log.Fatal("Failed to start the server : ", server.ListenAndServe())
}
