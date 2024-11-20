package routes

import (
	"github.com/Rohit-Prasanna/todo/handlers"

	"github.com/gorilla/mux"
)

// InitRoutes initializes all API routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todos/{userId}", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	return router
}
