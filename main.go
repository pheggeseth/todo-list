package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func main() {
	http.HandleFunc("/", todosHandler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
