package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello, world!")

	if err != nil {
		fmt.Println("Error occur: ", err)
	}

	fmt.Printf("Number of bytes written: %d", n)
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "About page")

	if err != nil {
		fmt.Println("Error occur: ", err)
	}

	fmt.Printf("Number of bytes written: %d", n)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server port number: %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
