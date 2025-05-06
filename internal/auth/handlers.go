package auth

import (
	"net/http"

	"github.com/Aman5681/notify/internal/db/models"
	"github.com/gin-gonic/gin"
)

var MockDB = make(map[string]models.User) // Temporary in-memory store

func LoginHandler(c *gin.Context) {
	var input struct {
		EmailId  string `json:"emailId"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, ok := MockDB[input.EmailId]
	if !ok || CheckPassword(user.Password, input.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid creds"})
		return
	}

	authToken, refreshToken, _ := GenerateTokens(user.EmailId)
	c.JSON(http.StatusOK, gin.H{"authToken": authToken, "refreshToken": refreshToken})
}

func RegisterHandler(c *gin.Context) {
	var input struct {
		EmailId  string `json:"emailId"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inpur"})
		return
	}

	hashed, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		UserId:   input.EmailId, // just for mock
		EmailId:  input.EmailId,
		Password: hashed,
		Role:     input.Role,
	}
	MockDB[user.EmailId] = user
	c.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
}

func NotifyHandler(c *gin.Context) {
	// TODO: Implement notification logic
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}
