package main

import (
	"github.com/Rohit-Prasanna/todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(r)

	// Start the server
	err := r.Run(":8080")
	if err != nil {
		return
	} // Runs on localhost:8080
}
