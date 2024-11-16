package routes

import (
	"github.com/Rohit-Prasanna/todo/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
		api.GET("/users/:id", handlers.GetUserByID)
	}
}
