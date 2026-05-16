package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		claims := user.(jwt.MapClaims)
		if claims["role"] != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Admin Access Only",
			})
			c.Abort()
			return
		}

		c.Next()
	}

}
