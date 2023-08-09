package questions

import (
	"fmt"
	"log"
	"net/http"
)

// 建立一個基本的 HTTP 伺服器，該伺服器能夠處理 GET 請求並回傳一個簡單的 "Hello, World!" 響應。
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func HttpServer() {
	http.HandleFunc("/", HomeHandler)

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
