package api

import (
	"github.com/Aman5681/notify/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", auth.LoginHandler)
	r.POST("/register", auth.RegisterHandler)
	r.POST("/refresh-auth-token", auth.RefreshHandler)
	// protected routes

	api := r.Group("/api")
	{
		api.Use(auth.JWTMiddleware())
		{
			api.POST("/notify", auth.NotifyHandler)
		}
	}

	return r
}
