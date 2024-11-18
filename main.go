package main

import (
	"github.com/Rohit-Prasanna/todo/db"
	"github.com/Rohit-Prasanna/todo/routes"
	"log"
	"net/http"
)

func main() {
	// Connect to MongoDB
	if err := db.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.DisconnectDB()

	// Initialize routes
	router := routes.InitRoutes()

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
