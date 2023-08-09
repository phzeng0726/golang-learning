package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 使用 Go 的標準庫，建立一個簡單的 JSON API，能夠存取和操作一個待辦事項清單。

type Task struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

var tasks = []Task{
	{Id: "1", Title: "1 task", IsDone: false},
	{Id: "2", Title: "2 task", IsDone: false},
	{Id: "3", Title: "3 task", IsDone: false},
	{Id: "4", Title: "4 task", IsDone: false},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTaskIndexById(id string) (int, error) {
	for i, t := range tasks {
		if t.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Task ID not found")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}
	// 找出id的i，並更新
	i, err := getTaskIndexById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := tasks[i]
	task.IsDone = !task.IsDone
	tasks[i] = task

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}
	// 找出id的i，並刪除
	i, err := getTaskIndexById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks = append(tasks[:i], tasks[i+1:]...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	case http.MethodPatch:
		updateTask(w, r)
	case http.MethodDelete:
		deleteTask(w, r)
	}
}

func main() {
	http.HandleFunc("/tasks", taskHandler)

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
