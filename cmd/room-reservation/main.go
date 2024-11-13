package main

import (
	"flag"
	"fmt"
)

func main(){
	port := flag.String("port", "8080", "port to run the application on")
	env := flag.String("env", "production", "application environment (development, production)")

	flag.Parse()

	fmt.Printf("Running on port: %s\n", *port)
	fmt.Printf("Environment: %s\n", *env)
	fmt.Println("d")
}
