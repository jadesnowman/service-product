package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := true
		c.Set("example", "12345")

		if !status {
			c.JSON(403, gin.H{"message": "Your request is not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
