package auth

import (
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT verification
		c.Next()
	}
}
