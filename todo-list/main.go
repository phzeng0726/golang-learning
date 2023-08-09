package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// 1. 使用 Go 的標準庫，建立一個簡單的 JSON API，能夠存取和操作一個待辦事項清單。
// 2. 創建一個簡單的 To-Do 應用，使用檔案儲存任務清單，能夠新增、刪除和列印任務。
type Task struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

var tasks []Task
var taskFilePath string = "./data/tasks.json"

func readTaskJson() error {
	filePtr, err := os.Open(taskFilePath)
	if err != nil {
		fmt.Println("Cannot open tasks.json")
		return err
	}
	defer filePtr.Close()

	err = json.NewDecoder(filePtr).Decode(&tasks)

	if err != nil {
		fmt.Println("Decode json failed", err.Error())
	} else {
		fmt.Println("Decode json succeed")
		fmt.Println(tasks)
	}

	return err
}

func editTaskJson() error {
	filePtr, err := os.Create(taskFilePath)
	if err != nil {
		fmt.Println("Cannot open tasks.json")
		return err
	}
	defer filePtr.Close()

	err = json.NewEncoder(filePtr).Encode(&tasks)

	if err != nil {
		fmt.Println("Encode json failed", err.Error())
	} else {
		fmt.Println("Encode json succeed")
		fmt.Println(tasks)
	}

	return err
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
	editTaskJson()
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
	editTaskJson()
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
	editTaskJson()
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
		return
	case http.MethodPost:
		createTask(w, r)
		return
	case http.MethodPatch:
		updateTask(w, r)
		return
	case http.MethodDelete:
		deleteTask(w, r)
		return
	}
}

func main() {
	readTaskJson()

	http.HandleFunc("/tasks", taskHandler)

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
