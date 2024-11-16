package main

import (
	"github.com/gin-gonic/gin"
	"todo/routes"
)

func main() {
	r := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(r)

	// Start the server
	r.Run(":8080") // Runs on localhost:8080
}
