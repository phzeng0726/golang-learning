package main

import (
	"fmt"
	"learning/pkg/config"
	"learning/pkg/handlers"
	"learning/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	fmt.Println(app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server port number: %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
