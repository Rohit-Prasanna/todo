package routes

import (
	"github.com/Rohit-Prasanna/todo/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Fetch all todos
	router.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")

	// Create a new todo
	router.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")

	// Update a todo by ID
	router.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")

	// Delete a todo by ID
	router.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	return router
}
