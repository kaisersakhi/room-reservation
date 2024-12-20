package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"room-reservation/internal/config"
	"room-reservation/internal/render"
	"time"

	"github.com/alexedwards/scs/v2"
)

func SetPortAndEnv(app *config.AppConfig) {
	port := flag.String("port", "4000", "port to run the application on")
	env := flag.String("env", "development", "application environment (development, production)")

	flag.Parse()

	portErr := app.SetPort(*port)
	envErr := app.SetEnv(*env)

	if portErr != nil || envErr != nil { 
		log.Fatal(portErr, envErr)
	}

	fmt.Printf("Running on port: %s\n", *port)
	fmt.Printf("Environment: %s\n", *env)
}

func SetTemplateCache(app *config.AppConfig) {
	templateCache, err := render.BuildTemplateCache()

	if err != nil {
		log.Fatal("error while buding the template cache", err)
	}

	app.SetTemplateCache(templateCache)
}


func newSessionManager() *scs.SessionManager {
	session := scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	return session
}