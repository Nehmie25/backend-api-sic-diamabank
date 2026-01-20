package handler

import (
	"net/http"
	"strings"

	"GoWebapitest/internal/core/port"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(token port.TokenGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		claims, err := token.Validate(strings.TrimPrefix(auth, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["id"])
		c.Set("user_name", claims["nom"])
		c.Set("user_role", claims["role"])
		c.Set("user_email", claims["email"])

		c.Next()
	}
}
