package api

import (
	"github.com/Aman5681/notify/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", LoginHandler)
		api.POST("/register", RegisterHandler)
		api.Use(auth.JWTMiddleware())
		{
			api.POST("/notify", NotifyHandler)
		}
	}

	return r
}
