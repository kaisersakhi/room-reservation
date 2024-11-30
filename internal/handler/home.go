package handler

import (
	"log"
	"net/http"
	"room-reservation/internal/render"
	"room-reservation/internal/repository"
)

var repo *repository.Repository

func SetRepo(r *repository.Repository) {
	repo = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	repo.App.SessionManager.Put(r.Context(), "remote_ip", r.RemoteAddr)
	log.Println("Value store inside session is: ", repo.App.SessionManager.GetString(r.Context(), "remote_ip"))
	render.Render(w, "home.html")
}
