package server

import (
	"user_service/handlers"
	"user_service/helpers"
	"user_service/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the server and sets up the routes
func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	api := r.Group("/api")

	// Api Versioning
	v1 := api.Group("/v1")

	// Unprotected routes
	v1.POST(helpers.SignUpURL, userHandler.SignUp)
	v1.POST(helpers.LoginURL, userHandler.Login)

	// V1 Protected routes
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.POST(helpers.AddUserDetails, userHandler.AddUserDetails)
		v1.GET(helpers.GetUserDetails, userHandler.GetUserDetails)
	}
}
