package response

import (
	"github.com/gin-gonic/gin"
)

func WithSuccess(c *gin.Context, statusCode int, message string, payload any) {
	c.JSON(statusCode, gin.H{
		"status":  "success",
		"error":   nil,
		"message": message,
		"payload": payload,
	})
}

func WithError(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"error":   err.Error(),
		"message": message,
		"payload": nil,
	})
}
