package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := "Bearer "
		token := c.Request.Header.Get("Authorization")

		if !strings.Contains(token, bearer) {
			c.JSON(403, gin.H{"message": "Your request is not authorized"})
			c.Abort()
			return
		}

		splitToken := strings.Split(token, bearer)
		if len(splitToken) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied"})
			c.Abort()
			return
		}

		valid, err := ValidateToken(splitToken[1], JWT_SIGNATURE_KEY)
		if err != nil {
			c.JSON(403, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		claims, ok := valid.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(403, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Set("userInfo", claims)
		c.Next()
	}
}
