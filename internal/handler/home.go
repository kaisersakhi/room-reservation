package handler

import (
	"log"
	"net/http"
	"room-reservation/internal/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home route hit")
	render.Render(w, "home.html")
}
