package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "user_service/handlers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	fmt.Println("Server started at http://127.0.0.1:8080")

	// Write an example endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// // Unprotected routes
	// r.POST("/login", handlers.Login)
	// r.POST("/signup", handlers.Signup)

	// // Protected routes
	// protected := r.Group("/api")
	// protected.Use(auth.AuthMiddleware())
	// {
	// 	protected.GET("/users/:id", handlers.GetUser)
	// 	protected.PUT("/users/:id", handlers.UpdateUser)
	// 	protected.DELETE("/users/:id", handlers.DeleteUser)
	// }

	return r
}
