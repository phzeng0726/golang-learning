package questions

import (
	"fmt"
	"log"
	"net/http"
)

// 通用元件
func InitRoute() {
	http.HandleFunc("/", HomeHandler)
	// http.HandleFunc("/notes", GetNotes).Methods("GET")
	// http.HandleFunc("/notes", CreateNote).Methods("POST")
	// http.HandleFunc("/notes/{id}", UpdateNote).Methods("PUT")
	// http.HandleFunc("/notes/{id}", DeleteNote).Methods("DELETE")
}

func HttpServer() {
	InitRoute()

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
