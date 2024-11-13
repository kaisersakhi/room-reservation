package routes

import (
	"net/http"
	"room-reservation/internal/config"
	"room-reservation/internal/handler"

	"github.com/go-chi/chi/v5"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handler.Home)

	return mux
}