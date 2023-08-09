package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 使用 Go 的 net/http 庫，建立一個簡單的 RESTful API，能夠處理用戶的 CRUD（新增、讀取、更新和刪除）請求。

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: "1", Name: "Rita", Age: 26},
	{ID: "2", Name: "Yong", Age: 27},
	{ID: "3", Name: "Mona", Age: 23},
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	var updated bool

	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if updateUser.ID == u.ID {
			users[i] = updateUser
			updated = true
			break
		}
	}

	if !updated {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	i, err := getUserIndexById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	users = append(users[:i], users[i+1:]...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserIndexById(id string) (int, error) {
	for i, u := range users {
		if u.ID == id {
			return i, nil
		}
	}

	return -1, errors.New("User not found")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

func main() {
	http.HandleFunc("/users", userHandler)

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
