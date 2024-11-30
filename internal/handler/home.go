package handler

import (
	"net/http"
	"room-reservation/internal/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.html")
}
