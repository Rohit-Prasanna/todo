package routes

import (
	"github.com/gin-gonic/gin"
	"todo/handlers"
)

func InitializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
		api.GET("/users/:id", handlers.GetUserByID)
	}
}
