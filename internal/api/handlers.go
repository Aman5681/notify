package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// TODO: Implement login
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}

func RegisterHandler(c *gin.Context) {
	// TODO: Implement registration
	c.JSON(http.StatusOK, gin.H{"message": "Register"})
}

func NotifyHandler(c *gin.Context) {
	// TODO: Implement notification logic
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}
