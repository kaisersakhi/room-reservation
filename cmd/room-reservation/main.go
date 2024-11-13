package main

import (
	"flag"
	"fmt"
	"log"
	"room-reservation/internal/config"
)


func main(){
	app := config.NewAppConfig()
	SetPortAndEnv(app)
}

func SetPortAndEnv(app *config.AppConfig){
	port := flag.String("port", "8080", "port to run the application on")
	env := flag.String("env", "production", "application environment (development, production)")

	flag.Parse()

	fmt.Printf("Running on port: %s\n", *port)
	fmt.Printf("Environment: %s\n", *env)

	portErr := app.SetPort(*port)
	envErr := app.SetEnv(*env)

	if portErr != nil || envErr != nil {
		log.Fatal(portErr, envErr)
	}
}
