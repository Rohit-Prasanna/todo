package main

import (
	"github.com/Rohit-Prasanna/todo/db"
	"github.com/Rohit-Prasanna/todo/routes"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	// Connect to MongoDB
	if err := db.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.DisconnectDB()

	// Initialize routes
	router := routes.InitRoutes()

	// Apply CORS middleware to the router
	corsHandler := cors.Default().Handler(router)

	// Get the port from environment variable (Render provides this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not provided
	}

	// Start the server with the CORS handler on the correct port
	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
