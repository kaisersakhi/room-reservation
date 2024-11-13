package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"room-reservation/internal/config"
	"room-reservation/internal/render"
	"room-reservation/internal/routes"
)

func main() {
	app := config.NewAppConfig()
	SetPortAndEnv(app)
	SetTemplateCache(app)

	server := &http.Server{
		Addr:    ":" + app.GetPort(),
		Handler: routes.Routes(app),
	}

	log.Fatal("Failed to start the server : ", server.ListenAndServe())
}

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
