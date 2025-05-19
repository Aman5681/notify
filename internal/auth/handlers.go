package auth

import (
	"net/http"

	"github.com/Aman5681/notify/internal/db"
	"github.com/Aman5681/notify/internal/db/models"
	"github.com/Aman5681/notify/internal/db/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserId    string `json:"userId"`
	EmailId   string `json:"emailId"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

var MockDB = make(map[string]user) // Temporary in-memory store

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

	user := user{
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

func SignUpUserHandler(c *gin.Context) {
	var input struct {
		EmailId   string `json:"emailId"`
		Password  string `json:"password"`
		Role      string `json:"role"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Phone     string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inpur"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		UserId:    uuid.New(),
		EmailId:   input.EmailId,
		Password:  string(hashedPassword),
		Role:      input.Role,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
	}

	repo := repositories.UserRepository{DB: db.DB}
	if err := repo.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "userId": user.UserId})
}
