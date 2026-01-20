package handler

import (
	"GoWebapitest/internal/core/service"

	"github.com/gin-gonic/gin"
)

func HistoriqueMiddleware(audit *service.HistoriqueService) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, ok := c.Get("user_id")
		if !ok {
			return
		}

		userName, _ := c.Get("user_name")

		err := audit.LogAction(c.FullPath(), userID.(string), userName.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
}
