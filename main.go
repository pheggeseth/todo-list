package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

var todos = []Todo{
	{Id: 1, Title: "Learn to code a simple GO web server", Done: false},
	{Id: 2, Title: "Build a todo list app backend", Done: false},
	{Id: 3, Title: "Serve a UI for the app", Done: false },
}

var nextID = len(todos)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}

		nextID++

		todos = append(todos, Todo{
			Id: nextID,
			Title: title,
			Done: false,
		})

		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	todoId := r.URL.Path[len("/todos/"):]

	for i, todo := range todos {
		if strconv.Itoa(todo.Id) == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Todo with id deleted"))
		}
	}
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", deleteTodoHandler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
