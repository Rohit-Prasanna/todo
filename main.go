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

	// Configure CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500", "https://rohit-prasanna.github.io/todo-list-withgo_backend/"}, // Replace with your frontend's URLs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	// Apply CORS middleware to the router
	corsHandler := cors.New(corsOptions).Handler(router)

	// Get the port from environment variables (Render provides this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no port is specified
	}

	// Start the server with the CORS handler on the correct port
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
