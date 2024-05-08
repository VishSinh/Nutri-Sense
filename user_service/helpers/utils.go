package helpers

import (
	"github.com/gin-gonic/gin"
)

func ResponseObj(status int, success bool, message string, data any, ctx *gin.Context) {
	ctx.JSON(status, gin.H{
		"status":  status,
		"success": success,
		"message": message,
		"data":    data,
	})
}
