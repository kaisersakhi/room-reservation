package handler

import (
	"log"
	"net/http"
	"room-reservation/internal/render"
)

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("about about hello")
	remote_ip := repo.App.SessionManager.GetString(r.Context(), "remote_ip")
	log.Println("User ip is : ", remote_ip)

	render.Render(w, "about.html")
}