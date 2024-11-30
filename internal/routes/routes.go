package routes

import (
	"net/http"
	"room-reservation/internal/config"
	"room-reservation/internal/handler"
	mw "room-reservation/internal/middleware"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(mw.UserPathHit)
	mux.Use(mw.NoSurf)
	mux.Use(mw.LoadSessionManager)

	mux.Get("/", handler.Home)
	mux.Get("/about", handler.About)

	return mux
}