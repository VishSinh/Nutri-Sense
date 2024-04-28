package helpers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ResponseObj(status int, success bool, message string, data any, ctx *gin.Context) {
	ctx.JSON(status, gin.H{
		"status":  status,
		"success": success,
		"message": message,
		"data":    data,
	})
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
