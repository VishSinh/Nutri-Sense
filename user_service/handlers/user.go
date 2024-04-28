package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"user_service/helpers"
	"user_service/models"
)

type UserHandler struct {
	DB *gorm.DB
}

func Ping(ctx *gin.Context) {
	helpers.ResponseObj(200, false, "Meow", gin.H{"msg": "Hello"}, ctx)
}

func (uh *UserHandler) SignUp(ctx *gin.Context) {

	var user struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.ResponseObj(http.StatusBadRequest, false, "Invalid request payload", gin.H{}, ctx)
		return
	}

	var existingUser models.User
	if err := uh.DB.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser).Error; err == nil {
		if existingUser.Email == user.Email {
			helpers.ResponseObj(http.StatusBadRequest, false, "Email already exists", gin.H{}, ctx)
		} else {
			helpers.ResponseObj(http.StatusBadRequest, false, "Username already exists", gin.H{}, ctx)
		}
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "An error occured", gin.H{}, ctx)
		return
	}

	newUser := models.User{
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	if err := uh.DB.Create(&newUser).Error; err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "Failed to create user", gin.H{}, ctx)
		return
	}

	responseData := gin.H{
		"ID":       newUser.ID,
		"Email":    newUser.Email,
		"Username": newUser.Username,
	}

	helpers.ResponseObj(http.StatusCreated, true, "User created successfully", responseData, ctx)

}
