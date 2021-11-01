package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := "Bearer "
		token := c.Request.Header.Get("Authorization")

		if !strings.Contains(token, bearer) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Your request is not authorized"})
			c.Abort()
			return
		}

		splitToken := strings.Split(token, bearer)
		if len(splitToken) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "An authorization token was not supplied"})
			c.Abort()
			return
		}

		valid, err := ValidateToken(splitToken[1], JWT_SIGNATURE_KEY)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		claims, ok := valid.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Set("userInfo", claims)
		c.Next()
	}
}
