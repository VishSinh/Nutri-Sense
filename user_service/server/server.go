package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"user_service/handlers"
	"user_service/helpers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	log.Println("Server started at " + helpers.ServerURL)

	// Initialize handlers with the database
	userHandler := handlers.UserHandler{DB: db}

	SetupRoutes(r, &userHandler)

	return r
}
