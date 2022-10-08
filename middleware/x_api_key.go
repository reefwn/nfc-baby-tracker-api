package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("Authorization")

		if key != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
		}
	}
}
