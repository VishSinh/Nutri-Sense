package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"user_service/helpers"
	"user_service/helpers/auth"
	"user_service/models"
)

type UserHandler struct {
	DB *gorm.DB
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

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "An error occured", gin.H{}, ctx)
		return
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "An error occured", gin.H{}, ctx)
		return
	}

	newUser := models.User{
		ID:       userId,
		Email:    user.Email,
		Username: user.Username,
		Password: hashedPassword,
	}

	if err := uh.DB.Create(&newUser).Error; err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "Failed to create user", gin.H{}, ctx)
		return
	}

	token, err := auth.CreateToken(newUser.ID)
	if err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "Failed to create user", gin.H{}, ctx)
		return
	}

	responseData := gin.H{
		"id":    newUser.ID,
		"token": token,
	}

	helpers.ResponseObj(http.StatusCreated, true, "User created successfully", responseData, ctx)

}

func (uh *UserHandler) Login(ctx *gin.Context) {

	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.ResponseObj(http.StatusBadRequest, false, "Invalid request payload", gin.H{}, ctx)
		return
	}

	var existingUser models.User
	if err := uh.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		helpers.ResponseObj(http.StatusNotFound, false, "User not found", gin.H{}, ctx)
		return
	}

	if !auth.VerifyPassword(existingUser.Password, user.Password) {
		helpers.ResponseObj(http.StatusUnauthorized, false, "Invalid credentials", gin.H{}, ctx)
		return
	}

	token, err := auth.CreateToken(existingUser.ID)
	if err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "Failed to create user", gin.H{}, ctx)
		return
	}

	responseData := gin.H{
		"id":    existingUser.ID,
		"token": token,
	}

	helpers.ResponseObj(http.StatusOK, true, "User logged in successfully", responseData, ctx)
}

func (uh *UserHandler) AddUserDetails(ctx *gin.Context) {
	var userDetails struct {
		Name   string  `json:"name" binding:"required"`
		Age    int     `json:"age" binding:"required"`
		Weight float32 `json:"weight" binding:"required"`
		Height *int    `json:"height"`
	}

	if err := ctx.ShouldBindJSON(&userDetails); err != nil {
		helpers.ResponseObj(http.StatusBadRequest, false, "Invalid request payload", gin.H{}, ctx)
		return
	}

	if userDetails.Age < 1 || userDetails.Weight < 1 || (userDetails.Height != nil && *userDetails.Height < 1) {
		helpers.ResponseObj(http.StatusBadRequest, false, "Invalid request payload", gin.H{}, ctx)
		return
	}

	userID := ctx.MustGet("userID").(uuid.UUID)
	if userID == uuid.Nil {
		helpers.ResponseObj(http.StatusUnauthorized, false, "Invalid user", gin.H{}, ctx)
		return
	}

	var existingUser models.User
	if err := uh.DB.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		helpers.ResponseObj(http.StatusNotFound, false, "User not found", gin.H{}, ctx)
		return
	}

	newUserDetails := models.UserDetails{
		UserID: userID,
		Name:   userDetails.Name,
		Age:    userDetails.Age,
		Weight: userDetails.Weight,
	}

	if userDetails.Height != nil {
		newUserDetails.Height = *userDetails.Height
	}

	if err := uh.DB.Create(&newUserDetails).Error; err != nil {
		helpers.ResponseObj(http.StatusInternalServerError, false, "Failed to add user details", gin.H{}, ctx)
		return
	}

	helpers.ResponseObj(http.StatusCreated, true, "User details added successfully", gin.H{}, ctx)
}

func (uh *UserHandler) GetUserDetails(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uuid.UUID)
	if userID == uuid.Nil {
		helpers.ResponseObj(http.StatusUnauthorized, false, "Invalid user", gin.H{}, ctx)
		return
	}

	var userDetails models.UserDetails
	if err := uh.DB.Where("user_id = ?", userID).First(&userDetails).Error; err != nil {
		helpers.ResponseObj(http.StatusNotFound, false, "User details not found", gin.H{}, ctx)
		return
	}

	var user models.User
	if err := uh.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		helpers.ResponseObj(http.StatusNotFound, false, "User not found", gin.H{}, ctx)
		return
	}

	responseData := gin.H{
		"username": user.Username,
		"name":     userDetails.Name,
		"age":      userDetails.Age,
		"weight":   userDetails.Weight,
		"height":   userDetails.Height,
	}

	helpers.ResponseObj(http.StatusOK, true, "User details fetched successfully", responseData, ctx)
}
