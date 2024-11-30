package middleware

import (
	"log"
	"net/http"
	"room-reservation/internal/repository"

	"github.com/justinas/nosurf"
)

var repo *repository.Repository

func SetRepo(r *repository.Repository) {
	repo = r
}

func UserPathHit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("Client requesting : ", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode, 
	})

	return csrfHandler
}

func LoadSessionManager(next http.Handler) http.Handler {
	return repo.App.SessionManager.LoadAndSave(next)
}