package main

import (
	"fmt"
	"learning/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server port number: %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
